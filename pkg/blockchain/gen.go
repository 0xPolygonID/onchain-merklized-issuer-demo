// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blockchain

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ClaimBuilderClaimData is an auto generated low-level Go binding around an user-defined struct.
type ClaimBuilderClaimData struct {
	SchemaHash            *big.Int
	IdPosition            uint8
	Expirable             bool
	Updatable             bool
	MerklizedRootPosition uint8
	Version               uint32
	Id                    *big.Int
	RevocationNonce       uint64
	ExpirationDate        uint64
	MerklizedRoot         *big.Int
	IndexDataSlotA        *big.Int
	IndexDataSlotB        *big.Int
	ValueDataSlotA        *big.Int
	ValueDataSlotB        *big.Int
}

// IOnchainCredentialStatusResolverCredentialStatus is an auto generated low-level Go binding around an user-defined struct.
type IOnchainCredentialStatusResolverCredentialStatus struct {
	Issuer IOnchainCredentialStatusResolverIdentityStateRoots
	Mtp    IOnchainCredentialStatusResolverProof
}

// IOnchainCredentialStatusResolverIdentityStateRoots is an auto generated low-level Go binding around an user-defined struct.
type IOnchainCredentialStatusResolverIdentityStateRoots struct {
	State              *big.Int
	ClaimsTreeRoot     *big.Int
	RevocationTreeRoot *big.Int
	RootOfRoots        *big.Int
}

// IOnchainCredentialStatusResolverProof is an auto generated low-level Go binding around an user-defined struct.
type IOnchainCredentialStatusResolverProof struct {
	Root         *big.Int
	Existence    bool
	Siblings     []*big.Int
	Index        *big.Int
	Value        *big.Int
	AuxExistence bool
	AuxIndex     *big.Int
	AuxValue     *big.Int
}

// IdentityLibRoots is an auto generated low-level Go binding around an user-defined struct.
type IdentityLibRoots struct {
	ClaimsRoot      *big.Int
	RevocationsRoot *big.Int
	RootsRoot       *big.Int
}

// SmtLibProof is an auto generated low-level Go binding around an user-defined struct.
type SmtLibProof struct {
	Root         *big.Int
	Existence    bool
	Siblings     []*big.Int
	Index        *big.Int
	Value        *big.Int
	AuxExistence bool
	AuxIndex     *big.Int
	AuxValue     *big.Int
}

// IdentityMetaData contains all meta data concerning the Identity contract.
var IdentityMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"claim\",\"type\":\"uint256[8]\"}],\"name\":\"addClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"claim\",\"type\":\"uint256[8]\"}],\"name\":\"addClaimAndTransit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"hashIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hashValue\",\"type\":\"uint256\"}],\"name\":\"addClaimHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"hashIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"hashValue\",\"type\":\"uint256\"}],\"name\":\"addClaimHashAndTransit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"schemaHash\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"idPosition\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"expirable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"updatable\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"merklizedRootPosition\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expirationDate\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"merklizedRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexDataSlotA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexDataSlotB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueDataSlotA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueDataSlotB\",\"type\":\"uint256\"}],\"internalType\":\"structClaimBuilder.ClaimData\",\"name\":\"claimData\",\"type\":\"tuple\"}],\"name\":\"buildClaim\",\"outputs\":[{\"internalType\":\"uint256[8]\",\"name\":\"\",\"type\":\"uint256[8]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"calcIdentityState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimIndexHash\",\"type\":\"uint256\"}],\"name\":\"getClaimProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimIndexHash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getClaimProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getClaimsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getIsOldStateGenesis\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedClaimsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedRevocationsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedRootsRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLatestPublishedState\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getRevocationProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationStatus\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootOfRoots\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.IdentityStateRoots\",\"name\":\"issuer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.Proof\",\"name\":\"mtp\",\"type\":\"tuple\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.CredentialStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"nonce\",\"type\":\"uint64\"}],\"name\":\"getRevocationStatusByIdAndState\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootOfRoots\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.IdentityStateRoots\",\"name\":\"issuer\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.Proof\",\"name\":\"mtp\",\"type\":\"tuple\"}],\"internalType\":\"structIOnchainCredentialStatusResolver.CredentialStatus\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRevocationsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"}],\"name\":\"getRootProof\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"claimsTreeRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"}],\"name\":\"getRootProofByRoot\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"root\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"existence\",\"type\":\"bool\"},{\"internalType\":\"uint256[]\",\"name\":\"siblings\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"auxExistence\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"auxIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"auxValue\",\"type\":\"uint256\"}],\"internalType\":\"structSmtLib.Proof\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"state\",\"type\":\"uint256\"}],\"name\":\"getRootsByState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"claimsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"revocationsRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rootsRoot\",\"type\":\"uint256\"}],\"internalType\":\"structIdentityLib.Roots\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRootsTreeRoot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSmtDepth\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stateContractAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"newClaimData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"schemaHash\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"idPosition\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"expirable\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"updatable\",\"type\":\"bool\"},{\"internalType\":\"uint8\",\"name\":\"merklizedRootPosition\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"version\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"expirationDate\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"merklizedRoot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexDataSlotA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexDataSlotB\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueDataSlotA\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valueDataSlotB\",\"type\":\"uint256\"}],\"internalType\":\"structClaimBuilder.ClaimData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"}],\"name\":\"revokeClaim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"revocationNonce\",\"type\":\"uint64\"}],\"name\":\"revokeClaimAndTransit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transitState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IdentityABI is the input ABI used to generate the binding from.
// Deprecated: Use IdentityMetaData.ABI instead.
var IdentityABI = IdentityMetaData.ABI

