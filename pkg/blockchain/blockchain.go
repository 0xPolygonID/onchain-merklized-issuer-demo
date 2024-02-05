package blockchain

import (
	"context"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	ethclientlib "github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

const blockConfirmations = 3

func IsRevokedCredential(
	ctx context.Context,
	ethclient *ethclientlib.Client,
	contractAddress string,
	nonce uint64,
) (bool, error) {
	contract, err := NewIdentity(
		common.HexToAddress(contractAddress),
		ethclient,
	)
	if err != nil {
		return false, errors.Wrapf(err, "failed to instantiate contract")
	}
	opts := &bind.CallOpts{
		Context: ctx,
	}
	proof, err := contract.GetRevocationProof(opts, nonce)
	if err != nil {
		return false, errors.Wrapf(err, "failed to get revocation proof")
	}
	return proof.Existence, nil
}

func RevokeCredential(
	ctx context.Context,
	ethclient *ethclientlib.Client,
	pk string,
	contractAddress string,
	nonce uint64,
) (*types.Transaction, error) {
	contract, err := NewIdentity(
		common.HexToAddress(contractAddress),
		ethclient,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to instantiate contract")
	}
	tx, err := SendTransaction(ctx, ethclient, pk, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.RevokeClaimAndTransit(opts, nonce)
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to send transaction")
	}
	return tx, nil
}

func IssueCredential(
	ctx context.Context,
	ethclient *ethclientlib.Client,
	pk string,
	contractAddress string,
	hindex, vindex *big.Int,
) (*types.Transaction, error) {
	contract, err := NewIdentity(
		common.HexToAddress(contractAddress),
		ethclient,
	)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to instantiate contract")
	}
	tx, err := SendTransaction(ctx, ethclient, pk, func(opts *bind.TransactOpts) (*types.Transaction, error) {
		return contract.AddClaimHashAndTransit(opts, hindex, vindex)
	})
	if err != nil {
		return nil, errors.Wrapf(err, "failed to send transaction")
	}
	return tx, nil
}

func ReadMTPProof(
	ctx context.Context,
	hidx *big.Int,
	ethclient *ethclientlib.Client,
	contractAddress string,
) (proof SmtLibProof, latestState *big.Int, roots IdentityLibRoots, err error) {
	contract, err := NewIdentity(
		common.HexToAddress(contractAddress),
		ethclient,
	)
	if err != nil {
		return proof, latestState, roots, errors.Wrapf(err, "failed to instantiate contract")
	}
	proof, err = contract.GetClaimProof(&bind.CallOpts{Context: ctx}, hidx)
	if err != nil {
		return proof, latestState, roots, err
	}
	bigState, err := contract.GetLatestPublishedState(&bind.CallOpts{Context: ctx})
	if err != nil {
		return proof, latestState, roots, err
	}
	roots, err = contract.GetRootsByState(&bind.CallOpts{Context: ctx}, bigState)
	if err != nil {
		return proof, latestState, roots, err
	}

	return proof, bigState, roots, nil
}

func SendTransaction(
	ctx context.Context,
	ethclient *ethclientlib.Client,
	pk string,
	call func(*bind.TransactOpts) (*types.Transaction, error),
) (*types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(pk)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse private key")
	}

	chainID, err := ethclient.ChainID(ctx)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get chain id")
	}

	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to create transactor")
	}
	opts.Context = ctx

	return call(opts)
}

func WaitTransaction(ctx context.Context, ethclient *ethclientlib.Client, tx *types.Transaction) error {
	receipt, err := bind.WaitMined(ctx, ethclient, tx)
	if err != nil {
		return errors.Wrapf(err, "failed to wait transaction")
	}

	tick := time.NewTicker(5 * time.Second)
	defer tick.Stop()

	for {
		<-tick.C
		header, err := ethclient.HeaderByNumber(ctx, nil)
		if err != nil {
			return errors.Wrapf(err, "failed to get header")
		}
		passedBlocks := big.NewInt(0).Sub(header.Number, receipt.BlockNumber)
		if big.NewInt(blockConfirmations).Cmp(passedBlocks) == -1 {
			return nil
		}
	}
}
