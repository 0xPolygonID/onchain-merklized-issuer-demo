package issuer

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"encoding/json"
	"io"
	"log/slog"
	"math/big"
	"net/http"
	"net/url"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	ethclientlib "github.com/ethereum/go-ethereum/ethclient"
	core "github.com/iden3/go-iden3-core/v2"
	"github.com/iden3/go-iden3-core/v2/w3c"
	"github.com/iden3/go-merkletree-sql/v2"
	jsonproc "github.com/iden3/go-schema-processor/v2/json"
	"github.com/iden3/go-schema-processor/v2/merklize"
	"github.com/iden3/go-schema-processor/v2/processor"
	"github.com/iden3/go-schema-processor/v2/verifiable"
	"github.com/iden3/go-service-template/pkg/blockchain"
	"github.com/iden3/go-service-template/pkg/credential"
	"github.com/iden3/go-service-template/pkg/logger"
	"github.com/iden3/go-service-template/pkg/repository"
	"github.com/pkg/errors"
)

type IPFSCli interface {
	Cat(string) (io.ReadCloser, error)
}

type IssuerService struct {
	repository   *repository.CredentialRepository
	ethclients   map[string]*ethclientlib.Client
	privateKeys  map[string]string // NOTE: for production version change to vault
	merklizeOpts []merklize.MerklizeOption
	ipfsCli      IPFSCli
}

func NewIssuerService(
	credentialRepository *repository.CredentialRepository,
	ethclients map[string]*ethclientlib.Client,
	privateKeys map[string]string,
	ipfsCli IPFSCli,
	merklizeOpts ...merklize.MerklizeOption,
) *IssuerService {
	return &IssuerService{
		repository:   credentialRepository,
		ethclients:   ethclients,
		privateKeys:  privateKeys,
		merklizeOpts: merklizeOpts,
		ipfsCli:      ipfsCli,
	}
}

func (is *IssuerService) GetUserCredentials(
	ctx context.Context,
	issuer *w3c.DID,
	subject string,
	schemaType string,
) ([]*verifiable.W3CCredential, error) {
	return is.repository.GetByUser(ctx, issuer.String(), subject, schemaType)
}

func (is *IssuerService) GetCredentialByID(
	ctx context.Context,
	issuer *w3c.DID,
	credentialID string,
) (*verifiable.W3CCredential, error) {
	return is.repository.GetByID(ctx, issuer.String(), credentialID)
}

func (is *IssuerService) GetIssuersList(_ context.Context) []string {
	issuers := make([]string, 0, len(is.privateKeys))
	for issuer := range is.privateKeys {
		issuers = append(issuers, issuer)
	}
	return issuers
}

func (is *IssuerService) IsRevokedCredential(
	ctx context.Context,
	issuer *w3c.DID,
	nonce uint64,
) (bool, error) {
	ethclient, contractAddress, err := is.lookforEthConnectForDID(issuer)
	if err != nil {
		return false, errors.Wrapf(err, "error getting ethclient for issuer: '%s'", issuer)
	}

	revoked, err := blockchain.IsRevokedCredential(
		ctx,
		ethclient,
		contractAddress,
		nonce,
	)
	if err != nil {
		return false,
			errors.Wrapf(err, "error getting revocation proof for nonce: %d", nonce)
	}

	return revoked, nil
}

func (is *IssuerService) RevokeCredential(
	ctx context.Context,
	issuer *w3c.DID,
	nonce uint64,
) error {
	ethClient, contractAddress, err := is.lookforEthConnectForDID(issuer)
	if err != nil {
		return errors.Wrapf(err, "error getting ethclient for issuer: '%s'", issuer)
	}

	issuerPK, found := is.privateKeys[issuer.String()]
	if !found {
		return errors.Errorf("private key not found for issuer: '%s'", issuer)
	}

	transaction, err := blockchain.RevokeCredential(
		ctx,
		ethClient,
		issuerPK,
		contractAddress,
		nonce,
	)
	if err != nil {
		return errors.Wrap(err, "error sending transaction")
	}

	if err := blockchain.WaitTransaction(ctx, ethClient, transaction); err != nil {
		return errors.Wrap(err, "error waiting transaction")
	}

	return nil
}