// Identity is an auto generated Go binding around an Ethereum contract.
type Identity struct {
	IdentityCaller     // Read-only binding to the contract
	IdentityTransactor // Write-only binding to the contract
	IdentityFilterer   // Log filterer for contract events
}

// IdentityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentitySession struct {
	Contract     *Identity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityCallerSession struct {
	Contract *IdentityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdentityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityTransactorSession struct {
	Contract     *IdentityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRaw struct {
	Contract *Identity // Generic contract binding to access the raw methods on
}

// IdentityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityCallerRaw struct {
	Contract *IdentityCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityTransactorRaw struct {
	Contract *IdentityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentity creates a new instance of Identity, bound to a specific deployed contract.
func NewIdentity(address common.Address, backend bind.ContractBackend) (*Identity, error) {
	contract, err := bindIdentity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// NewIdentityCaller creates a new read-only instance of Identity, bound to a specific deployed contract.
func NewIdentityCaller(address common.Address, caller bind.ContractCaller) (*IdentityCaller, error) {
	contract, err := bindIdentity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityCaller{contract: contract}, nil
}

// NewIdentityTransactor creates a new write-only instance of Identity, bound to a specific deployed contract.
func NewIdentityTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityTransactor, error) {
	contract, err := bindIdentity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityTransactor{contract: contract}, nil
}

// NewIdentityFilterer creates a new log filterer instance of Identity, bound to a specific deployed contract.
func NewIdentityFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFilterer, error) {
	contract, err := bindIdentity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFilterer{contract: contract}, nil
}

// bindIdentity binds a generic wrapper to an already deployed contract.
func bindIdentity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IdentityMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.IdentityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transact(opts, method, params...)
}

// BuildClaim is a free data retrieval call binding the contract method 0x73c08244.
//
// Solidity: function buildClaim((uint256,uint8,bool,bool,uint8,uint32,uint256,uint64,uint64,uint256,uint256,uint256,uint256,uint256) claimData) pure returns(uint256[8])
func (_Identity *IdentityCaller) BuildClaim(opts *bind.CallOpts, claimData ClaimBuilderClaimData) ([8]*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "buildClaim", claimData)

	if err != nil {
		return *new([8]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([8]*big.Int)).(*[8]*big.Int)

	return out0, err

}

// BuildClaim is a free data retrieval call binding the contract method 0x73c08244.
//
// Solidity: function buildClaim((uint256,uint8,bool,bool,uint8,uint32,uint256,uint64,uint64,uint256,uint256,uint256,uint256,uint256) claimData) pure returns(uint256[8])
func (_Identity *IdentitySession) BuildClaim(claimData ClaimBuilderClaimData) ([8]*big.Int, error) {
	return _Identity.Contract.BuildClaim(&_Identity.CallOpts, claimData)
}

// BuildClaim is a free data retrieval call binding the contract method 0x73c08244.
//
// Solidity: function buildClaim((uint256,uint8,bool,bool,uint8,uint32,uint256,uint64,uint64,uint256,uint256,uint256,uint256,uint256) claimData) pure returns(uint256[8])
func (_Identity *IdentityCallerSession) BuildClaim(claimData ClaimBuilderClaimData) ([8]*big.Int, error) {
	return _Identity.Contract.BuildClaim(&_Identity.CallOpts, claimData)
}

// CalcIdentityState is a free data retrieval call binding the contract method 0x4ddad0b3.
//
// Solidity: function calcIdentityState() view returns(uint256)
func (_Identity *IdentityCaller) CalcIdentityState(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "calcIdentityState")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcIdentityState is a free data retrieval call binding the contract method 0x4ddad0b3.
//
// Solidity: function calcIdentityState() view returns(uint256)
func (_Identity *IdentitySession) CalcIdentityState() (*big.Int, error) {
	return _Identity.Contract.CalcIdentityState(&_Identity.CallOpts)
}

// CalcIdentityState is a free data retrieval call binding the contract method 0x4ddad0b3.
//
// Solidity: function calcIdentityState() view returns(uint256)
func (_Identity *IdentityCallerSession) CalcIdentityState() (*big.Int, error) {
	return _Identity.Contract.CalcIdentityState(&_Identity.CallOpts)
}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCaller) GetClaimProof(opts *bind.CallOpts, claimIndexHash *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getClaimProof", claimIndexHash)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentitySession) GetClaimProof(claimIndexHash *big.Int) (SmtLibProof, error) {
	return _Identity.Contract.GetClaimProof(&_Identity.CallOpts, claimIndexHash)
}

// GetClaimProof is a free data retrieval call binding the contract method 0xb57a40cb.
//
// Solidity: function getClaimProof(uint256 claimIndexHash) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCallerSession) GetClaimProof(claimIndexHash *big.Int) (SmtLibProof, error) {
	return _Identity.Contract.GetClaimProof(&_Identity.CallOpts, claimIndexHash)
}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCaller) GetClaimProofByRoot(opts *bind.CallOpts, claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getClaimProofByRoot", claimIndexHash, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentitySession) GetClaimProofByRoot(claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	return _Identity.Contract.GetClaimProofByRoot(&_Identity.CallOpts, claimIndexHash, root)
}

// GetClaimProofByRoot is a free data retrieval call binding the contract method 0x310d0d5b.
//
// Solidity: function getClaimProofByRoot(uint256 claimIndexHash, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCallerSession) GetClaimProofByRoot(claimIndexHash *big.Int, root *big.Int) (SmtLibProof, error) {
	return _Identity.Contract.GetClaimProofByRoot(&_Identity.CallOpts, claimIndexHash, root)
}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_Identity *IdentityCaller) GetClaimsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getClaimsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_Identity *IdentitySession) GetClaimsTreeRoot() (*big.Int, error) {
	return _Identity.Contract.GetClaimsTreeRoot(&_Identity.CallOpts)
}

// GetClaimsTreeRoot is a free data retrieval call binding the contract method 0x3df432fc.
//
// Solidity: function getClaimsTreeRoot() view returns(uint256)
func (_Identity *IdentityCallerSession) GetClaimsTreeRoot() (*big.Int, error) {
	return _Identity.Contract.GetClaimsTreeRoot(&_Identity.CallOpts)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_Identity *IdentityCaller) GetId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_Identity *IdentitySession) GetId() (*big.Int, error) {
	return _Identity.Contract.GetId(&_Identity.CallOpts)
}

// GetId is a free data retrieval call binding the contract method 0x5d1ca631.
//
// Solidity: function getId() view returns(uint256)
func (_Identity *IdentityCallerSession) GetId() (*big.Int, error) {
	return _Identity.Contract.GetId(&_Identity.CallOpts)
}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_Identity *IdentityCaller) GetIsOldStateGenesis(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getIsOldStateGenesis")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_Identity *IdentitySession) GetIsOldStateGenesis() (bool, error) {
	return _Identity.Contract.GetIsOldStateGenesis(&_Identity.CallOpts)
}

// GetIsOldStateGenesis is a free data retrieval call binding the contract method 0xf84c7c1e.
//
// Solidity: function getIsOldStateGenesis() view returns(bool)
func (_Identity *IdentityCallerSession) GetIsOldStateGenesis() (bool, error) {
	return _Identity.Contract.GetIsOldStateGenesis(&_Identity.CallOpts)
}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_Identity *IdentityCaller) GetLatestPublishedClaimsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getLatestPublishedClaimsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_Identity *IdentitySession) GetLatestPublishedClaimsRoot() (*big.Int, error) {
	return _Identity.Contract.GetLatestPublishedClaimsRoot(&_Identity.CallOpts)
}