func (is *IssuerService) IssueCredential(
	ctx context.Context,
	issuerDID *w3c.DID,
	credentialRequest *credential.CredentialRequest,
) (string, error) {
	if credentialRequest.SubjectPosition != "" {
		return "", errors.New("subject position must be empty for merklized credentials")
	}
	if credentialRequest.CredentialSchema == "" {
		return "", errors.New("credential schema is required")
	}

	if credentialRequest.RevNonce == nil {
		nonce, err := randInt64()
		if err != nil {
			return "", errors.Wrap(err, "error generating nonce")
		}
		credentialRequest.RevNonce = &nonce
	}
	if credentialRequest.MerklizedRootPosition == verifiable.CredentialMerklizedRootPositionNone {
		credentialRequest.MerklizedRootPosition = verifiable.CredentialMerklizedRootPositionIndex
	}

	schemaRawURL, err := url.Parse(credentialRequest.CredentialSchema)
	if err != nil {
		return "", errors.Wrap(err, "error parsing schema url")
	}

	var credentialSchemaBody io.ReadCloser
	switch schemaRawURL.Scheme {
	case "ipfs":
		credentialSchemaBody, err = is.ipfsCli.Cat(schemaRawURL.Host)
		if err != nil {
			return "", errors.Wrap(err, "error getting schema")
		}
	case "http", "https":
		resp, err := http.DefaultClient.Get(credentialRequest.CredentialSchema)
		if err != nil {
			return "", errors.Wrap(err, "error getting schema")
		}
		if resp.StatusCode != http.StatusOK {
			_ = resp.Body.Close()
			return "", errors.Errorf("error getting schema: %s", resp.Status)
		}
		credentialSchemaBody = resp.Body
	default:
		return "", errors.Errorf("unsupported schema url '%s'", schemaRawURL.Scheme)
	}
	defer credentialSchemaBody.Close()

	var jsonSchemaMetadata credential.JSONSchemaMetadata
	if err = json.NewDecoder(credentialSchemaBody).Decode(&jsonSchemaMetadata); err != nil {
		return "", errors.Wrap(err, "error decoding json schema to metadata")
	}
	verifiableCredential, err := credential.NewCredential(issuerDID, credentialRequest, jsonSchemaMetadata)
	if err != nil {
		return "", errors.Wrap(err, "error creating credential")
	}

	printVerifiableCredential(verifiableCredential)

	coreClaim, err := jsonproc.Parser{}.ParseClaim(
		ctx,
		*verifiableCredential,
		&processor.CoreClaimOptions{
			RevNonce:              *credentialRequest.RevNonce,
			Version:               credentialRequest.Version,
			SubjectPosition:       credentialRequest.SubjectPosition,
			MerklizedRootPosition: credentialRequest.MerklizedRootPosition,
			Updatable:             false,
			MerklizerOpts:         is.merklizeOpts,
		})
	if err != nil {
		return "", errors.Wrap(err, "error build core claim")
	}

	ethclient, contractAddress, err := is.lookforEthConnectForDID(issuerDID)
	if err != nil {
		return "", errors.Wrapf(err, "error getting ethclient for issuer: '%s'", issuerDID)
	}
	issuerPK, found := is.privateKeys[issuerDID.String()]
	if !found {
		return "", errors.Errorf("private key not found for issuer: '%s'", issuerDID)
	}

	hindex, err := coreClaim.HIndex()
	if err != nil {
		return "", errors.Wrap(err, "error getting hash index")
	}
	vindex, err := coreClaim.HValue()
	if err != nil {
		return "", errors.Wrap(err, "error getting verification method index")
	}

	transaction, err := blockchain.IssueCredential(
		ctx,
		ethclient,
		issuerPK,
		contractAddress,
		hindex, vindex,
	)
	if err != nil {
		return "", errors.Wrap(err, "error sending transaction")
	}

	if err := blockchain.WaitTransaction(ctx, ethclient, transaction); err != nil {
		return "", errors.Wrap(err, "error waiting transaction")
	}

	mtpProof, err := extractMTPProof(
		ctx, hindex, ethclient, contractAddress)
	if err != nil {
		return "", errors.Wrap(err, "error extracting mtp proof")
	}
	mtpProof.IssuerData.ID = issuerDID.String()
	mtpProof.CoreClaim, err = coreClaim.Hex()
	if err != nil {
		return "", errors.Wrap(err, "error getting core claim")
	}
	verifiableCredential.Proof = verifiable.CredentialProofs{&mtpProof}

	id, err := is.repository.Create(ctx, verifiableCredential)
	if err != nil {
		return "", errors.Wrap(err, "error creating credential")
	}
	return id, nil
}