// GetLatestPublishedClaimsRoot is a free data retrieval call binding the contract method 0x523b8136.
//
// Solidity: function getLatestPublishedClaimsRoot() view returns(uint256)
func (_Identity *IdentityCallerSession) GetLatestPublishedClaimsRoot() (*big.Int, error) {
	return _Identity.Contract.GetLatestPublishedClaimsRoot(&_Identity.CallOpts)
}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_Identity *IdentityCaller) GetLatestPublishedRevocationsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getLatestPublishedRevocationsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_Identity *IdentitySession) GetLatestPublishedRevocationsRoot() (*big.Int, error) {
	return _Identity.Contract.GetLatestPublishedRevocationsRoot(&_Identity.CallOpts)
}

// GetLatestPublishedRevocationsRoot is a free data retrieval call binding the contract method 0x9674cfa4.
//
// Solidity: function getLatestPublishedRevocationsRoot() view returns(uint256)
func (_Identity *IdentityCallerSession) GetLatestPublishedRevocationsRoot() (*big.Int, error) {
	return _Identity.Contract.GetLatestPublishedRevocationsRoot(&_Identity.CallOpts)
}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_Identity *IdentityCaller) GetLatestPublishedRootsRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getLatestPublishedRootsRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_Identity *IdentitySession) GetLatestPublishedRootsRoot() (*big.Int, error) {
	return _Identity.Contract.GetLatestPublishedRootsRoot(&_Identity.CallOpts)
}

// GetLatestPublishedRootsRoot is a free data retrieval call binding the contract method 0xc6365a3b.
//
// Solidity: function getLatestPublishedRootsRoot() view returns(uint256)
func (_Identity *IdentityCallerSession) GetLatestPublishedRootsRoot() (*big.Int, error) {
	return _Identity.Contract.GetLatestPublishedRootsRoot(&_Identity.CallOpts)
}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_Identity *IdentityCaller) GetLatestPublishedState(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getLatestPublishedState")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_Identity *IdentitySession) GetLatestPublishedState() (*big.Int, error) {
	return _Identity.Contract.GetLatestPublishedState(&_Identity.CallOpts)
}

// GetLatestPublishedState is a free data retrieval call binding the contract method 0x3d59ec60.
//
// Solidity: function getLatestPublishedState() view returns(uint256)
func (_Identity *IdentityCallerSession) GetLatestPublishedState() (*big.Int, error) {
	return _Identity.Contract.GetLatestPublishedState(&_Identity.CallOpts)
}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCaller) GetRevocationProof(opts *bind.CallOpts, revocationNonce uint64) (SmtLibProof, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getRevocationProof", revocationNonce)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentitySession) GetRevocationProof(revocationNonce uint64) (SmtLibProof, error) {
	return _Identity.Contract.GetRevocationProof(&_Identity.CallOpts, revocationNonce)
}

// GetRevocationProof is a free data retrieval call binding the contract method 0x26485063.
//
// Solidity: function getRevocationProof(uint64 revocationNonce) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCallerSession) GetRevocationProof(revocationNonce uint64) (SmtLibProof, error) {
	return _Identity.Contract.GetRevocationProof(&_Identity.CallOpts, revocationNonce)
}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCaller) GetRevocationProofByRoot(opts *bind.CallOpts, revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getRevocationProofByRoot", revocationNonce, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentitySession) GetRevocationProofByRoot(revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	return _Identity.Contract.GetRevocationProofByRoot(&_Identity.CallOpts, revocationNonce, root)
}

// GetRevocationProofByRoot is a free data retrieval call binding the contract method 0xe26ecb0b.
//
// Solidity: function getRevocationProofByRoot(uint64 revocationNonce, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCallerSession) GetRevocationProofByRoot(revocationNonce uint64, root *big.Int) (SmtLibProof, error) {
	return _Identity.Contract.GetRevocationProofByRoot(&_Identity.CallOpts, revocationNonce, root)
}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_Identity *IdentityCaller) GetRevocationStatus(opts *bind.CallOpts, id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getRevocationStatus", id, nonce)

	if err != nil {
		return *new(IOnchainCredentialStatusResolverCredentialStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IOnchainCredentialStatusResolverCredentialStatus)).(*IOnchainCredentialStatusResolverCredentialStatus)

	return out0, err

}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_Identity *IdentitySession) GetRevocationStatus(id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _Identity.Contract.GetRevocationStatus(&_Identity.CallOpts, id, nonce)
}

// GetRevocationStatus is a free data retrieval call binding the contract method 0x110c96a7.
//
// Solidity: function getRevocationStatus(uint256 id, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_Identity *IdentityCallerSession) GetRevocationStatus(id *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _Identity.Contract.GetRevocationStatus(&_Identity.CallOpts, id, nonce)
}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_Identity *IdentityCaller) GetRevocationStatusByIdAndState(opts *bind.CallOpts, id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getRevocationStatusByIdAndState", id, state, nonce)

	if err != nil {
		return *new(IOnchainCredentialStatusResolverCredentialStatus), err
	}

	out0 := *abi.ConvertType(out[0], new(IOnchainCredentialStatusResolverCredentialStatus)).(*IOnchainCredentialStatusResolverCredentialStatus)

	return out0, err

}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_Identity *IdentitySession) GetRevocationStatusByIdAndState(id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _Identity.Contract.GetRevocationStatusByIdAndState(&_Identity.CallOpts, id, state, nonce)
}

// GetRevocationStatusByIdAndState is a free data retrieval call binding the contract method 0xaad72921.
//
// Solidity: function getRevocationStatusByIdAndState(uint256 id, uint256 state, uint64 nonce) view returns(((uint256,uint256,uint256,uint256),(uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256)))
func (_Identity *IdentityCallerSession) GetRevocationStatusByIdAndState(id *big.Int, state *big.Int, nonce uint64) (IOnchainCredentialStatusResolverCredentialStatus, error) {
	return _Identity.Contract.GetRevocationStatusByIdAndState(&_Identity.CallOpts, id, state, nonce)
}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_Identity *IdentityCaller) GetRevocationsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getRevocationsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_Identity *IdentitySession) GetRevocationsTreeRoot() (*big.Int, error) {
	return _Identity.Contract.GetRevocationsTreeRoot(&_Identity.CallOpts)
}

// GetRevocationsTreeRoot is a free data retrieval call binding the contract method 0x01c85c77.
//
// Solidity: function getRevocationsTreeRoot() view returns(uint256)
func (_Identity *IdentityCallerSession) GetRevocationsTreeRoot() (*big.Int, error) {
	return _Identity.Contract.GetRevocationsTreeRoot(&_Identity.CallOpts)
}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 claimsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCaller) GetRootProof(opts *bind.CallOpts, claimsTreeRoot *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getRootProof", claimsTreeRoot)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 claimsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentitySession) GetRootProof(claimsTreeRoot *big.Int) (SmtLibProof, error) {
	return _Identity.Contract.GetRootProof(&_Identity.CallOpts, claimsTreeRoot)
}

// GetRootProof is a free data retrieval call binding the contract method 0xc1e32733.
//
// Solidity: function getRootProof(uint256 claimsTreeRoot) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCallerSession) GetRootProof(claimsTreeRoot *big.Int) (SmtLibProof, error) {
	return _Identity.Contract.GetRootProof(&_Identity.CallOpts, claimsTreeRoot)
}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCaller) GetRootProofByRoot(opts *bind.CallOpts, claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getRootProofByRoot", claimsTreeRoot, root)

	if err != nil {
		return *new(SmtLibProof), err
	}

	out0 := *abi.ConvertType(out[0], new(SmtLibProof)).(*SmtLibProof)

	return out0, err

}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentitySession) GetRootProofByRoot(claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	return _Identity.Contract.GetRootProofByRoot(&_Identity.CallOpts, claimsTreeRoot, root)
}

// GetRootProofByRoot is a free data retrieval call binding the contract method 0x2d5c4f25.
//
// Solidity: function getRootProofByRoot(uint256 claimsTreeRoot, uint256 root) view returns((uint256,bool,uint256[],uint256,uint256,bool,uint256,uint256))
func (_Identity *IdentityCallerSession) GetRootProofByRoot(claimsTreeRoot *big.Int, root *big.Int) (SmtLibProof, error) {
	return _Identity.Contract.GetRootProofByRoot(&_Identity.CallOpts, claimsTreeRoot, root)
}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_Identity *IdentityCaller) GetRootsByState(opts *bind.CallOpts, state *big.Int) (IdentityLibRoots, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getRootsByState", state)

	if err != nil {
		return *new(IdentityLibRoots), err
	}

	out0 := *abi.ConvertType(out[0], new(IdentityLibRoots)).(*IdentityLibRoots)

	return out0, err

}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_Identity *IdentitySession) GetRootsByState(state *big.Int) (IdentityLibRoots, error) {
	return _Identity.Contract.GetRootsByState(&_Identity.CallOpts, state)
}