func (is *IssuerService) lookforEthConnectForDID(did *w3c.DID) (ethclient *ethclientlib.Client, contractAddress string, err error) {
	issuerID, err := core.IDFromDID(*did)
	if err != nil {
		return nil, "", errors.Wrap(err, "error getting issuer ID")
	}
	networkID, err := core.ChainIDfromDID(*did)
	if err != nil {
		return nil, "", errors.Wrapf(err, "network not found for did '%s'", did)
	}
	ethclient, found := is.ethclients[strconv.Itoa(int(networkID))]
	if !found {
		return nil, "", errors.Errorf("ethclient not found for network id '%d'", networkID)
	}
	contractAddressBytes, err := core.EthAddressFromID(issuerID)
	if err != nil {
		return nil, "", errors.Wrap(err, "error getting contract address")
	}
	return ethclient, common.BytesToAddress(contractAddressBytes[:]).Hex(), nil
}

func extractMTPProof(
	ctx context.Context,
	hindex *big.Int,
	ethclient *ethclientlib.Client,
	contractAddress string,
) (verifiable.Iden3SparseMerkleTreeProof, error) {
	proof, latestState, roots, err := blockchain.ReadMTPProof(
		ctx,
		hindex,
		ethclient,
		contractAddress,
	)
	if err != nil {
		return verifiable.Iden3SparseMerkleTreeProof{}, err
	}
	if !proof.Existence {
		return verifiable.Iden3SparseMerkleTreeProof{},
			errors.Errorf("hindex '%s' does not exist in the issuer's claim tree", hindex)
	}

	state, err := merkletree.NewHashFromBigInt(latestState)
	if err != nil {
		return verifiable.Iden3SparseMerkleTreeProof{}, err
	}
	rootOfRoots, err := merkletree.NewHashFromBigInt(roots.RootsRoot)
	if err != nil {
		return verifiable.Iden3SparseMerkleTreeProof{}, err
	}
	claimTreeRoot, err := merkletree.NewHashFromBigInt(roots.ClaimsRoot)
	if err != nil {
		return verifiable.Iden3SparseMerkleTreeProof{}, err
	}
	revocationTreeRoot, err := merkletree.NewHashFromBigInt(roots.RevocationsRoot)
	if err != nil {
		return verifiable.Iden3SparseMerkleTreeProof{}, err
	}

	mtp, err := convertChainProofToMerkleProof(&proof)
	if err != nil {
		return verifiable.Iden3SparseMerkleTreeProof{}, err
	}

	return verifiable.Iden3SparseMerkleTreeProof{
		Type: verifiable.Iden3SparseMerkleTreeProofType,
		IssuerData: verifiable.IssuerData{
			State: verifiable.State{
				RootOfRoots:        strpoint(rootOfRoots.Hex()),
				ClaimsTreeRoot:     strpoint(claimTreeRoot.Hex()),
				RevocationTreeRoot: strpoint(revocationTreeRoot.Hex()),
				Value:              strpoint(state.Hex()),
			},
		},
		MTP: mtp,
	}, nil
}

func convertChainProofToMerkleProof(proof *blockchain.SmtLibProof) (*merkletree.Proof, error) {
	nodeAuxIndex, err := merkletree.NewHashFromBigInt(
		proof.AuxIndex,
	)
	if err != nil {
		return nil, err
	}
	nodeAuxValue, err := merkletree.NewHashFromBigInt(
		proof.AuxValue,
	)
	if err != nil {
		return nil, err
	}

	siblings := make([]*merkletree.Hash, 0, len(proof.Siblings))
	for _, s := range proof.Siblings {
		h, err := merkletree.NewHashFromBigInt(s)
		if err != nil {
			return nil, err
		}
		siblings = append(siblings, h)
	}

	if proof.AuxExistence {
		return merkletree.NewProofFromData(
			proof.Existence,
			siblings,
			&merkletree.NodeAux{
				Key:   nodeAuxIndex,
				Value: nodeAuxValue,
			},
		)
	}

	return merkletree.NewProofFromData(
		proof.Existence,
		siblings,
		nil,
	)
}

func strpoint(s string) *string {
	return &s
}

func randInt64() (uint64, error) {
	var buf [8]byte
	_, err := rand.Read(buf[:4])
	if err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(buf[:]), nil
}

func printVerifiableCredential(data *verifiable.W3CCredential) {
	credentialBytes, err := json.Marshal(data)
	if err != nil {
		logger.WithError(err).Error("error marshalizing credential for logging")
	}
	logger.Debug("verifiable credential: ", slog.String("credential", string(credentialBytes)))
}