// GetRootsByState is a free data retrieval call binding the contract method 0xb8db6871.
//
// Solidity: function getRootsByState(uint256 state) view returns((uint256,uint256,uint256))
func (_Identity *IdentityCallerSession) GetRootsByState(state *big.Int) (IdentityLibRoots, error) {
	return _Identity.Contract.GetRootsByState(&_Identity.CallOpts, state)
}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_Identity *IdentityCaller) GetRootsTreeRoot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getRootsTreeRoot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_Identity *IdentitySession) GetRootsTreeRoot() (*big.Int, error) {
	return _Identity.Contract.GetRootsTreeRoot(&_Identity.CallOpts)
}

// GetRootsTreeRoot is a free data retrieval call binding the contract method 0xda68a0b1.
//
// Solidity: function getRootsTreeRoot() view returns(uint256)
func (_Identity *IdentityCallerSession) GetRootsTreeRoot() (*big.Int, error) {
	return _Identity.Contract.GetRootsTreeRoot(&_Identity.CallOpts)
}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_Identity *IdentityCaller) GetSmtDepth(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getSmtDepth")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_Identity *IdentitySession) GetSmtDepth() (*big.Int, error) {
	return _Identity.Contract.GetSmtDepth(&_Identity.CallOpts)
}

// GetSmtDepth is a free data retrieval call binding the contract method 0x3f0c6648.
//
// Solidity: function getSmtDepth() pure returns(uint256)
func (_Identity *IdentityCallerSession) GetSmtDepth() (*big.Int, error) {
	return _Identity.Contract.GetSmtDepth(&_Identity.CallOpts)
}

// NewClaimData is a free data retrieval call binding the contract method 0x5b9c895d.
//
// Solidity: function newClaimData() pure returns((uint256,uint8,bool,bool,uint8,uint32,uint256,uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_Identity *IdentityCaller) NewClaimData(opts *bind.CallOpts) (ClaimBuilderClaimData, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "newClaimData")

	if err != nil {
		return *new(ClaimBuilderClaimData), err
	}

	out0 := *abi.ConvertType(out[0], new(ClaimBuilderClaimData)).(*ClaimBuilderClaimData)

	return out0, err

}

// NewClaimData is a free data retrieval call binding the contract method 0x5b9c895d.
//
// Solidity: function newClaimData() pure returns((uint256,uint8,bool,bool,uint8,uint32,uint256,uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_Identity *IdentitySession) NewClaimData() (ClaimBuilderClaimData, error) {
	return _Identity.Contract.NewClaimData(&_Identity.CallOpts)
}

// NewClaimData is a free data retrieval call binding the contract method 0x5b9c895d.
//
// Solidity: function newClaimData() pure returns((uint256,uint8,bool,bool,uint8,uint32,uint256,uint64,uint64,uint256,uint256,uint256,uint256,uint256))
func (_Identity *IdentityCallerSession) NewClaimData() (ClaimBuilderClaimData, error) {
	return _Identity.Contract.NewClaimData(&_Identity.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identity *IdentityCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identity *IdentitySession) Owner() (common.Address, error) {
	return _Identity.Contract.Owner(&_Identity.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Identity *IdentityCallerSession) Owner() (common.Address, error) {
	return _Identity.Contract.Owner(&_Identity.CallOpts)
}

// AddClaim is a paid mutator transaction binding the contract method 0x5aa7396b.
//
// Solidity: function addClaim(uint256[8] claim) returns()
func (_Identity *IdentityTransactor) AddClaim(opts *bind.TransactOpts, claim [8]*big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addClaim", claim)
}

// AddClaim is a paid mutator transaction binding the contract method 0x5aa7396b.
//
// Solidity: function addClaim(uint256[8] claim) returns()
func (_Identity *IdentitySession) AddClaim(claim [8]*big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddClaim(&_Identity.TransactOpts, claim)
}

// AddClaim is a paid mutator transaction binding the contract method 0x5aa7396b.
//
// Solidity: function addClaim(uint256[8] claim) returns()
func (_Identity *IdentityTransactorSession) AddClaim(claim [8]*big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddClaim(&_Identity.TransactOpts, claim)
}

// AddClaimAndTransit is a paid mutator transaction binding the contract method 0x412363d3.
//
// Solidity: function addClaimAndTransit(uint256[8] claim) returns()
func (_Identity *IdentityTransactor) AddClaimAndTransit(opts *bind.TransactOpts, claim [8]*big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addClaimAndTransit", claim)
}

// AddClaimAndTransit is a paid mutator transaction binding the contract method 0x412363d3.
//
// Solidity: function addClaimAndTransit(uint256[8] claim) returns()
func (_Identity *IdentitySession) AddClaimAndTransit(claim [8]*big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddClaimAndTransit(&_Identity.TransactOpts, claim)
}

// AddClaimAndTransit is a paid mutator transaction binding the contract method 0x412363d3.
//
// Solidity: function addClaimAndTransit(uint256[8] claim) returns()
func (_Identity *IdentityTransactorSession) AddClaimAndTransit(claim [8]*big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddClaimAndTransit(&_Identity.TransactOpts, claim)
}

// AddClaimHash is a paid mutator transaction binding the contract method 0x5584b091.
//
// Solidity: function addClaimHash(uint256 hashIndex, uint256 hashValue) returns()
func (_Identity *IdentityTransactor) AddClaimHash(opts *bind.TransactOpts, hashIndex *big.Int, hashValue *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addClaimHash", hashIndex, hashValue)
}

// AddClaimHash is a paid mutator transaction binding the contract method 0x5584b091.
//
// Solidity: function addClaimHash(uint256 hashIndex, uint256 hashValue) returns()
func (_Identity *IdentitySession) AddClaimHash(hashIndex *big.Int, hashValue *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddClaimHash(&_Identity.TransactOpts, hashIndex, hashValue)
}

// AddClaimHash is a paid mutator transaction binding the contract method 0x5584b091.
//
// Solidity: function addClaimHash(uint256 hashIndex, uint256 hashValue) returns()
func (_Identity *IdentityTransactorSession) AddClaimHash(hashIndex *big.Int, hashValue *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddClaimHash(&_Identity.TransactOpts, hashIndex, hashValue)
}

// AddClaimHashAndTransit is a paid mutator transaction binding the contract method 0x961fc416.
//
// Solidity: function addClaimHashAndTransit(uint256 hashIndex, uint256 hashValue) returns()
func (_Identity *IdentityTransactor) AddClaimHashAndTransit(opts *bind.TransactOpts, hashIndex *big.Int, hashValue *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addClaimHashAndTransit", hashIndex, hashValue)
}

// AddClaimHashAndTransit is a paid mutator transaction binding the contract method 0x961fc416.
//
// Solidity: function addClaimHashAndTransit(uint256 hashIndex, uint256 hashValue) returns()
func (_Identity *IdentitySession) AddClaimHashAndTransit(hashIndex *big.Int, hashValue *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddClaimHashAndTransit(&_Identity.TransactOpts, hashIndex, hashValue)
}

// AddClaimHashAndTransit is a paid mutator transaction binding the contract method 0x961fc416.
//
// Solidity: function addClaimHashAndTransit(uint256 hashIndex, uint256 hashValue) returns()
func (_Identity *IdentityTransactorSession) AddClaimHashAndTransit(hashIndex *big.Int, hashValue *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddClaimHashAndTransit(&_Identity.TransactOpts, hashIndex, hashValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _stateContractAddr) returns()
func (_Identity *IdentityTransactor) Initialize(opts *bind.TransactOpts, _stateContractAddr common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "initialize", _stateContractAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _stateContractAddr) returns()
func (_Identity *IdentitySession) Initialize(_stateContractAddr common.Address) (*types.Transaction, error) {
	return _Identity.Contract.Initialize(&_Identity.TransactOpts, _stateContractAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _stateContractAddr) returns()
func (_Identity *IdentityTransactorSession) Initialize(_stateContractAddr common.Address) (*types.Transaction, error) {
	return _Identity.Contract.Initialize(&_Identity.TransactOpts, _stateContractAddr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identity *IdentityTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identity *IdentitySession) RenounceOwnership() (*types.Transaction, error) {
	return _Identity.Contract.RenounceOwnership(&_Identity.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Identity *IdentityTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Identity.Contract.RenounceOwnership(&_Identity.TransactOpts)
}

// RevokeClaim is a paid mutator transaction binding the contract method 0x891eb608.
//
// Solidity: function revokeClaim(uint64 revocationNonce) returns()
func (_Identity *IdentityTransactor) RevokeClaim(opts *bind.TransactOpts, revocationNonce uint64) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "revokeClaim", revocationNonce)
}

// RevokeClaim is a paid mutator transaction binding the contract method 0x891eb608.
//
// Solidity: function revokeClaim(uint64 revocationNonce) returns()
func (_Identity *IdentitySession) RevokeClaim(revocationNonce uint64) (*types.Transaction, error) {
	return _Identity.Contract.RevokeClaim(&_Identity.TransactOpts, revocationNonce)
}

// RevokeClaim is a paid mutator transaction binding the contract method 0x891eb608.
//
// Solidity: function revokeClaim(uint64 revocationNonce) returns()
func (_Identity *IdentityTransactorSession) RevokeClaim(revocationNonce uint64) (*types.Transaction, error) {
	return _Identity.Contract.RevokeClaim(&_Identity.TransactOpts, revocationNonce)
}

// RevokeClaimAndTransit is a paid mutator transaction binding the contract method 0xf2a8ed5a.
//
// Solidity: function revokeClaimAndTransit(uint64 revocationNonce) returns()
func (_Identity *IdentityTransactor) RevokeClaimAndTransit(opts *bind.TransactOpts, revocationNonce uint64) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "revokeClaimAndTransit", revocationNonce)
}

// RevokeClaimAndTransit is a paid mutator transaction binding the contract method 0xf2a8ed5a.
//
// Solidity: function revokeClaimAndTransit(uint64 revocationNonce) returns()
func (_Identity *IdentitySession) RevokeClaimAndTransit(revocationNonce uint64) (*types.Transaction, error) {
	return _Identity.Contract.RevokeClaimAndTransit(&_Identity.TransactOpts, revocationNonce)
}

// RevokeClaimAndTransit is a paid mutator transaction binding the contract method 0xf2a8ed5a.
//
// Solidity: function revokeClaimAndTransit(uint64 revocationNonce) returns()
func (_Identity *IdentityTransactorSession) RevokeClaimAndTransit(revocationNonce uint64) (*types.Transaction, error) {
	return _Identity.Contract.RevokeClaimAndTransit(&_Identity.TransactOpts, revocationNonce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identity *IdentityTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identity *IdentitySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identity.Contract.TransferOwnership(&_Identity.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Identity *IdentityTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Identity.Contract.TransferOwnership(&_Identity.TransactOpts, newOwner)
}

// TransitState is a paid mutator transaction binding the contract method 0x6c87e2fa.
//
// Solidity: function transitState() returns()
func (_Identity *IdentityTransactor) TransitState(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "transitState")
}

// TransitState is a paid mutator transaction binding the contract method 0x6c87e2fa.
//
// Solidity: function transitState() returns()
func (_Identity *IdentitySession) TransitState() (*types.Transaction, error) {
	return _Identity.Contract.TransitState(&_Identity.TransactOpts)
}

// TransitState is a paid mutator transaction binding the contract method 0x6c87e2fa.
//
// Solidity: function transitState() returns()
func (_Identity *IdentityTransactorSession) TransitState() (*types.Transaction, error) {
	return _Identity.Contract.TransitState(&_Identity.TransactOpts)
}

// IdentityInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Identity contract.
type IdentityInitializedIterator struct {
	Event *IdentityInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityInitialized represents a Initialized event raised by the Identity contract.
type IdentityInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Identity *IdentityFilterer) FilterInitialized(opts *bind.FilterOpts) (*IdentityInitializedIterator, error) {

	logs, sub, err := _Identity.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &IdentityInitializedIterator{contract: _Identity.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Identity *IdentityFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *IdentityInitialized) (event.Subscription, error) {

	logs, sub, err := _Identity.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityInitialized)
				if err := _Identity.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Identity *IdentityFilterer) ParseInitialized(log types.Log) (*IdentityInitialized, error) {
	event := new(IdentityInitialized)
	if err := _Identity.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IdentityOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Identity contract.
type IdentityOwnershipTransferredIterator struct {
	Event *IdentityOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IdentityOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IdentityOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IdentityOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IdentityOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IdentityOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IdentityOwnershipTransferred represents a OwnershipTransferred event raised by the Identity contract.
type IdentityOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identity *IdentityFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IdentityOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Identity.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IdentityOwnershipTransferredIterator{contract: _Identity.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identity *IdentityFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IdentityOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Identity.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IdentityOwnershipTransferred)
				if err := _Identity.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Identity *IdentityFilterer) ParseOwnershipTransferred(log types.Log) (*IdentityOwnershipTransferred, error) {
	event := new(IdentityOwnershipTransferred)
	if err := _Identity.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
