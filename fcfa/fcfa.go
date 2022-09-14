// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fcfa

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
)

// FeeCollector is an auto generated low-level Go binding around an user-defined struct.
type FeeCollector struct {
	AccountPubKey common.Address
	Numerator     *big.Int
	Denominateur  *big.Int
}

// FcfaMetaData contains all meta data concerning the Fcfa contract.
var FcfaMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddKYCFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApproveDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ApproveWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"canceler\",\"type\":\"address\"}],\"name\":\"CancelDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"canceler\",\"type\":\"address\"}],\"name\":\"CancelWithdraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"remover\",\"type\":\"address\"}],\"name\":\"RemoveKYCFlag\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"AGENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"BURNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FREEZE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"KYC_AGENT\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGMENT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactionId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"addUnconfirmedPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"approveDeposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"approveWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burnTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"canceler\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"cancelDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"canceler\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"cancelWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactionId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"commercepay\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactionId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"confirmUnconfirmedPayment\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"freezeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"giveRoleToAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"accountPubKey\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"numerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominateur\",\"type\":\"uint256\"}],\"internalType\":\"structFeeCollector[]\",\"name\":\"_feeCollectors\",\"type\":\"tuple[]\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactionId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"refund\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"remover\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeKycFlag\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"role\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeroleFromAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"unFreezeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"transactionnId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60a06040523073ffffffffffffffffffffffffffffffffffffffff1660809073ffffffffffffffffffffffffffffffffffffffff168152503480156200004457600080fd5b50620000556200005b60201b60201c565b62000206565b600060019054906101000a900460ff1615620000ae576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000a590620001a9565b60405180910390fd5b60ff801660008054906101000a900460ff1660ff161015620001205760ff6000806101000a81548160ff021916908360ff1602179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860ff604051620001179190620001e9565b60405180910390a15b565b600082825260208201905092915050565b7f496e697469616c697a61626c653a20636f6e747261637420697320696e69746960008201527f616c697a696e6700000000000000000000000000000000000000000000000000602082015250565b60006200019160278362000122565b91506200019e8262000133565b604082019050919050565b60006020820190508181036000830152620001c48162000182565b9050919050565b600060ff82169050919050565b620001e381620001cb565b82525050565b6000602082019050620002006000830184620001d8565b92915050565b608051618b7d6200023e6000396000818161152c015281816115ba015281816119f401528181611a820152611b320152618b7d6000f3fe6080604052600436106102ff5760003560e01c806370a0823111610190578063c087abdb116100dc578063e63ab1e911610095578063f26c159f1161006f578063f26c159f14610bb5578063f9613bb414610bde578063f9e8572314610c1b578063fe82df1514610c46576102ff565b8063e63ab1e914610b24578063ea785a5e14610b4f578063f19b7df014610b78576102ff565b8063c087abdb14610a04578063c3d017d614610a2d578063d539139314610a6a578063d547741f14610a95578063dd62ed3e14610abe578063e1825cf714610afb576102ff565b806395d89b4111610149578063a457c2d711610123578063a457c2d714610936578063a9059cbb14610973578063b90df18c146109b0578063bf376af6146109d9576102ff565b806395d89b41146108b75780639913c7bd146108e2578063a217fddf1461090b576102ff565b806370a082311461079757806379cc6790146107d45780638456cb59146107fd578063854ed6491461081457806391d148541461083d57806392b4dfd21461087a576102ff565b80633659cfe61161024f5780634e09b9571161020857806353cc2fae116101e257806353cc2fae146106dd5780635c975abb146107065780635e19e7b91461073157806366e2cd861461076e576102ff565b80634e09b9571461066d5780634f1ef2861461069657806352d1902d146106b2576102ff565b80633659cfe614610573578063395093511461059c5780633bcb220e146105d95780633f4ba83a1461060457806340c10f191461061b57806342966c6814610644576102ff565b806323b872dd116102bc5780632d7cdeb1116102965780632d7cdeb1146104cd5780632f2ff15d146104f6578063313ce5671461051f57806336568abe1461054a576102ff565b806323b872dd14610428578063248a9ca314610465578063282c51f3146104a2576102ff565b806301ffc9a71461030457806306fdde0314610341578063095ea7b31461036c57806318160ddd146103a957806318449c23146103d457806322459e18146103fd575b600080fd5b34801561031057600080fd5b5061032b60048036038101906103269190616705565b610c83565b604051610338919061674d565b60405180910390f35b34801561034d57600080fd5b50610356610cfd565b60405161036391906167f8565b60405180910390f35b34801561037857600080fd5b50610393600480360381019061038e91906168ae565b610d8f565b6040516103a0919061674d565b60405180910390f35b3480156103b557600080fd5b506103be610db2565b6040516103cb91906168fd565b60405180910390f35b3480156103e057600080fd5b506103fb60048036038101906103f69190616aee565b610dbc565b005b34801561040957600080fd5b50610412611063565b60405161041f9190616b92565b60405180910390f35b34801561043457600080fd5b5061044f600480360381019061044a9190616bad565b611087565b60405161045c919061674d565b60405180910390f35b34801561047157600080fd5b5061048c60048036038101906104879190616c2c565b6110b6565b6040516104999190616b92565b60405180910390f35b3480156104ae57600080fd5b506104b76110d6565b6040516104c49190616b92565b60405180910390f35b3480156104d957600080fd5b506104f460048036038101906104ef9190616c59565b6110fa565b005b34801561050257600080fd5b5061051d60048036038101906105189190616cf8565b61147d565b005b34801561052b57600080fd5b5061053461149e565b6040516105419190616d54565b60405180910390f35b34801561055657600080fd5b50610571600480360381019061056c9190616cf8565b6114a7565b005b34801561057f57600080fd5b5061059a60048036038101906105959190616d6f565b61152a565b005b3480156105a857600080fd5b506105c360048036038101906105be91906168ae565b6116b2565b6040516105d0919061674d565b60405180910390f35b3480156105e557600080fd5b506105ee6116e9565b6040516105fb9190616b92565b60405180910390f35b34801561061057600080fd5b5061061961170d565b005b34801561062757600080fd5b50610642600480360381019061063d91906168ae565b611742565b005b34801561065057600080fd5b5061066b60048036038101906106669190616d9c565b61177b565b005b34801561067957600080fd5b50610694600480360381019061068f9190616dc9565b61178f565b005b6106b060048036038101906106ab9190616e68565b6119f2565b005b3480156106be57600080fd5b506106c7611b2e565b6040516106d49190616b92565b60405180910390f35b3480156106e957600080fd5b5061070460048036038101906106ff9190616d6f565b611be7565b005b34801561071257600080fd5b5061071b611c65565b604051610728919061674d565b60405180910390f35b34801561073d57600080fd5b5061075860048036038101906107539190616aee565b611c7c565b604051610765919061674d565b60405180910390f35b34801561077a57600080fd5b5061079560048036038101906107909190616ff5565b612148565b005b3480156107a357600080fd5b506107be60048036038101906107b99190616d6f565b6123e7565b6040516107cb91906168fd565b60405180910390f35b3480156107e057600080fd5b506107fb60048036038101906107f691906168ae565b612430565b005b34801561080957600080fd5b50610812612450565b005b34801561082057600080fd5b5061083b60048036038101906108369190616aee565b612485565b005b34801561084957600080fd5b50610864600480360381019061085f9190616cf8565b61272c565b604051610871919061674d565b60405180910390f35b34801561088657600080fd5b506108a1600480360381019061089c919061709c565b612797565b6040516108ae919061674d565b60405180910390f35b3480156108c357600080fd5b506108cc612895565b6040516108d991906167f8565b60405180910390f35b3480156108ee57600080fd5b5061090960048036038101906109049190616c59565b612927565b005b34801561091757600080fd5b50610920612c26565b60405161092d9190616b92565b60405180910390f35b34801561094257600080fd5b5061095d600480360381019061095891906168ae565b612c2d565b60405161096a919061674d565b60405180910390f35b34801561097f57600080fd5b5061099a600480360381019061099591906168ae565b612ca4565b6040516109a7919061674d565b60405180910390f35b3480156109bc57600080fd5b506109d760048036038101906109d29190616dc9565b612cbb565b005b3480156109e557600080fd5b506109ee612f1e565b6040516109fb9190616b92565b60405180910390f35b348015610a1057600080fd5b50610a2b6004803603810190610a269190616aee565b612f42565b005b348015610a3957600080fd5b50610a546004803603810190610a4f9190616aee565b6132af565b604051610a61919061674d565b60405180910390f35b348015610a7657600080fd5b50610a7f613623565b604051610a8c9190616b92565b60405180910390f35b348015610aa157600080fd5b50610abc6004803603810190610ab79190616cf8565b613647565b005b348015610aca57600080fd5b50610ae56004803603810190610ae0919061711f565b613668565b604051610af291906168fd565b60405180910390f35b348015610b0757600080fd5b50610b226004803603810190610b1d919061715f565b6136ef565b005b348015610b3057600080fd5b50610b396139cb565b604051610b469190616b92565b60405180910390f35b348015610b5b57600080fd5b50610b766004803603810190610b7191906168ae565b6139ef565b005b348015610b8457600080fd5b50610b9f6004803603810190610b9a91906171e2565b613a28565b604051610bac919061674d565b60405180910390f35b348015610bc157600080fd5b50610bdc6004803603810190610bd79190616d6f565b61402e565b005b348015610bea57600080fd5b50610c056004803603810190610c009190616aee565b6140b5565b604051610c12919061674d565b60405180910390f35b348015610c2757600080fd5b50610c3061456b565b604051610c3d9190616b92565b60405180910390f35b348015610c5257600080fd5b50610c6d6004803603810190610c689190617295565b61458f565b604051610c7a919061674d565b60405180910390f35b60007f7965db0b000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610cf65750610cf58261474b565b5b9050919050565b606060368054610d0c90617333565b80601f0160208091040260200160405190810160405280929190818152602001828054610d3890617333565b8015610d855780601f10610d5a57610100808354040283529160200191610d85565b820191906000526020600020905b815481529060010190602001808311610d6857829003601f168201915b5050505050905090565b600080610d9a6147b5565b9050610da78185856147bd565b600191505092915050565b6000603554905090565b60008383604051602001610dd19291906173e8565b6040516020818303038152906040528051906020012090506000610df58284614986565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525090610e9d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e9491906167f8565b60405180910390fd5b508373ffffffffffffffffffffffffffffffffffffffff1661019486604051610ec69190617410565b908152602001604051809103902060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161480610f405750610f3f7fcab5a0bfe0b79d2c4b1c2e02599fa044d115b7511f9659307cb42769509677098561272c565b5b6040518060400160405280600281526020017f313700000000000000000000000000000000000000000000000000000000000081525090610fb7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610fae91906167f8565b60405180910390fd5b50600161019486604051610fcb9190617410565b908152602001604051809103902060030160156101000a81548160ff02191690836004811115610ffe57610ffd617427565b5b02179055508373ffffffffffffffffffffffffffffffffffffffff16856040516110289190617410565b60405180910390207f2d365e7e41ce37b636dac5f661aadcd49e813866124180449e057e59ea79c32360405160405180910390a35050505050565b7fcab5a0bfe0b79d2c4b1c2e02599fa044d115b7511f9659307cb427695096770981565b6000806110926147b5565b905061109f8582856149ab565b6110aa858585614a37565b60019150509392505050565b600060fb6000838152602001908152602001600020600101549050919050565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84881565b600084848460405160200161111193929190617477565b60405160208183030381529060405280519060200120905060006111358284614986565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f3700000000000000000000000000000000000000000000000000000000000000815250906111dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111d491906167f8565b60405180910390fd5b50846111e8856123e7565b10156040518060400160405280600281526020017f313600000000000000000000000000000000000000000000000000000000000081525090611261576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161125891906167f8565b60405180910390fd5b506040518060c001604052808781526020018681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001600160028111156112c5576112c4617427565b5b8152602001600060048111156112de576112dd617427565b5b815250610194876040516112f29190617410565b90815260200160405180910390206000820151816000019081611315919061765c565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160146101000a81548160ff021916908360028111156113d9576113d8617427565b5b021790555060a08201518160030160156101000a81548160ff0219169083600481111561140957611408617427565b5b02179055509050508373ffffffffffffffffffffffffffffffffffffffff16866040516114369190617410565b60405180910390207f7a5e5c901f9da945e2028bd646eb1f6842dff416e862cfa502e0d3408635143a8760405161146d91906168fd565b60405180910390a3505050505050565b611486826110b6565b61148f81614df1565b6114998383614e05565b505050565b60006012905090565b6114af6147b5565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161461151c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611513906177a0565b60405180910390fd5b6115268282614ee6565b5050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff16036115b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115af90617832565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166115f7614fc8565b73ffffffffffffffffffffffffffffffffffffffff161461164d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611644906178c4565b60405180910390fd5b6116568161501f565b6116af81600067ffffffffffffffff81111561167557611674616922565b5b6040519080825280601f01601f1916602001820160405280156116a75781602001600182028036833780820191505090505b506000615030565b50565b6000806116bd6147b5565b90506116de8185856116cf8589613668565b6116d99190617913565b6147bd565b600191505092915050565b7fb30e940f336aba0589b4a82bc81c64c246cab8e8644222b4626d02f90a7786c281565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61173781614df1565b61173f61519e565b50565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a661176c81614df1565b6117768383615201565b505050565b61178c6117866147b5565b82615361565b50565b60008484846040516020016117a693929190617947565b60405160208183030381529060405280519060200120905060006117ca8284614986565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525090611872576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161186991906167f8565b60405180910390fd5b506000801b856040516020016118889190617410565b6040516020818303038152906040528051906020012014156040518060400160405280600281526020017f313400000000000000000000000000000000000000000000000000000000000081525090611917576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161190e91906167f8565b60405180910390fd5b506119427ff4a313569b11ee41331e9b404dc5be718fed2beb1eb41bb868890c5accf2789b8561272c565b6040518060400160405280600281526020017f3134000000000000000000000000000000000000000000000000000000000000815250906119b9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016119b091906167f8565b60405180910390fd5b506119ea856040516020016119ce9190617410565b6040516020818303038152906040528051906020012087614e05565b505050505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1603611a80576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a7790617832565b60405180910390fd5b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16611abf614fc8565b73ffffffffffffffffffffffffffffffffffffffff1614611b15576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b0c906178c4565b60405180910390fd5b611b1e8261501f565b611b2a82826001615030565b5050565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163073ffffffffffffffffffffffffffffffffffffffff1614611bbe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bb5906179f2565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b905090565b7f5789b43a60de35bcedee40618ae90979bab7d1315fd4b079234241bdab19936d611c1181614df1565b61019260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690555050565b6000609760009054906101000a900460ff16905090565b6000808484604051602001611c929291906173e8565b6040516020818303038152906040528051906020012090506000611cb68285614986565b90508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525090611d5e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d5591906167f8565b60405180910390fd5b50600061019487604051611d729190617410565b90815260200160405180910390206040518060c0016040529081600082018054611d9b90617333565b80601f0160208091040260200160405190810160405280929190818152602001828054611dc790617333565b8015611e145780601f10611de957610100808354040283529160200191611e14565b820191906000526020600020905b815481529060010190602001808311611df757829003601f168201915b50505050508152602001600182015481526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160149054906101000a900460ff166002811115611ef857611ef7617427565b5b6002811115611f0a57611f09617427565b5b81526020016003820160159054906101000a900460ff166004811115611f3357611f32617427565b5b6004811115611f4557611f44617427565b5b815250509050806060015173ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600281526020017f313700000000000000000000000000000000000000000000000000000000000081525090611ff5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fec91906167f8565b60405180910390fd5b5060005b610191805490508110156120db576000610191828154811061201e5761201d617a12565b5b906000526020600020906003020160020154610191838154811061204557612044617a12565b5b90600052602060002090600302016001015484602001516120669190617a41565b6120709190617aca565b90506120c7610191838154811061208a57612089617a12565b5b906000526020600020906003020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846060015183615539565b5080806120d390617afb565b915050611ff9565b506120ef8282604001518360200151615539565b6004610194886040516121029190617410565b908152602001604051809103902060030160156101000a81548160ff0219169083600481111561213557612134617427565b5b0217905550600193505050509392505050565b60008060019054906101000a900460ff161590508080156121795750600160008054906101000a900460ff1660ff16105b806121a65750612188306157bb565b1580156121a55750600160008054906101000a900460ff1660ff16145b5b6121e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121dc90617bb5565b60405180910390fd5b60016000806101000a81548160ff021916908360ff1602179055508015612222576001600060016101000a81548160ff0219169083151502179055505b61222c83836157de565b61223461583b565b61223c61588c565b6122446158e5565b61224c615936565b6122596000801b33614e05565b6122837f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a33614e05565b6122ad7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a633614e05565b6122d77f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a84833614e05565b6123017f5789b43a60de35bcedee40618ae90979bab7d1315fd4b079234241bdab19936d33614e05565b61232b7fcab5a0bfe0b79d2c4b1c2e02599fa044d115b7511f9659307cb427695096770933614e05565b6123557ff4a313569b11ee41331e9b404dc5be718fed2beb1eb41bb868890c5accf2789b33614e05565b61237f7fb30e940f336aba0589b4a82bc81c64c246cab8e8644222b4626d02f90a7786c233614e05565b61238884615987565b80156123e15760008060016101000a81548160ff0219169083151502179055507f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb384740249860016040516123d89190617c10565b60405180910390a15b50505050565b6000603360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050919050565b6124428261243c6147b5565b836149ab565b61244c8282615361565b5050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a61247a81614df1565b612482615acf565b50565b6000838360405160200161249a9291906173e8565b60405160208183030381529060405280519060200120905060006124be8284614986565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525090612566576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161255d91906167f8565b60405180910390fd5b508373ffffffffffffffffffffffffffffffffffffffff166101948660405161258f9190617410565b908152602001604051809103902060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148061260957506126087fcab5a0bfe0b79d2c4b1c2e02599fa044d115b7511f9659307cb42769509677098561272c565b5b6040518060400160405280600281526020017f313700000000000000000000000000000000000000000000000000000000000081525090612680576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161267791906167f8565b60405180910390fd5b506001610194866040516126949190617410565b908152602001604051809103902060030160156101000a81548160ff021916908360048111156126c7576126c6617427565b5b02179055508373ffffffffffffffffffffffffffffffffffffffff16856040516126f19190617410565b60405180910390207ff8006234b98a8a5d1104283e880b00bf159e0c65d70ab13ab479ae8cf6a4066260405160405180910390a35050505050565b600060fb600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16905092915050565b6000808585856040516020016127af93929190617c2b565b60405160208183030381529060405280519060200120905060006127d38285614986565b90508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f37000000000000000000000000000000000000000000000000000000000000008152509061287b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161287291906167f8565b60405180910390fd5b50612887818888614a37565b600192505050949350505050565b6060603780546128a490617333565b80601f01602080910402602001604051908101604052809291908181526020018280546128d090617333565b801561291d5780601f106128f25761010080835404028352916020019161291d565b820191906000526020600020905b81548152906001019060200180831161290057829003601f168201915b5050505050905090565b600084848460405160200161293e93929190617477565b60405160208183030381529060405280519060200120905060006129628284614986565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525090612a0a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a0191906167f8565b60405180910390fd5b506040518060c001604052808781526020018681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff16815260200160006002811115612a6e57612a6d617427565b5b815260200160006004811115612a8757612a86617427565b5b81525061019487604051612a9b9190617410565b90815260200160405180910390206000820151816000019081612abe919061765c565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160146101000a81548160ff02191690836002811115612b8257612b81617427565b5b021790555060a08201518160030160156101000a81548160ff02191690836004811115612bb257612bb1617427565b5b02179055509050508373ffffffffffffffffffffffffffffffffffffffff1686604051612bdf9190617410565b60405180910390207fd327b35e36b3981157588978d60961f5c09dc2926008abb81dd77b1197a416ed87604051612c1691906168fd565b60405180910390a3505050505050565b6000801b81565b600080612c386147b5565b90506000612c468286613668565b905083811015612c8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c8290617cda565b60405180910390fd5b612c9882868684036147bd565b60019250505092915050565b6000612cb1338484614a37565b6001905092915050565b6000848484604051602001612cd293929190617947565b6040516020818303038152906040528051906020012090506000612cf68284614986565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525090612d9e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d9591906167f8565b60405180910390fd5b506000801b85604051602001612db49190617410565b6040516020818303038152906040528051906020012014156040518060400160405280600281526020017f313400000000000000000000000000000000000000000000000000000000000081525090612e43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612e3a91906167f8565b60405180910390fd5b50612e6e7ff4a313569b11ee41331e9b404dc5be718fed2beb1eb41bb868890c5accf2789b8561272c565b6040518060400160405280600281526020017f313400000000000000000000000000000000000000000000000000000000000081525090612ee5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612edc91906167f8565b60405180910390fd5b50612f1685604051602001612efa9190617410565b6040516020818303038152906040528051906020012087614ee6565b505050505050565b7ff4a313569b11ee41331e9b404dc5be718fed2beb1eb41bb868890c5accf2789b81565b60008383604051602001612f579291906173e8565b6040516020818303038152906040528051906020012090506000612f7b8284614986565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525090613023576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161301a91906167f8565b60405180910390fd5b5061304e7fcab5a0bfe0b79d2c4b1c2e02599fa044d115b7511f9659307cb42769509677098561272c565b6040518060400160405280600281526020017f3134000000000000000000000000000000000000000000000000000000000000815250906130c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016130bc91906167f8565b60405180910390fd5b50600260048111156130da576130d9617427565b5b610194866040516130eb9190617410565b908152602001604051809103902060030160159054906101000a900460ff16600481111561311c5761311b617427565b5b14156040518060400160405280600281526020017f313700000000000000000000000000000000000000000000000000000000000081525090613195576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161318c91906167f8565b60405180910390fd5b50613204610194866040516131aa9190617410565b908152602001604051809103902060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16610194876040516131ed9190617410565b908152602001604051809103902060010154615361565b6002610194866040516132179190617410565b908152602001604051809103902060030160156101000a81548160ff0219169083600481111561324a57613249617427565b5b02179055508373ffffffffffffffffffffffffffffffffffffffff16856040516132749190617410565b60405180910390207fdf78545f0bfafc5998a08d764f2ac01a93d1b37fdf8a34e5ef83d439d331c23060405160405180910390a35050505050565b60008084846040516020016132c59291906173e8565b60405160208183030381529060405280519060200120905060006132e98285614986565b90508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525090613391576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161338891906167f8565b60405180910390fd5b506133bc7fcab5a0bfe0b79d2c4b1c2e02599fa044d115b7511f9659307cb42769509677098661272c565b6040518060400160405280600281526020017f313400000000000000000000000000000000000000000000000000000000000081525090613433576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161342a91906167f8565b60405180910390fd5b506002600481111561344857613447617427565b5b610194876040516134599190617410565b908152602001604051809103902060030160159054906101000a900460ff16600481111561348a57613489617427565b5b14156040518060400160405280600281526020017f313700000000000000000000000000000000000000000000000000000000000081525090613503576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134fa91906167f8565b60405180910390fd5b50613572610194876040516135189190617410565b908152602001604051809103902060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166101948860405161355b9190617410565b908152602001604051809103902060010154615201565b6002610194876040516135859190617410565b908152602001604051809103902060030160156101000a81548160ff021916908360048111156135b8576135b7617427565b5b02179055508473ffffffffffffffffffffffffffffffffffffffff16866040516135e29190617410565b60405180910390207ff92a27ff89640f3ea2022b7152bd3bd2e0193157d72edb6d9ae6c8f2230108b760405160405180910390a36001925050509392505050565b7f9f2df0fed2c77648de5860a4cc508cd0818c85b8b8a1ab4ceeef8d981c8956a681565b613650826110b6565b61365981614df1565b6136638383614ee6565b505050565b6000603460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905092915050565b6000846040516020016137029190617cfa565b60405160208183030381529060405280519060200120905060006137268284614986565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f3700000000000000000000000000000000000000000000000000000000000000815250906137ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137c591906167f8565b60405180910390fd5b506137f97fb30e940f336aba0589b4a82bc81c64c246cab8e8644222b4626d02f90a7786c28561272c565b6040518060400160405280600281526020017f313400000000000000000000000000000000000000000000000000000000000081525090613870576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161386791906167f8565b60405180910390fd5b508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614156040518060400160405280600281526020017f313700000000000000000000000000000000000000000000000000000000000081525090613918576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161390f91906167f8565b60405180910390fd5b5061019360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81549060ff02191690558473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167ffb456b249009c196b58086480400da57f56b6214a5bccad64b4f952e909969e760405160405180910390a3505050505050565b7f65d7a28e3265b37a6474929f336521b332c1681b933f6cb9f3376673440d862a81565b7f3c11d16cbaffd01df69ce1c404f6340ee057498f5f00246190ea54220576a848613a1981614df1565b613a238383615361565b505050565b60008086868686604051602001613a429493929190617d15565b6040516020818303038152906040528051906020012090506000613a668285614986565b90508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525090613b0e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b0591906167f8565b60405180910390fd5b506001151561019260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514156040518060400160405280600281526020017f313800000000000000000000000000000000000000000000000000000000000081525090613bdc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613bd391906167f8565b60405180910390fd5b506001151561019260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514156040518060400160405280600281526020017f313800000000000000000000000000000000000000000000000000000000000081525090613caa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613ca191906167f8565b60405180910390fd5b5085613cb5826123e7565b10156040518060400160405280600281526020017f313000000000000000000000000000000000000000000000000000000000000081525090613d2e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d2591906167f8565b60405180910390fd5b506000805b61019180549050811015613e52576000613dd66101918381548110613d5b57613d5a617a12565b5b906000526020600020906003020160020154613dc86101918581548110613d8557613d84617a12565b5b9060005260206000209060030201600101546101948f604051613da89190617410565b908152602001604051809103902060010154615b3290919063ffffffff16565b615b4890919063ffffffff16565b9050613e29846101918481548110613df157613df0617a12565b5b906000526020600020906003020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683615539565b613e3c8184615b5e90919063ffffffff16565b9250508080613e4a90617afb565b915050613d33565b50613e708289613e6b848b615b7490919063ffffffff16565b615539565b6040518060c001604052808a81526020018881526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018973ffffffffffffffffffffffffffffffffffffffff168152602001600280811115613ed257613ed1617427565b5b815260200160026004811115613eeb57613eea617427565b5b8152506101948a604051613eff9190617410565b90815260200160405180910390206000820151816000019081613f22919061765c565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160146101000a81548160ff02191690836002811115613fe657613fe5617427565b5b021790555060a08201518160030160156101000a81548160ff0219169083600481111561401657614015617427565b5b02179055509050506001935050505095945050505050565b7f5789b43a60de35bcedee40618ae90979bab7d1315fd4b079234241bdab19936d61405881614df1565b600161019260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b60008084846040516020016140cb9291906173e8565b60405160208183030381529060405280519060200120905060006140ef8285614986565b90508473ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146040518060400160405280600181526020017f370000000000000000000000000000000000000000000000000000000000000081525090614197576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161418e91906167f8565b60405180910390fd5b506001151561019260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514156040518060400160405280600281526020017f313800000000000000000000000000000000000000000000000000000000000081525090614265576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161425c91906167f8565b60405180910390fd5b50610194866040516142779190617410565b908152602001604051809103902060010154614292826123e7565b10156040518060400160405280600281526020017f31300000000000000000000000000000000000000000000000000000000000008152509061430b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161430291906167f8565b60405180910390fd5b506000805b6101918054905081101561442f5760006143b3610191838154811061433857614337617a12565b5b9060005260206000209060030201600201546143a5610191858154811061436257614361617a12565b5b9060005260206000209060030201600101546101948d6040516143859190617410565b908152602001604051809103902060010154615b3290919063ffffffff16565b615b4890919063ffffffff16565b90506144068461019184815481106143ce576143cd617a12565b5b906000526020600020906003020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683615539565b6144198184615b5e90919063ffffffff16565b925050808061442790617afb565b915050614310565b506144b182610194896040516144459190617410565b908152602001604051809103902060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166144ac846101948c60405161448c9190617410565b908152602001604051809103902060010154615b7490919063ffffffff16565b615539565b85610194886040516144c39190617410565b908152602001604051809103902060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506002610194886040516145259190617410565b908152602001604051809103902060030160156101000a81548160ff0219169083600481111561455857614557617427565b5b0217905550600193505050509392505050565b7f5789b43a60de35bcedee40618ae90979bab7d1315fd4b079234241bdab19936d81565b60006040518060c00160405280858152602001838152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020016002808111156145f4576145f3617427565b5b81526020016003600481111561460d5761460c617427565b5b815250610194856040516146219190617410565b90815260200160405180910390206000820151816000019081614644919061765c565b506020820151816001015560408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160146101000a81548160ff0219169083600281111561470857614707617427565b5b021790555060a08201518160030160156101000a81548160ff0219169083600481111561473857614737617427565b5b0217905550905050600190509392505050565b60007f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b600033905090565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff160361482c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161482390617dd1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff160361489b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161489290617e63565b60405180910390fd5b80603460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258360405161497991906168fd565b60405180910390a3505050565b60006149a38261499585615b8a565b615bba90919063ffffffff16565b905092915050565b60006149b78484613668565b90507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8114614a315781811015614a23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614a1a90617ecf565b60405180910390fd5b614a3084848484036147bd565b5b50505050565b6001151561019260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514156040518060400160405280600281526020017f313800000000000000000000000000000000000000000000000000000000000081525090614b04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614afb91906167f8565b60405180910390fd5b506001151561019260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16151514156040518060400160405280600281526020017f313800000000000000000000000000000000000000000000000000000000000081525090614bd2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614bc991906167f8565b60405180910390fd5b50600080600090505b61019180549050811015614c725760006101918281548110614c0057614bff617a12565b5b9060005260206000209060030201600201546101918381548110614c2757614c26617a12565b5b90600052602060002090600302016001015485614c449190617a41565b614c4e9190617aca565b90508083614c5c9190617913565b9250508080614c6a90617afb565b915050614bdb565b508082614c7f9190617913565b614c88856123e7565b10156040518060400160405280600281526020017f313600000000000000000000000000000000000000000000000000000000000081525090614d01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614cf891906167f8565b60405180910390fd5b5060005b61019180549050811015614ddf5760006101918281548110614d2a57614d29617a12565b5b9060005260206000209060030201600201546101918381548110614d5157614d50617a12565b5b90600052602060002090600302016001015485614d6e9190617a41565b614d789190617aca565b9050614dcb866101918481548110614d9357614d92617a12565b5b906000526020600020906003020160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683615539565b508080614dd790617afb565b915050614d05565b50614deb848484615539565b50505050565b614e0281614dfd6147b5565b615be1565b50565b614e0f828261272c565b614ee257600160fb600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550614e876147b5565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b614ef0828261272c565b15614fc457600060fb600084815260200190815260200160002060000160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550614f696147b5565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b6000614ff67f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b615c7e565b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6000801b61502c81614df1565b5050565b61505c7f4910fdfa16fed3260ed0e7147f7cc6da11a60208b5b9406d12a635614ffd914360001b615c88565b60000160009054906101000a900460ff16156150805761507b83615c92565b615199565b8273ffffffffffffffffffffffffffffffffffffffff166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156150e857506040513d601f19601f820116820180604052508101906150e59190617f04565b60015b615127576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161511e90617fa3565b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b811461518c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161518390618035565b60405180910390fd5b50615198838383615d4b565b5b505050565b6151a6615d77565b6000609760006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa6151ea6147b5565b6040516151f79190618064565b60405180910390a1565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603615270576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615267906180cb565b60405180910390fd5b61527c60008383615dc0565b806035600082825461528e9190617913565b9250508190555080603360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060008282546152e49190617913565b925050819055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161534991906168fd565b60405180910390a361535d60008383615dd8565b5050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036153d0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016153c79061815d565b60405180910390fd5b6153dc82600083615dc0565b6000603360008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002054905081811015615463576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161545a906181ef565b60405180910390fd5b818103603360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081603560008282546154bb919061820f565b92505081905550600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161552091906168fd565b60405180910390a361553483600084615dd8565b505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036155a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161559f906182b5565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603615617576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161560e90618347565b60405180910390fd5b615622838383615dc0565b6000603360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050818110156156a9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016156a0906183d9565b60405180910390fd5b818103603360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081603360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600082825461573e9190617913565b925050819055508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef846040516157a291906168fd565b60405180910390a36157b5848484615dd8565b50505050565b6000808273ffffffffffffffffffffffffffffffffffffffff163b119050919050565b600060019054906101000a900460ff1661582d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016158249061846b565b60405180910390fd5b6158378282615ddd565b5050565b600060019054906101000a900460ff1661588a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016158819061846b565b60405180910390fd5b565b600060019054906101000a900460ff166158db576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016158d29061846b565b60405180910390fd5b6158e3615e50565b565b600060019054906101000a900460ff16615934576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161592b9061846b565b60405180910390fd5b565b600060019054906101000a900460ff16615985576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161597c9061846b565b60405180910390fd5b565b600081511115615acc5760005b8151811015615aca57600060405180606001604052808484815181106159bd576159bc617a12565b5b60200260200101516000015173ffffffffffffffffffffffffffffffffffffffff1681526020018484815181106159f7576159f6617a12565b5b6020026020010151602001518152602001848481518110615a1b57615a1a617a12565b5b602002602001015160400151815250905061019181908060018154018082558091505060019003906000526020600020906003020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010155604082015181600201555050508080615ac290617afb565b915050615994565b505b50565b615ad7615ebc565b6001609760006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258615b1b6147b5565b604051615b289190618064565b60405180910390a1565b60008183615b409190617a41565b905092915050565b60008183615b569190617aca565b905092915050565b60008183615b6c9190617913565b905092915050565b60008183615b82919061820f565b905092915050565b600081604051602001615b9d91906184f8565b604051602081830303815290604052805190602001209050919050565b6000806000615bc98585615f06565b91509150615bd681615f87565b819250505092915050565b615beb828261272c565b615c7a57615c108173ffffffffffffffffffffffffffffffffffffffff166014616153565b615c1e8360001c6020616153565b604051602001615c2f9291906185b6565b6040516020818303038152906040526040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615c7191906167f8565b60405180910390fd5b5050565b6000819050919050565b6000819050919050565b615c9b816157bb565b615cda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615cd190618662565b60405180910390fd5b80615d077f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc60001b615c7e565b60000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b615d548361638f565b600082511180615d615750805b15615d7257615d7083836163de565b505b505050565b615d7f611c65565b615dbe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615db5906186ce565b60405180910390fd5b565b615dc8615ebc565b615dd38383836164c2565b505050565b505050565b600060019054906101000a900460ff16615e2c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615e239061846b565b60405180910390fd5b8160369081615e3b919061765c565b508060379081615e4b919061765c565b505050565b600060019054906101000a900460ff16615e9f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615e969061846b565b60405180910390fd5b6000609760006101000a81548160ff021916908315150217905550565b615ec4611c65565b15615f04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615efb9061873a565b60405180910390fd5b565b6000806041835103615f475760008060006020860151925060408601519150606086015160001a9050615f3b878285856164c7565b94509450505050615f80565b6040835103615f77576000806020850151915060408501519050615f6c8683836165d3565b935093505050615f80565b60006002915091505b9250929050565b60006004811115615f9b57615f9a617427565b5b816004811115615fae57615fad617427565b5b03156161505760016004811115615fc857615fc7617427565b5b816004811115615fdb57615fda617427565b5b0361601b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616012906187a6565b60405180910390fd5b6002600481111561602f5761602e617427565b5b81600481111561604257616041617427565b5b03616082576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161607990618812565b60405180910390fd5b6003600481111561609657616095617427565b5b8160048111156160a9576160a8617427565b5b036160e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016160e0906188a4565b60405180910390fd5b6004808111156160fc576160fb617427565b5b81600481111561610f5761610e617427565b5b0361614f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161614690618936565b60405180910390fd5b5b50565b6060600060028360026161669190617a41565b6161709190617913565b67ffffffffffffffff81111561618957616188616922565b5b6040519080825280601f01601f1916602001820160405280156161bb5781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106161f3576161f2617a12565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061625757616256617a12565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600060018460026162979190617a41565b6162a19190617913565b90505b6001811115616341577f3031323334353637383961626364656600000000000000000000000000000000600f8616601081106162e3576162e2617a12565b5b1a60f81b8282815181106162fa576162f9617a12565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600485901c94508061633a90618956565b90506162a4565b5060008414616385576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161637c906189cb565b60405180910390fd5b8091505092915050565b61639881615c92565b8073ffffffffffffffffffffffffffffffffffffffff167fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b60405160405180910390a250565b60606163e9836157bb565b616428576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161641f90618a5d565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff16846040516164509190618ac4565b600060405180830381855af49150503d806000811461648b576040519150601f19603f3d011682016040523d82523d6000602084013e616490565b606091505b50915091506164b88282604051806060016040528060278152602001618b2160279139616632565b9250505092915050565b505050565b6000807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08360001c11156165025760006003915091506165ca565b601b8560ff161415801561651a5750601c8560ff1614155b1561652c5760006004915091506165ca565b6000600187878787604051600081526020016040526040516165519493929190618adb565b6020604051602081039080840390855afa158015616573573d6000803e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16036165c1576000600192509250506165ca565b80600092509250505b94509492505050565b60008060007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60001b841690506000601b60ff8660001c901c6166169190617913565b9050616624878288856164c7565b935093505050935093915050565b6060831561664257829050616692565b6000835111156166555782518084602001fd5b816040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161668991906167f8565b60405180910390fd5b9392505050565b6000604051905090565b600080fd5b600080fd5b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6166e2816166ad565b81146166ed57600080fd5b50565b6000813590506166ff816166d9565b92915050565b60006020828403121561671b5761671a6166a3565b5b6000616729848285016166f0565b91505092915050565b60008115159050919050565b61674781616732565b82525050565b6000602082019050616762600083018461673e565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156167a2578082015181840152602081019050616787565b60008484015250505050565b6000601f19601f8301169050919050565b60006167ca82616768565b6167d48185616773565b93506167e4818560208601616784565b6167ed816167ae565b840191505092915050565b6000602082019050818103600083015261681281846167bf565b905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006168458261681a565b9050919050565b6168558161683a565b811461686057600080fd5b50565b6000813590506168728161684c565b92915050565b6000819050919050565b61688b81616878565b811461689657600080fd5b50565b6000813590506168a881616882565b92915050565b600080604083850312156168c5576168c46166a3565b5b60006168d385828601616863565b92505060206168e485828601616899565b9150509250929050565b6168f781616878565b82525050565b600060208201905061691260008301846168ee565b92915050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61695a826167ae565b810181811067ffffffffffffffff8211171561697957616978616922565b5b80604052505050565b600061698c616699565b90506169988282616951565b919050565b600067ffffffffffffffff8211156169b8576169b7616922565b5b6169c1826167ae565b9050602081019050919050565b82818337600083830152505050565b60006169f06169eb8461699d565b616982565b905082815260208101848484011115616a0c57616a0b61691d565b5b616a178482856169ce565b509392505050565b600082601f830112616a3457616a33616918565b5b8135616a448482602086016169dd565b91505092915050565b600067ffffffffffffffff821115616a6857616a67616922565b5b616a71826167ae565b9050602081019050919050565b6000616a91616a8c84616a4d565b616982565b905082815260208101848484011115616aad57616aac61691d565b5b616ab88482856169ce565b509392505050565b600082601f830112616ad557616ad4616918565b5b8135616ae5848260208601616a7e565b91505092915050565b600080600060608486031215616b0757616b066166a3565b5b600084013567ffffffffffffffff811115616b2557616b246166a8565b5b616b3186828701616a1f565b9350506020616b4286828701616863565b925050604084013567ffffffffffffffff811115616b6357616b626166a8565b5b616b6f86828701616ac0565b9150509250925092565b6000819050919050565b616b8c81616b79565b82525050565b6000602082019050616ba76000830184616b83565b92915050565b600080600060608486031215616bc657616bc56166a3565b5b6000616bd486828701616863565b9350506020616be586828701616863565b9250506040616bf686828701616899565b9150509250925092565b616c0981616b79565b8114616c1457600080fd5b50565b600081359050616c2681616c00565b92915050565b600060208284031215616c4257616c416166a3565b5b6000616c5084828501616c17565b91505092915050565b60008060008060808587031215616c7357616c726166a3565b5b600085013567ffffffffffffffff811115616c9157616c906166a8565b5b616c9d87828801616a1f565b9450506020616cae87828801616899565b9350506040616cbf87828801616863565b925050606085013567ffffffffffffffff811115616ce057616cdf6166a8565b5b616cec87828801616ac0565b91505092959194509250565b60008060408385031215616d0f57616d0e6166a3565b5b6000616d1d85828601616c17565b9250506020616d2e85828601616863565b9150509250929050565b600060ff82169050919050565b616d4e81616d38565b82525050565b6000602082019050616d696000830184616d45565b92915050565b600060208284031215616d8557616d846166a3565b5b6000616d9384828501616863565b91505092915050565b600060208284031215616db257616db16166a3565b5b6000616dc084828501616899565b91505092915050565b60008060008060808587031215616de357616de26166a3565b5b6000616df187828801616863565b945050602085013567ffffffffffffffff811115616e1257616e116166a8565b5b616e1e87828801616a1f565b9350506040616e2f87828801616863565b925050606085013567ffffffffffffffff811115616e5057616e4f6166a8565b5b616e5c87828801616ac0565b91505092959194509250565b60008060408385031215616e7f57616e7e6166a3565b5b6000616e8d85828601616863565b925050602083013567ffffffffffffffff811115616eae57616ead6166a8565b5b616eba85828601616ac0565b9150509250929050565b600067ffffffffffffffff821115616edf57616ede616922565b5b602082029050602081019050919050565b600080fd5b600080fd5b600060608284031215616f1057616f0f616ef5565b5b616f1a6060616982565b90506000616f2a84828501616863565b6000830152506020616f3e84828501616899565b6020830152506040616f5284828501616899565b60408301525092915050565b6000616f71616f6c84616ec4565b616982565b90508083825260208201905060608402830185811115616f9457616f93616ef0565b5b835b81811015616fbd5780616fa98882616efa565b845260208401935050606081019050616f96565b5050509392505050565b600082601f830112616fdc57616fdb616918565b5b8135616fec848260208601616f5e565b91505092915050565b60008060006060848603121561700e5761700d6166a3565b5b600084013567ffffffffffffffff81111561702c5761702b6166a8565b5b61703886828701616fc7565b935050602084013567ffffffffffffffff811115617059576170586166a8565b5b61706586828701616a1f565b925050604084013567ffffffffffffffff811115617086576170856166a8565b5b61709286828701616a1f565b9150509250925092565b600080600080608085870312156170b6576170b56166a3565b5b60006170c487828801616863565b94505060206170d587828801616899565b93505060406170e687828801616863565b925050606085013567ffffffffffffffff811115617107576171066166a8565b5b61711387828801616ac0565b91505092959194509250565b60008060408385031215617136576171356166a3565b5b600061714485828601616863565b925050602061715585828601616863565b9150509250929050565b60008060008060808587031215617179576171786166a3565b5b600061718787828801616863565b945050602061719887828801616863565b93505060406171a987828801616863565b925050606085013567ffffffffffffffff8111156171ca576171c96166a8565b5b6171d687828801616ac0565b91505092959194509250565b600080600080600060a086880312156171fe576171fd6166a3565b5b600086013567ffffffffffffffff81111561721c5761721b6166a8565b5b61722888828901616a1f565b955050602061723988828901616863565b945050604061724a88828901616899565b935050606061725b88828901616863565b925050608086013567ffffffffffffffff81111561727c5761727b6166a8565b5b61728888828901616ac0565b9150509295509295909350565b6000806000606084860312156172ae576172ad6166a3565b5b600084013567ffffffffffffffff8111156172cc576172cb6166a8565b5b6172d886828701616a1f565b93505060206172e986828701616863565b92505060406172fa86828701616899565b9150509250925092565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061734b57607f821691505b60208210810361735e5761735d617304565b5b50919050565b600081905092915050565b600061737a82616768565b6173848185617364565b9350617394818560208601616784565b80840191505092915050565b60008160601b9050919050565b60006173b8826173a0565b9050919050565b60006173ca826173ad565b9050919050565b6173e26173dd8261683a565b6173bf565b82525050565b60006173f4828561736f565b915061740082846173d1565b6014820191508190509392505050565b600061741c828461736f565b915081905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6000819050919050565b61747161746c82616878565b617456565b82525050565b6000617483828661736f565b915061748f8285617460565b60208201915061749f82846173d1565b601482019150819050949350505050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026175127fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826174d5565b61751c86836174d5565b95508019841693508086168417925050509392505050565b6000819050919050565b600061755961755461754f84616878565b617534565b616878565b9050919050565b6000819050919050565b6175738361753e565b61758761757f82617560565b8484546174e2565b825550505050565b600090565b61759c61758f565b6175a781848461756a565b505050565b5b818110156175cb576175c0600082617594565b6001810190506175ad565b5050565b601f821115617610576175e1816174b0565b6175ea846174c5565b810160208510156175f9578190505b61760d617605856174c5565b8301826175ac565b50505b505050565b600082821c905092915050565b600061763360001984600802617615565b1980831691505092915050565b600061764c8383617622565b9150826002028217905092915050565b61766582616768565b67ffffffffffffffff81111561767e5761767d616922565b5b6176888254617333565b6176938282856175cf565b600060209050601f8311600181146176c657600084156176b4578287015190505b6176be8582617640565b865550617726565b601f1984166176d4866174b0565b60005b828110156176fc578489015182556001820191506020850194506020810190506176d7565b868310156177195784890151617715601f891682617622565b8355505b6001600288020188555050505b505050505050565b7f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008201527f20726f6c657320666f722073656c660000000000000000000000000000000000602082015250565b600061778a602f83616773565b91506177958261772e565b604082019050919050565b600060208201905081810360008301526177b98161777d565b9050919050565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f64656c656761746563616c6c0000000000000000000000000000000000000000602082015250565b600061781c602c83616773565b9150617827826177c0565b604082019050919050565b6000602082019050818103600083015261784b8161780f565b9050919050565b7f46756e6374696f6e206d7573742062652063616c6c6564207468726f7567682060008201527f6163746976652070726f78790000000000000000000000000000000000000000602082015250565b60006178ae602c83616773565b91506178b982617852565b604082019050919050565b600060208201905081810360008301526178dd816178a1565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061791e82616878565b915061792983616878565b9250828201905080821115617941576179406178e4565b5b92915050565b600061795382866173d1565b601482019150617963828561736f565b915061796f82846173d1565b601482019150819050949350505050565b7f555550535570677261646561626c653a206d757374206e6f742062652063616c60008201527f6c6564207468726f7567682064656c656761746563616c6c0000000000000000602082015250565b60006179dc603883616773565b91506179e782617980565b604082019050919050565b60006020820190508181036000830152617a0b816179cf565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000617a4c82616878565b9150617a5783616878565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615617a9057617a8f6178e4565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000617ad582616878565b9150617ae083616878565b925082617af057617aef617a9b565b5b828204905092915050565b6000617b0682616878565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203617b3857617b376178e4565b5b600182019050919050565b7f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160008201527f647920696e697469616c697a6564000000000000000000000000000000000000602082015250565b6000617b9f602e83616773565b9150617baa82617b43565b604082019050919050565b60006020820190508181036000830152617bce81617b92565b9050919050565b6000819050919050565b6000617bfa617bf5617bf084617bd5565b617534565b616d38565b9050919050565b617c0a81617bdf565b82525050565b6000602082019050617c256000830184617c01565b92915050565b6000617c3782866173d1565b601482019150617c478285617460565b602082019150617c5782846173d1565b601482019150819050949350505050565b7f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f7760008201527f207a65726f000000000000000000000000000000000000000000000000000000602082015250565b6000617cc4602583616773565b9150617ccf82617c68565b604082019050919050565b60006020820190508181036000830152617cf381617cb7565b9050919050565b6000617d0682846173d1565b60148201915081905092915050565b6000617d21828761736f565b9150617d2d82866173d1565b601482019150617d3d8285617460565b602082019150617d4d82846173d1565b60148201915081905095945050505050565b7f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008201527f7265737300000000000000000000000000000000000000000000000000000000602082015250565b6000617dbb602483616773565b9150617dc682617d5f565b604082019050919050565b60006020820190508181036000830152617dea81617dae565b9050919050565b7f45524332303a20617070726f766520746f20746865207a65726f20616464726560008201527f7373000000000000000000000000000000000000000000000000000000000000602082015250565b6000617e4d602283616773565b9150617e5882617df1565b604082019050919050565b60006020820190508181036000830152617e7c81617e40565b9050919050565b7f45524332303a20696e73756666696369656e7420616c6c6f77616e6365000000600082015250565b6000617eb9601d83616773565b9150617ec482617e83565b602082019050919050565b60006020820190508181036000830152617ee881617eac565b9050919050565b600081519050617efe81616c00565b92915050565b600060208284031215617f1a57617f196166a3565b5b6000617f2884828501617eef565b91505092915050565b7f45524331393637557067726164653a206e657720696d706c656d656e7461746960008201527f6f6e206973206e6f742055555053000000000000000000000000000000000000602082015250565b6000617f8d602e83616773565b9150617f9882617f31565b604082019050919050565b60006020820190508181036000830152617fbc81617f80565b9050919050565b7f45524331393637557067726164653a20756e737570706f727465642070726f7860008201527f6961626c65555549440000000000000000000000000000000000000000000000602082015250565b600061801f602983616773565b915061802a82617fc3565b604082019050919050565b6000602082019050818103600083015261804e81618012565b9050919050565b61805e8161683a565b82525050565b60006020820190506180796000830184618055565b92915050565b7f45524332303a206d696e7420746f20746865207a65726f206164647265737300600082015250565b60006180b5601f83616773565b91506180c08261807f565b602082019050919050565b600060208201905081810360008301526180e4816180a8565b9050919050565b7f45524332303a206275726e2066726f6d20746865207a65726f2061646472657360008201527f7300000000000000000000000000000000000000000000000000000000000000602082015250565b6000618147602183616773565b9150618152826180eb565b604082019050919050565b600060208201905081810360008301526181768161813a565b9050919050565b7f45524332303a206275726e20616d6f756e7420657863656564732062616c616e60008201527f6365000000000000000000000000000000000000000000000000000000000000602082015250565b60006181d9602283616773565b91506181e48261817d565b604082019050919050565b60006020820190508181036000830152618208816181cc565b9050919050565b600061821a82616878565b915061822583616878565b925082820390508181111561823d5761823c6178e4565b5b92915050565b7f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008201527f6472657373000000000000000000000000000000000000000000000000000000602082015250565b600061829f602583616773565b91506182aa82618243565b604082019050919050565b600060208201905081810360008301526182ce81618292565b9050919050565b7f45524332303a207472616e7366657220746f20746865207a65726f206164647260008201527f6573730000000000000000000000000000000000000000000000000000000000602082015250565b6000618331602383616773565b915061833c826182d5565b604082019050919050565b6000602082019050818103600083015261836081618324565b9050919050565b7f45524332303a207472616e7366657220616d6f756e742065786365656473206260008201527f616c616e63650000000000000000000000000000000000000000000000000000602082015250565b60006183c3602683616773565b91506183ce82618367565b604082019050919050565b600060208201905081810360008301526183f2816183b6565b9050919050565b7f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960008201527f6e697469616c697a696e67000000000000000000000000000000000000000000602082015250565b6000618455602b83616773565b9150618460826183f9565b604082019050919050565b6000602082019050818103600083015261848481618448565b9050919050565b7f19457468657265756d205369676e6564204d6573736167653a0a333200000000600082015250565b60006184c1601c83617364565b91506184cc8261848b565b601c82019050919050565b6000819050919050565b6184f26184ed82616b79565b6184d7565b82525050565b6000618503826184b4565b915061850f82846184e1565b60208201915081905092915050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000600082015250565b6000618554601783617364565b915061855f8261851e565b601782019050919050565b7f206973206d697373696e6720726f6c6520000000000000000000000000000000600082015250565b60006185a0601183617364565b91506185ab8261856a565b601182019050919050565b60006185c182618547565b91506185cd828561736f565b91506185d882618593565b91506185e4828461736f565b91508190509392505050565b7f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60008201527f6f74206120636f6e747261637400000000000000000000000000000000000000602082015250565b600061864c602d83616773565b9150618657826185f0565b604082019050919050565b6000602082019050818103600083015261867b8161863f565b9050919050565b7f5061757361626c653a206e6f7420706175736564000000000000000000000000600082015250565b60006186b8601483616773565b91506186c382618682565b602082019050919050565b600060208201905081810360008301526186e7816186ab565b9050919050565b7f5061757361626c653a2070617573656400000000000000000000000000000000600082015250565b6000618724601083616773565b915061872f826186ee565b602082019050919050565b6000602082019050818103600083015261875381618717565b9050919050565b7f45434453413a20696e76616c6964207369676e61747572650000000000000000600082015250565b6000618790601883616773565b915061879b8261875a565b602082019050919050565b600060208201905081810360008301526187bf81618783565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265206c656e67746800600082015250565b60006187fc601f83616773565b9150618807826187c6565b602082019050919050565b6000602082019050818103600083015261882b816187ef565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b600061888e602283616773565b915061889982618832565b604082019050919050565b600060208201905081810360008301526188bd81618881565b9050919050565b7f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008201527f7565000000000000000000000000000000000000000000000000000000000000602082015250565b6000618920602283616773565b915061892b826188c4565b604082019050919050565b6000602082019050818103600083015261894f81618913565b9050919050565b600061896182616878565b915060008203618974576189736178e4565b5b600182039050919050565b7f537472696e67733a20686578206c656e67746820696e73756666696369656e74600082015250565b60006189b5602083616773565b91506189c08261897f565b602082019050919050565b600060208201905081810360008301526189e4816189a8565b9050919050565b7f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f60008201527f6e74726163740000000000000000000000000000000000000000000000000000602082015250565b6000618a47602683616773565b9150618a52826189eb565b604082019050919050565b60006020820190508181036000830152618a7681618a3a565b9050919050565b600081519050919050565b600081905092915050565b6000618a9e82618a7d565b618aa88185618a88565b9350618ab8818560208601616784565b80840191505092915050565b6000618ad08284618a93565b915081905092915050565b6000608082019050618af06000830187616b83565b618afd6020830186616d45565b618b0a6040830185616b83565b618b176060830184616b83565b9594505050505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220579ef8e5f16ae5378b2c8f891d2f644cdc21a109f676a53620411541abfae7bb64736f6c63430008100033",
}

// FcfaABI is the input ABI used to generate the binding from.
// Deprecated: Use FcfaMetaData.ABI instead.
var FcfaABI = FcfaMetaData.ABI

// FcfaBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FcfaMetaData.Bin instead.
var FcfaBin = FcfaMetaData.Bin

// DeployFcfa deploys a new Ethereum contract, binding an instance of Fcfa to it.
func DeployFcfa(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Fcfa, error) {
	parsed, err := FcfaMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FcfaBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Fcfa{FcfaCaller: FcfaCaller{contract: contract}, FcfaTransactor: FcfaTransactor{contract: contract}, FcfaFilterer: FcfaFilterer{contract: contract}}, nil
}

// Fcfa is an auto generated Go binding around an Ethereum contract.
type Fcfa struct {
	FcfaCaller     // Read-only binding to the contract
	FcfaTransactor // Write-only binding to the contract
	FcfaFilterer   // Log filterer for contract events
}

// FcfaCaller is an auto generated read-only Go binding around an Ethereum contract.
type FcfaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FcfaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FcfaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FcfaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FcfaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FcfaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FcfaSession struct {
	Contract     *Fcfa             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FcfaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FcfaCallerSession struct {
	Contract *FcfaCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// FcfaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FcfaTransactorSession struct {
	Contract     *FcfaTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FcfaRaw is an auto generated low-level Go binding around an Ethereum contract.
type FcfaRaw struct {
	Contract *Fcfa // Generic contract binding to access the raw methods on
}

// FcfaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FcfaCallerRaw struct {
	Contract *FcfaCaller // Generic read-only contract binding to access the raw methods on
}

// FcfaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FcfaTransactorRaw struct {
	Contract *FcfaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFcfa creates a new instance of Fcfa, bound to a specific deployed contract.
func NewFcfa(address common.Address, backend bind.ContractBackend) (*Fcfa, error) {
	contract, err := bindFcfa(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fcfa{FcfaCaller: FcfaCaller{contract: contract}, FcfaTransactor: FcfaTransactor{contract: contract}, FcfaFilterer: FcfaFilterer{contract: contract}}, nil
}

// NewFcfaCaller creates a new read-only instance of Fcfa, bound to a specific deployed contract.
func NewFcfaCaller(address common.Address, caller bind.ContractCaller) (*FcfaCaller, error) {
	contract, err := bindFcfa(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FcfaCaller{contract: contract}, nil
}

// NewFcfaTransactor creates a new write-only instance of Fcfa, bound to a specific deployed contract.
func NewFcfaTransactor(address common.Address, transactor bind.ContractTransactor) (*FcfaTransactor, error) {
	contract, err := bindFcfa(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FcfaTransactor{contract: contract}, nil
}

// NewFcfaFilterer creates a new log filterer instance of Fcfa, bound to a specific deployed contract.
func NewFcfaFilterer(address common.Address, filterer bind.ContractFilterer) (*FcfaFilterer, error) {
	contract, err := bindFcfa(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FcfaFilterer{contract: contract}, nil
}

// bindFcfa binds a generic wrapper to an already deployed contract.
func bindFcfa(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FcfaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fcfa *FcfaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fcfa.Contract.FcfaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fcfa *FcfaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fcfa.Contract.FcfaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fcfa *FcfaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fcfa.Contract.FcfaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fcfa *FcfaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fcfa.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fcfa *FcfaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fcfa.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fcfa *FcfaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fcfa.Contract.contract.Transact(opts, method, params...)
}

// AGENTROLE is a free data retrieval call binding the contract method 0x22459e18.
//
// Solidity: function AGENT_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCaller) AGENTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "AGENT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AGENTROLE is a free data retrieval call binding the contract method 0x22459e18.
//
// Solidity: function AGENT_ROLE() view returns(bytes32)
func (_Fcfa *FcfaSession) AGENTROLE() ([32]byte, error) {
	return _Fcfa.Contract.AGENTROLE(&_Fcfa.CallOpts)
}

// AGENTROLE is a free data retrieval call binding the contract method 0x22459e18.
//
// Solidity: function AGENT_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCallerSession) AGENTROLE() ([32]byte, error) {
	return _Fcfa.Contract.AGENTROLE(&_Fcfa.CallOpts)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_Fcfa *FcfaSession) BURNERROLE() ([32]byte, error) {
	return _Fcfa.Contract.BURNERROLE(&_Fcfa.CallOpts)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCallerSession) BURNERROLE() ([32]byte, error) {
	return _Fcfa.Contract.BURNERROLE(&_Fcfa.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Fcfa *FcfaSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Fcfa.Contract.DEFAULTADMINROLE(&_Fcfa.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Fcfa.Contract.DEFAULTADMINROLE(&_Fcfa.CallOpts)
}

// FREEZEROLE is a free data retrieval call binding the contract method 0xf9e85723.
//
// Solidity: function FREEZE_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCaller) FREEZEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "FREEZE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FREEZEROLE is a free data retrieval call binding the contract method 0xf9e85723.
//
// Solidity: function FREEZE_ROLE() view returns(bytes32)
func (_Fcfa *FcfaSession) FREEZEROLE() ([32]byte, error) {
	return _Fcfa.Contract.FREEZEROLE(&_Fcfa.CallOpts)
}

// FREEZEROLE is a free data retrieval call binding the contract method 0xf9e85723.
//
// Solidity: function FREEZE_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCallerSession) FREEZEROLE() ([32]byte, error) {
	return _Fcfa.Contract.FREEZEROLE(&_Fcfa.CallOpts)
}

// KYCAGENT is a free data retrieval call binding the contract method 0x3bcb220e.
//
// Solidity: function KYC_AGENT() view returns(bytes32)
func (_Fcfa *FcfaCaller) KYCAGENT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "KYC_AGENT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KYCAGENT is a free data retrieval call binding the contract method 0x3bcb220e.
//
// Solidity: function KYC_AGENT() view returns(bytes32)
func (_Fcfa *FcfaSession) KYCAGENT() ([32]byte, error) {
	return _Fcfa.Contract.KYCAGENT(&_Fcfa.CallOpts)
}

// KYCAGENT is a free data retrieval call binding the contract method 0x3bcb220e.
//
// Solidity: function KYC_AGENT() view returns(bytes32)
func (_Fcfa *FcfaCallerSession) KYCAGENT() ([32]byte, error) {
	return _Fcfa.Contract.KYCAGENT(&_Fcfa.CallOpts)
}

// MANAGMENTROLE is a free data retrieval call binding the contract method 0xbf376af6.
//
// Solidity: function MANAGMENT_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCaller) MANAGMENTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "MANAGMENT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGMENTROLE is a free data retrieval call binding the contract method 0xbf376af6.
//
// Solidity: function MANAGMENT_ROLE() view returns(bytes32)
func (_Fcfa *FcfaSession) MANAGMENTROLE() ([32]byte, error) {
	return _Fcfa.Contract.MANAGMENTROLE(&_Fcfa.CallOpts)
}

// MANAGMENTROLE is a free data retrieval call binding the contract method 0xbf376af6.
//
// Solidity: function MANAGMENT_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCallerSession) MANAGMENTROLE() ([32]byte, error) {
	return _Fcfa.Contract.MANAGMENTROLE(&_Fcfa.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Fcfa *FcfaSession) MINTERROLE() ([32]byte, error) {
	return _Fcfa.Contract.MINTERROLE(&_Fcfa.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCallerSession) MINTERROLE() ([32]byte, error) {
	return _Fcfa.Contract.MINTERROLE(&_Fcfa.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Fcfa *FcfaSession) PAUSERROLE() ([32]byte, error) {
	return _Fcfa.Contract.PAUSERROLE(&_Fcfa.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Fcfa *FcfaCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Fcfa.Contract.PAUSERROLE(&_Fcfa.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Fcfa *FcfaCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Fcfa *FcfaSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Fcfa.Contract.Allowance(&_Fcfa.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Fcfa *FcfaCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Fcfa.Contract.Allowance(&_Fcfa.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Fcfa *FcfaCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Fcfa *FcfaSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Fcfa.Contract.BalanceOf(&_Fcfa.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Fcfa *FcfaCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Fcfa.Contract.BalanceOf(&_Fcfa.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Fcfa *FcfaCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Fcfa *FcfaSession) Decimals() (uint8, error) {
	return _Fcfa.Contract.Decimals(&_Fcfa.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Fcfa *FcfaCallerSession) Decimals() (uint8, error) {
	return _Fcfa.Contract.Decimals(&_Fcfa.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Fcfa *FcfaCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Fcfa *FcfaSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Fcfa.Contract.GetRoleAdmin(&_Fcfa.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Fcfa *FcfaCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Fcfa.Contract.GetRoleAdmin(&_Fcfa.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Fcfa *FcfaCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Fcfa *FcfaSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Fcfa.Contract.HasRole(&_Fcfa.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Fcfa *FcfaCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Fcfa.Contract.HasRole(&_Fcfa.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fcfa *FcfaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fcfa *FcfaSession) Name() (string, error) {
	return _Fcfa.Contract.Name(&_Fcfa.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Fcfa *FcfaCallerSession) Name() (string, error) {
	return _Fcfa.Contract.Name(&_Fcfa.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Fcfa *FcfaCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Fcfa *FcfaSession) Paused() (bool, error) {
	return _Fcfa.Contract.Paused(&_Fcfa.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Fcfa *FcfaCallerSession) Paused() (bool, error) {
	return _Fcfa.Contract.Paused(&_Fcfa.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fcfa *FcfaCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fcfa *FcfaSession) ProxiableUUID() ([32]byte, error) {
	return _Fcfa.Contract.ProxiableUUID(&_Fcfa.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Fcfa *FcfaCallerSession) ProxiableUUID() ([32]byte, error) {
	return _Fcfa.Contract.ProxiableUUID(&_Fcfa.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Fcfa *FcfaCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Fcfa *FcfaSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Fcfa.Contract.SupportsInterface(&_Fcfa.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Fcfa *FcfaCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Fcfa.Contract.SupportsInterface(&_Fcfa.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fcfa *FcfaCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fcfa *FcfaSession) Symbol() (string, error) {
	return _Fcfa.Contract.Symbol(&_Fcfa.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Fcfa *FcfaCallerSession) Symbol() (string, error) {
	return _Fcfa.Contract.Symbol(&_Fcfa.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Fcfa *FcfaCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fcfa.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Fcfa *FcfaSession) TotalSupply() (*big.Int, error) {
	return _Fcfa.Contract.TotalSupply(&_Fcfa.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Fcfa *FcfaCallerSession) TotalSupply() (*big.Int, error) {
	return _Fcfa.Contract.TotalSupply(&_Fcfa.CallOpts)
}

// AddUnconfirmedPayment is a paid mutator transaction binding the contract method 0xfe82df15.
//
// Solidity: function addUnconfirmedPayment(string transactionId, address to, uint256 amount) returns(bool)
func (_Fcfa *FcfaTransactor) AddUnconfirmedPayment(opts *bind.TransactOpts, transactionId string, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "addUnconfirmedPayment", transactionId, to, amount)
}

// AddUnconfirmedPayment is a paid mutator transaction binding the contract method 0xfe82df15.
//
// Solidity: function addUnconfirmedPayment(string transactionId, address to, uint256 amount) returns(bool)
func (_Fcfa *FcfaSession) AddUnconfirmedPayment(transactionId string, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.AddUnconfirmedPayment(&_Fcfa.TransactOpts, transactionId, to, amount)
}

// AddUnconfirmedPayment is a paid mutator transaction binding the contract method 0xfe82df15.
//
// Solidity: function addUnconfirmedPayment(string transactionId, address to, uint256 amount) returns(bool)
func (_Fcfa *FcfaTransactorSession) AddUnconfirmedPayment(transactionId string, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.AddUnconfirmedPayment(&_Fcfa.TransactOpts, transactionId, to, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Fcfa *FcfaTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Fcfa *FcfaSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.Approve(&_Fcfa.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Fcfa *FcfaTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.Approve(&_Fcfa.TransactOpts, spender, amount)
}

// ApproveDeposit is a paid mutator transaction binding the contract method 0xc3d017d6.
//
// Solidity: function approveDeposit(string transactionnId, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaTransactor) ApproveDeposit(opts *bind.TransactOpts, transactionnId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "approveDeposit", transactionnId, signer, signature)
}

// ApproveDeposit is a paid mutator transaction binding the contract method 0xc3d017d6.
//
// Solidity: function approveDeposit(string transactionnId, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaSession) ApproveDeposit(transactionnId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.ApproveDeposit(&_Fcfa.TransactOpts, transactionnId, signer, signature)
}

// ApproveDeposit is a paid mutator transaction binding the contract method 0xc3d017d6.
//
// Solidity: function approveDeposit(string transactionnId, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaTransactorSession) ApproveDeposit(transactionnId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.ApproveDeposit(&_Fcfa.TransactOpts, transactionnId, signer, signature)
}

// ApproveWithdraw is a paid mutator transaction binding the contract method 0xc087abdb.
//
// Solidity: function approveWithdraw(string transactionnId, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactor) ApproveWithdraw(opts *bind.TransactOpts, transactionnId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "approveWithdraw", transactionnId, signer, signature)
}

// ApproveWithdraw is a paid mutator transaction binding the contract method 0xc087abdb.
//
// Solidity: function approveWithdraw(string transactionnId, address signer, bytes signature) returns()
func (_Fcfa *FcfaSession) ApproveWithdraw(transactionnId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.ApproveWithdraw(&_Fcfa.TransactOpts, transactionnId, signer, signature)
}

// ApproveWithdraw is a paid mutator transaction binding the contract method 0xc087abdb.
//
// Solidity: function approveWithdraw(string transactionnId, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactorSession) ApproveWithdraw(transactionnId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.ApproveWithdraw(&_Fcfa.TransactOpts, transactionnId, signer, signature)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Fcfa *FcfaTransactor) Burn(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "burn", amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Fcfa *FcfaSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.Burn(&_Fcfa.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 amount) returns()
func (_Fcfa *FcfaTransactorSession) Burn(amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.Burn(&_Fcfa.TransactOpts, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Fcfa *FcfaTransactor) BurnFrom(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "burnFrom", account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Fcfa *FcfaSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.BurnFrom(&_Fcfa.TransactOpts, account, amount)
}

// BurnFrom is a paid mutator transaction binding the contract method 0x79cc6790.
//
// Solidity: function burnFrom(address account, uint256 amount) returns()
func (_Fcfa *FcfaTransactorSession) BurnFrom(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.BurnFrom(&_Fcfa.TransactOpts, account, amount)
}

// BurnTo is a paid mutator transaction binding the contract method 0xea785a5e.
//
// Solidity: function burnTo(address from, uint256 amount) returns()
func (_Fcfa *FcfaTransactor) BurnTo(opts *bind.TransactOpts, from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "burnTo", from, amount)
}

// BurnTo is a paid mutator transaction binding the contract method 0xea785a5e.
//
// Solidity: function burnTo(address from, uint256 amount) returns()
func (_Fcfa *FcfaSession) BurnTo(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.BurnTo(&_Fcfa.TransactOpts, from, amount)
}

// BurnTo is a paid mutator transaction binding the contract method 0xea785a5e.
//
// Solidity: function burnTo(address from, uint256 amount) returns()
func (_Fcfa *FcfaTransactorSession) BurnTo(from common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.BurnTo(&_Fcfa.TransactOpts, from, amount)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x854ed649.
//
// Solidity: function cancelDeposit(string transactionnId, address canceler, bytes signature) returns()
func (_Fcfa *FcfaTransactor) CancelDeposit(opts *bind.TransactOpts, transactionnId string, canceler common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "cancelDeposit", transactionnId, canceler, signature)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x854ed649.
//
// Solidity: function cancelDeposit(string transactionnId, address canceler, bytes signature) returns()
func (_Fcfa *FcfaSession) CancelDeposit(transactionnId string, canceler common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.CancelDeposit(&_Fcfa.TransactOpts, transactionnId, canceler, signature)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x854ed649.
//
// Solidity: function cancelDeposit(string transactionnId, address canceler, bytes signature) returns()
func (_Fcfa *FcfaTransactorSession) CancelDeposit(transactionnId string, canceler common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.CancelDeposit(&_Fcfa.TransactOpts, transactionnId, canceler, signature)
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0x18449c23.
//
// Solidity: function cancelWithdraw(string transactionnId, address canceler, bytes signature) returns()
func (_Fcfa *FcfaTransactor) CancelWithdraw(opts *bind.TransactOpts, transactionnId string, canceler common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "cancelWithdraw", transactionnId, canceler, signature)
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0x18449c23.
//
// Solidity: function cancelWithdraw(string transactionnId, address canceler, bytes signature) returns()
func (_Fcfa *FcfaSession) CancelWithdraw(transactionnId string, canceler common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.CancelWithdraw(&_Fcfa.TransactOpts, transactionnId, canceler, signature)
}

// CancelWithdraw is a paid mutator transaction binding the contract method 0x18449c23.
//
// Solidity: function cancelWithdraw(string transactionnId, address canceler, bytes signature) returns()
func (_Fcfa *FcfaTransactorSession) CancelWithdraw(transactionnId string, canceler common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.CancelWithdraw(&_Fcfa.TransactOpts, transactionnId, canceler, signature)
}

// Commercepay is a paid mutator transaction binding the contract method 0xf19b7df0.
//
// Solidity: function commercepay(string transactionId, address to, uint256 amount, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaTransactor) Commercepay(opts *bind.TransactOpts, transactionId string, to common.Address, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "commercepay", transactionId, to, amount, signer, signature)
}

// Commercepay is a paid mutator transaction binding the contract method 0xf19b7df0.
//
// Solidity: function commercepay(string transactionId, address to, uint256 amount, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaSession) Commercepay(transactionId string, to common.Address, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.Commercepay(&_Fcfa.TransactOpts, transactionId, to, amount, signer, signature)
}

// Commercepay is a paid mutator transaction binding the contract method 0xf19b7df0.
//
// Solidity: function commercepay(string transactionId, address to, uint256 amount, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaTransactorSession) Commercepay(transactionId string, to common.Address, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.Commercepay(&_Fcfa.TransactOpts, transactionId, to, amount, signer, signature)
}

// ConfirmUnconfirmedPayment is a paid mutator transaction binding the contract method 0xf9613bb4.
//
// Solidity: function confirmUnconfirmedPayment(string transactionId, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaTransactor) ConfirmUnconfirmedPayment(opts *bind.TransactOpts, transactionId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "confirmUnconfirmedPayment", transactionId, signer, signature)
}

// ConfirmUnconfirmedPayment is a paid mutator transaction binding the contract method 0xf9613bb4.
//
// Solidity: function confirmUnconfirmedPayment(string transactionId, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaSession) ConfirmUnconfirmedPayment(transactionId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.ConfirmUnconfirmedPayment(&_Fcfa.TransactOpts, transactionId, signer, signature)
}

// ConfirmUnconfirmedPayment is a paid mutator transaction binding the contract method 0xf9613bb4.
//
// Solidity: function confirmUnconfirmedPayment(string transactionId, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaTransactorSession) ConfirmUnconfirmedPayment(transactionId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.ConfirmUnconfirmedPayment(&_Fcfa.TransactOpts, transactionId, signer, signature)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Fcfa *FcfaTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Fcfa *FcfaSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.DecreaseAllowance(&_Fcfa.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Fcfa *FcfaTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.DecreaseAllowance(&_Fcfa.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x9913c7bd.
//
// Solidity: function deposit(string transactionnId, uint256 amount, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactor) Deposit(opts *bind.TransactOpts, transactionnId string, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "deposit", transactionnId, amount, signer, signature)
}

// Deposit is a paid mutator transaction binding the contract method 0x9913c7bd.
//
// Solidity: function deposit(string transactionnId, uint256 amount, address signer, bytes signature) returns()
func (_Fcfa *FcfaSession) Deposit(transactionnId string, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.Deposit(&_Fcfa.TransactOpts, transactionnId, amount, signer, signature)
}

// Deposit is a paid mutator transaction binding the contract method 0x9913c7bd.
//
// Solidity: function deposit(string transactionnId, uint256 amount, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactorSession) Deposit(transactionnId string, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.Deposit(&_Fcfa.TransactOpts, transactionnId, amount, signer, signature)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xf26c159f.
//
// Solidity: function freezeAccount(address account) returns()
func (_Fcfa *FcfaTransactor) FreezeAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "freezeAccount", account)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xf26c159f.
//
// Solidity: function freezeAccount(address account) returns()
func (_Fcfa *FcfaSession) FreezeAccount(account common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.FreezeAccount(&_Fcfa.TransactOpts, account)
}

// FreezeAccount is a paid mutator transaction binding the contract method 0xf26c159f.
//
// Solidity: function freezeAccount(address account) returns()
func (_Fcfa *FcfaTransactorSession) FreezeAccount(account common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.FreezeAccount(&_Fcfa.TransactOpts, account)
}

// GiveRoleToAddress is a paid mutator transaction binding the contract method 0x4e09b957.
//
// Solidity: function giveRoleToAddress(address account, string role, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactor) GiveRoleToAddress(opts *bind.TransactOpts, account common.Address, role string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "giveRoleToAddress", account, role, signer, signature)
}

// GiveRoleToAddress is a paid mutator transaction binding the contract method 0x4e09b957.
//
// Solidity: function giveRoleToAddress(address account, string role, address signer, bytes signature) returns()
func (_Fcfa *FcfaSession) GiveRoleToAddress(account common.Address, role string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.GiveRoleToAddress(&_Fcfa.TransactOpts, account, role, signer, signature)
}

// GiveRoleToAddress is a paid mutator transaction binding the contract method 0x4e09b957.
//
// Solidity: function giveRoleToAddress(address account, string role, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactorSession) GiveRoleToAddress(account common.Address, role string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.GiveRoleToAddress(&_Fcfa.TransactOpts, account, role, signer, signature)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Fcfa *FcfaTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Fcfa *FcfaSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.GrantRole(&_Fcfa.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Fcfa *FcfaTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.GrantRole(&_Fcfa.TransactOpts, role, account)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Fcfa *FcfaTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Fcfa *FcfaSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.IncreaseAllowance(&_Fcfa.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Fcfa *FcfaTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.IncreaseAllowance(&_Fcfa.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x66e2cd86.
//
// Solidity: function initialize((address,uint256,uint256)[] _feeCollectors, string name, string symbol) returns()
func (_Fcfa *FcfaTransactor) Initialize(opts *bind.TransactOpts, _feeCollectors []FeeCollector, name string, symbol string) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "initialize", _feeCollectors, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x66e2cd86.
//
// Solidity: function initialize((address,uint256,uint256)[] _feeCollectors, string name, string symbol) returns()
func (_Fcfa *FcfaSession) Initialize(_feeCollectors []FeeCollector, name string, symbol string) (*types.Transaction, error) {
	return _Fcfa.Contract.Initialize(&_Fcfa.TransactOpts, _feeCollectors, name, symbol)
}

// Initialize is a paid mutator transaction binding the contract method 0x66e2cd86.
//
// Solidity: function initialize((address,uint256,uint256)[] _feeCollectors, string name, string symbol) returns()
func (_Fcfa *FcfaTransactorSession) Initialize(_feeCollectors []FeeCollector, name string, symbol string) (*types.Transaction, error) {
	return _Fcfa.Contract.Initialize(&_Fcfa.TransactOpts, _feeCollectors, name, symbol)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Fcfa *FcfaTransactor) Mint(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "mint", to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Fcfa *FcfaSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.Mint(&_Fcfa.TransactOpts, to, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 amount) returns()
func (_Fcfa *FcfaTransactorSession) Mint(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.Mint(&_Fcfa.TransactOpts, to, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Fcfa *FcfaTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Fcfa *FcfaSession) Pause() (*types.Transaction, error) {
	return _Fcfa.Contract.Pause(&_Fcfa.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Fcfa *FcfaTransactorSession) Pause() (*types.Transaction, error) {
	return _Fcfa.Contract.Pause(&_Fcfa.TransactOpts)
}

// Refund is a paid mutator transaction binding the contract method 0x5e19e7b9.
//
// Solidity: function refund(string transactionId, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaTransactor) Refund(opts *bind.TransactOpts, transactionId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "refund", transactionId, signer, signature)
}

// Refund is a paid mutator transaction binding the contract method 0x5e19e7b9.
//
// Solidity: function refund(string transactionId, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaSession) Refund(transactionId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.Refund(&_Fcfa.TransactOpts, transactionId, signer, signature)
}

// Refund is a paid mutator transaction binding the contract method 0x5e19e7b9.
//
// Solidity: function refund(string transactionId, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaTransactorSession) Refund(transactionId string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.Refund(&_Fcfa.TransactOpts, transactionId, signer, signature)
}

// RemoveKycFlag is a paid mutator transaction binding the contract method 0xe1825cf7.
//
// Solidity: function removeKycFlag(address account, address remover, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactor) RemoveKycFlag(opts *bind.TransactOpts, account common.Address, remover common.Address, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "removeKycFlag", account, remover, signer, signature)
}

// RemoveKycFlag is a paid mutator transaction binding the contract method 0xe1825cf7.
//
// Solidity: function removeKycFlag(address account, address remover, address signer, bytes signature) returns()
func (_Fcfa *FcfaSession) RemoveKycFlag(account common.Address, remover common.Address, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.RemoveKycFlag(&_Fcfa.TransactOpts, account, remover, signer, signature)
}

// RemoveKycFlag is a paid mutator transaction binding the contract method 0xe1825cf7.
//
// Solidity: function removeKycFlag(address account, address remover, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactorSession) RemoveKycFlag(account common.Address, remover common.Address, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.RemoveKycFlag(&_Fcfa.TransactOpts, account, remover, signer, signature)
}

// RemoveroleFromAddress is a paid mutator transaction binding the contract method 0xb90df18c.
//
// Solidity: function removeroleFromAddress(address account, string role, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactor) RemoveroleFromAddress(opts *bind.TransactOpts, account common.Address, role string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "removeroleFromAddress", account, role, signer, signature)
}

// RemoveroleFromAddress is a paid mutator transaction binding the contract method 0xb90df18c.
//
// Solidity: function removeroleFromAddress(address account, string role, address signer, bytes signature) returns()
func (_Fcfa *FcfaSession) RemoveroleFromAddress(account common.Address, role string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.RemoveroleFromAddress(&_Fcfa.TransactOpts, account, role, signer, signature)
}

// RemoveroleFromAddress is a paid mutator transaction binding the contract method 0xb90df18c.
//
// Solidity: function removeroleFromAddress(address account, string role, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactorSession) RemoveroleFromAddress(account common.Address, role string, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.RemoveroleFromAddress(&_Fcfa.TransactOpts, account, role, signer, signature)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Fcfa *FcfaTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Fcfa *FcfaSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.RenounceRole(&_Fcfa.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Fcfa *FcfaTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.RenounceRole(&_Fcfa.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Fcfa *FcfaTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Fcfa *FcfaSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.RevokeRole(&_Fcfa.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Fcfa *FcfaTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.RevokeRole(&_Fcfa.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0x92b4dfd2.
//
// Solidity: function transfer(address to, uint256 amount, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "transfer", to, amount, signer, signature)
}

// Transfer is a paid mutator transaction binding the contract method 0x92b4dfd2.
//
// Solidity: function transfer(address to, uint256 amount, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaSession) Transfer(to common.Address, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.Transfer(&_Fcfa.TransactOpts, to, amount, signer, signature)
}

// Transfer is a paid mutator transaction binding the contract method 0x92b4dfd2.
//
// Solidity: function transfer(address to, uint256 amount, address signer, bytes signature) returns(bool)
func (_Fcfa *FcfaTransactorSession) Transfer(to common.Address, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.Transfer(&_Fcfa.TransactOpts, to, amount, signer, signature)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Fcfa *FcfaTransactor) Transfer0(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "transfer0", to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Fcfa *FcfaSession) Transfer0(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.Transfer0(&_Fcfa.TransactOpts, to, amount)
}

// Transfer0 is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Fcfa *FcfaTransactorSession) Transfer0(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.Transfer0(&_Fcfa.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Fcfa *FcfaTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Fcfa *FcfaSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.TransferFrom(&_Fcfa.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Fcfa *FcfaTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fcfa.Contract.TransferFrom(&_Fcfa.TransactOpts, from, to, amount)
}

// UnFreezeAccount is a paid mutator transaction binding the contract method 0x53cc2fae.
//
// Solidity: function unFreezeAccount(address account) returns()
func (_Fcfa *FcfaTransactor) UnFreezeAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "unFreezeAccount", account)
}

// UnFreezeAccount is a paid mutator transaction binding the contract method 0x53cc2fae.
//
// Solidity: function unFreezeAccount(address account) returns()
func (_Fcfa *FcfaSession) UnFreezeAccount(account common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.UnFreezeAccount(&_Fcfa.TransactOpts, account)
}

// UnFreezeAccount is a paid mutator transaction binding the contract method 0x53cc2fae.
//
// Solidity: function unFreezeAccount(address account) returns()
func (_Fcfa *FcfaTransactorSession) UnFreezeAccount(account common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.UnFreezeAccount(&_Fcfa.TransactOpts, account)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Fcfa *FcfaTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Fcfa *FcfaSession) Unpause() (*types.Transaction, error) {
	return _Fcfa.Contract.Unpause(&_Fcfa.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Fcfa *FcfaTransactorSession) Unpause() (*types.Transaction, error) {
	return _Fcfa.Contract.Unpause(&_Fcfa.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Fcfa *FcfaTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Fcfa *FcfaSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.UpgradeTo(&_Fcfa.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Fcfa *FcfaTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Fcfa.Contract.UpgradeTo(&_Fcfa.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Fcfa *FcfaTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Fcfa *FcfaSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.UpgradeToAndCall(&_Fcfa.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Fcfa *FcfaTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.UpgradeToAndCall(&_Fcfa.TransactOpts, newImplementation, data)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2d7cdeb1.
//
// Solidity: function withdraw(string transactionnId, uint256 amount, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactor) Withdraw(opts *bind.TransactOpts, transactionnId string, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.contract.Transact(opts, "withdraw", transactionnId, amount, signer, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2d7cdeb1.
//
// Solidity: function withdraw(string transactionnId, uint256 amount, address signer, bytes signature) returns()
func (_Fcfa *FcfaSession) Withdraw(transactionnId string, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.Withdraw(&_Fcfa.TransactOpts, transactionnId, amount, signer, signature)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2d7cdeb1.
//
// Solidity: function withdraw(string transactionnId, uint256 amount, address signer, bytes signature) returns()
func (_Fcfa *FcfaTransactorSession) Withdraw(transactionnId string, amount *big.Int, signer common.Address, signature []byte) (*types.Transaction, error) {
	return _Fcfa.Contract.Withdraw(&_Fcfa.TransactOpts, transactionnId, amount, signer, signature)
}

// FcfaAddKYCFlagIterator is returned from FilterAddKYCFlag and is used to iterate over the raw logs and unpacked data for AddKYCFlag events raised by the Fcfa contract.
type FcfaAddKYCFlagIterator struct {
	Event *FcfaAddKYCFlag // Event containing the contract specifics and raw log

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
func (it *FcfaAddKYCFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaAddKYCFlag)
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
		it.Event = new(FcfaAddKYCFlag)
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
func (it *FcfaAddKYCFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaAddKYCFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaAddKYCFlag represents a AddKYCFlag event raised by the Fcfa contract.
type FcfaAddKYCFlag struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAddKYCFlag is a free log retrieval operation binding the contract event 0xf3cedc1ecace48692d585b4713ace260c3001985c589e976c06bb45050c2dbc5.
//
// Solidity: event AddKYCFlag(address indexed account)
func (_Fcfa *FcfaFilterer) FilterAddKYCFlag(opts *bind.FilterOpts, account []common.Address) (*FcfaAddKYCFlagIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "AddKYCFlag", accountRule)
	if err != nil {
		return nil, err
	}
	return &FcfaAddKYCFlagIterator{contract: _Fcfa.contract, event: "AddKYCFlag", logs: logs, sub: sub}, nil
}

// WatchAddKYCFlag is a free log subscription operation binding the contract event 0xf3cedc1ecace48692d585b4713ace260c3001985c589e976c06bb45050c2dbc5.
//
// Solidity: event AddKYCFlag(address indexed account)
func (_Fcfa *FcfaFilterer) WatchAddKYCFlag(opts *bind.WatchOpts, sink chan<- *FcfaAddKYCFlag, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "AddKYCFlag", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaAddKYCFlag)
				if err := _Fcfa.contract.UnpackLog(event, "AddKYCFlag", log); err != nil {
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

// ParseAddKYCFlag is a log parse operation binding the contract event 0xf3cedc1ecace48692d585b4713ace260c3001985c589e976c06bb45050c2dbc5.
//
// Solidity: event AddKYCFlag(address indexed account)
func (_Fcfa *FcfaFilterer) ParseAddKYCFlag(log types.Log) (*FcfaAddKYCFlag, error) {
	event := new(FcfaAddKYCFlag)
	if err := _Fcfa.contract.UnpackLog(event, "AddKYCFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Fcfa contract.
type FcfaAdminChangedIterator struct {
	Event *FcfaAdminChanged // Event containing the contract specifics and raw log

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
func (it *FcfaAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaAdminChanged)
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
		it.Event = new(FcfaAdminChanged)
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
func (it *FcfaAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaAdminChanged represents a AdminChanged event raised by the Fcfa contract.
type FcfaAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Fcfa *FcfaFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*FcfaAdminChangedIterator, error) {

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &FcfaAdminChangedIterator{contract: _Fcfa.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Fcfa *FcfaFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *FcfaAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaAdminChanged)
				if err := _Fcfa.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Fcfa *FcfaFilterer) ParseAdminChanged(log types.Log) (*FcfaAdminChanged, error) {
	event := new(FcfaAdminChanged)
	if err := _Fcfa.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Fcfa contract.
type FcfaApprovalIterator struct {
	Event *FcfaApproval // Event containing the contract specifics and raw log

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
func (it *FcfaApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaApproval)
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
		it.Event = new(FcfaApproval)
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
func (it *FcfaApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaApproval represents a Approval event raised by the Fcfa contract.
type FcfaApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Fcfa *FcfaFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*FcfaApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &FcfaApprovalIterator{contract: _Fcfa.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Fcfa *FcfaFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *FcfaApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaApproval)
				if err := _Fcfa.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Fcfa *FcfaFilterer) ParseApproval(log types.Log) (*FcfaApproval, error) {
	event := new(FcfaApproval)
	if err := _Fcfa.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaApproveDepositIterator is returned from FilterApproveDeposit and is used to iterate over the raw logs and unpacked data for ApproveDeposit events raised by the Fcfa contract.
type FcfaApproveDepositIterator struct {
	Event *FcfaApproveDeposit // Event containing the contract specifics and raw log

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
func (it *FcfaApproveDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaApproveDeposit)
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
		it.Event = new(FcfaApproveDeposit)
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
func (it *FcfaApproveDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaApproveDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaApproveDeposit represents a ApproveDeposit event raised by the Fcfa contract.
type FcfaApproveDeposit struct {
	TransactionnId common.Hash
	Approver       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApproveDeposit is a free log retrieval operation binding the contract event 0xf92a27ff89640f3ea2022b7152bd3bd2e0193157d72edb6d9ae6c8f2230108b7.
//
// Solidity: event ApproveDeposit(string indexed transactionnId, address indexed approver)
func (_Fcfa *FcfaFilterer) FilterApproveDeposit(opts *bind.FilterOpts, transactionnId []string, approver []common.Address) (*FcfaApproveDepositIterator, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "ApproveDeposit", transactionnIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return &FcfaApproveDepositIterator{contract: _Fcfa.contract, event: "ApproveDeposit", logs: logs, sub: sub}, nil
}

// WatchApproveDeposit is a free log subscription operation binding the contract event 0xf92a27ff89640f3ea2022b7152bd3bd2e0193157d72edb6d9ae6c8f2230108b7.
//
// Solidity: event ApproveDeposit(string indexed transactionnId, address indexed approver)
func (_Fcfa *FcfaFilterer) WatchApproveDeposit(opts *bind.WatchOpts, sink chan<- *FcfaApproveDeposit, transactionnId []string, approver []common.Address) (event.Subscription, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "ApproveDeposit", transactionnIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaApproveDeposit)
				if err := _Fcfa.contract.UnpackLog(event, "ApproveDeposit", log); err != nil {
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

// ParseApproveDeposit is a log parse operation binding the contract event 0xf92a27ff89640f3ea2022b7152bd3bd2e0193157d72edb6d9ae6c8f2230108b7.
//
// Solidity: event ApproveDeposit(string indexed transactionnId, address indexed approver)
func (_Fcfa *FcfaFilterer) ParseApproveDeposit(log types.Log) (*FcfaApproveDeposit, error) {
	event := new(FcfaApproveDeposit)
	if err := _Fcfa.contract.UnpackLog(event, "ApproveDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaApproveWithdrawIterator is returned from FilterApproveWithdraw and is used to iterate over the raw logs and unpacked data for ApproveWithdraw events raised by the Fcfa contract.
type FcfaApproveWithdrawIterator struct {
	Event *FcfaApproveWithdraw // Event containing the contract specifics and raw log

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
func (it *FcfaApproveWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaApproveWithdraw)
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
		it.Event = new(FcfaApproveWithdraw)
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
func (it *FcfaApproveWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaApproveWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaApproveWithdraw represents a ApproveWithdraw event raised by the Fcfa contract.
type FcfaApproveWithdraw struct {
	TransactionnId common.Hash
	Approver       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterApproveWithdraw is a free log retrieval operation binding the contract event 0xdf78545f0bfafc5998a08d764f2ac01a93d1b37fdf8a34e5ef83d439d331c230.
//
// Solidity: event ApproveWithdraw(string indexed transactionnId, address indexed approver)
func (_Fcfa *FcfaFilterer) FilterApproveWithdraw(opts *bind.FilterOpts, transactionnId []string, approver []common.Address) (*FcfaApproveWithdrawIterator, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "ApproveWithdraw", transactionnIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return &FcfaApproveWithdrawIterator{contract: _Fcfa.contract, event: "ApproveWithdraw", logs: logs, sub: sub}, nil
}

// WatchApproveWithdraw is a free log subscription operation binding the contract event 0xdf78545f0bfafc5998a08d764f2ac01a93d1b37fdf8a34e5ef83d439d331c230.
//
// Solidity: event ApproveWithdraw(string indexed transactionnId, address indexed approver)
func (_Fcfa *FcfaFilterer) WatchApproveWithdraw(opts *bind.WatchOpts, sink chan<- *FcfaApproveWithdraw, transactionnId []string, approver []common.Address) (event.Subscription, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var approverRule []interface{}
	for _, approverItem := range approver {
		approverRule = append(approverRule, approverItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "ApproveWithdraw", transactionnIdRule, approverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaApproveWithdraw)
				if err := _Fcfa.contract.UnpackLog(event, "ApproveWithdraw", log); err != nil {
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

// ParseApproveWithdraw is a log parse operation binding the contract event 0xdf78545f0bfafc5998a08d764f2ac01a93d1b37fdf8a34e5ef83d439d331c230.
//
// Solidity: event ApproveWithdraw(string indexed transactionnId, address indexed approver)
func (_Fcfa *FcfaFilterer) ParseApproveWithdraw(log types.Log) (*FcfaApproveWithdraw, error) {
	event := new(FcfaApproveWithdraw)
	if err := _Fcfa.contract.UnpackLog(event, "ApproveWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Fcfa contract.
type FcfaBeaconUpgradedIterator struct {
	Event *FcfaBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *FcfaBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaBeaconUpgraded)
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
		it.Event = new(FcfaBeaconUpgraded)
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
func (it *FcfaBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaBeaconUpgraded represents a BeaconUpgraded event raised by the Fcfa contract.
type FcfaBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Fcfa *FcfaFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*FcfaBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &FcfaBeaconUpgradedIterator{contract: _Fcfa.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Fcfa *FcfaFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *FcfaBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaBeaconUpgraded)
				if err := _Fcfa.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Fcfa *FcfaFilterer) ParseBeaconUpgraded(log types.Log) (*FcfaBeaconUpgraded, error) {
	event := new(FcfaBeaconUpgraded)
	if err := _Fcfa.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaCancelDepositIterator is returned from FilterCancelDeposit and is used to iterate over the raw logs and unpacked data for CancelDeposit events raised by the Fcfa contract.
type FcfaCancelDepositIterator struct {
	Event *FcfaCancelDeposit // Event containing the contract specifics and raw log

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
func (it *FcfaCancelDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaCancelDeposit)
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
		it.Event = new(FcfaCancelDeposit)
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
func (it *FcfaCancelDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaCancelDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaCancelDeposit represents a CancelDeposit event raised by the Fcfa contract.
type FcfaCancelDeposit struct {
	TransactionnId common.Hash
	Canceler       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCancelDeposit is a free log retrieval operation binding the contract event 0xf8006234b98a8a5d1104283e880b00bf159e0c65d70ab13ab479ae8cf6a40662.
//
// Solidity: event CancelDeposit(string indexed transactionnId, address indexed canceler)
func (_Fcfa *FcfaFilterer) FilterCancelDeposit(opts *bind.FilterOpts, transactionnId []string, canceler []common.Address) (*FcfaCancelDepositIterator, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var cancelerRule []interface{}
	for _, cancelerItem := range canceler {
		cancelerRule = append(cancelerRule, cancelerItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "CancelDeposit", transactionnIdRule, cancelerRule)
	if err != nil {
		return nil, err
	}
	return &FcfaCancelDepositIterator{contract: _Fcfa.contract, event: "CancelDeposit", logs: logs, sub: sub}, nil
}

// WatchCancelDeposit is a free log subscription operation binding the contract event 0xf8006234b98a8a5d1104283e880b00bf159e0c65d70ab13ab479ae8cf6a40662.
//
// Solidity: event CancelDeposit(string indexed transactionnId, address indexed canceler)
func (_Fcfa *FcfaFilterer) WatchCancelDeposit(opts *bind.WatchOpts, sink chan<- *FcfaCancelDeposit, transactionnId []string, canceler []common.Address) (event.Subscription, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var cancelerRule []interface{}
	for _, cancelerItem := range canceler {
		cancelerRule = append(cancelerRule, cancelerItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "CancelDeposit", transactionnIdRule, cancelerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaCancelDeposit)
				if err := _Fcfa.contract.UnpackLog(event, "CancelDeposit", log); err != nil {
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

// ParseCancelDeposit is a log parse operation binding the contract event 0xf8006234b98a8a5d1104283e880b00bf159e0c65d70ab13ab479ae8cf6a40662.
//
// Solidity: event CancelDeposit(string indexed transactionnId, address indexed canceler)
func (_Fcfa *FcfaFilterer) ParseCancelDeposit(log types.Log) (*FcfaCancelDeposit, error) {
	event := new(FcfaCancelDeposit)
	if err := _Fcfa.contract.UnpackLog(event, "CancelDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaCancelWithdrawIterator is returned from FilterCancelWithdraw and is used to iterate over the raw logs and unpacked data for CancelWithdraw events raised by the Fcfa contract.
type FcfaCancelWithdrawIterator struct {
	Event *FcfaCancelWithdraw // Event containing the contract specifics and raw log

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
func (it *FcfaCancelWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaCancelWithdraw)
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
		it.Event = new(FcfaCancelWithdraw)
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
func (it *FcfaCancelWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaCancelWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaCancelWithdraw represents a CancelWithdraw event raised by the Fcfa contract.
type FcfaCancelWithdraw struct {
	TransactionnId common.Hash
	Canceler       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCancelWithdraw is a free log retrieval operation binding the contract event 0x2d365e7e41ce37b636dac5f661aadcd49e813866124180449e057e59ea79c323.
//
// Solidity: event CancelWithdraw(string indexed transactionnId, address indexed canceler)
func (_Fcfa *FcfaFilterer) FilterCancelWithdraw(opts *bind.FilterOpts, transactionnId []string, canceler []common.Address) (*FcfaCancelWithdrawIterator, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var cancelerRule []interface{}
	for _, cancelerItem := range canceler {
		cancelerRule = append(cancelerRule, cancelerItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "CancelWithdraw", transactionnIdRule, cancelerRule)
	if err != nil {
		return nil, err
	}
	return &FcfaCancelWithdrawIterator{contract: _Fcfa.contract, event: "CancelWithdraw", logs: logs, sub: sub}, nil
}

// WatchCancelWithdraw is a free log subscription operation binding the contract event 0x2d365e7e41ce37b636dac5f661aadcd49e813866124180449e057e59ea79c323.
//
// Solidity: event CancelWithdraw(string indexed transactionnId, address indexed canceler)
func (_Fcfa *FcfaFilterer) WatchCancelWithdraw(opts *bind.WatchOpts, sink chan<- *FcfaCancelWithdraw, transactionnId []string, canceler []common.Address) (event.Subscription, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var cancelerRule []interface{}
	for _, cancelerItem := range canceler {
		cancelerRule = append(cancelerRule, cancelerItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "CancelWithdraw", transactionnIdRule, cancelerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaCancelWithdraw)
				if err := _Fcfa.contract.UnpackLog(event, "CancelWithdraw", log); err != nil {
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

// ParseCancelWithdraw is a log parse operation binding the contract event 0x2d365e7e41ce37b636dac5f661aadcd49e813866124180449e057e59ea79c323.
//
// Solidity: event CancelWithdraw(string indexed transactionnId, address indexed canceler)
func (_Fcfa *FcfaFilterer) ParseCancelWithdraw(log types.Log) (*FcfaCancelWithdraw, error) {
	event := new(FcfaCancelWithdraw)
	if err := _Fcfa.contract.UnpackLog(event, "CancelWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Fcfa contract.
type FcfaDepositIterator struct {
	Event *FcfaDeposit // Event containing the contract specifics and raw log

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
func (it *FcfaDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaDeposit)
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
		it.Event = new(FcfaDeposit)
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
func (it *FcfaDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaDeposit represents a Deposit event raised by the Fcfa contract.
type FcfaDeposit struct {
	TransactionnId common.Hash
	Sender         common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xd327b35e36b3981157588978d60961f5c09dc2926008abb81dd77b1197a416ed.
//
// Solidity: event Deposit(string indexed transactionnId, address indexed sender, uint256 amount)
func (_Fcfa *FcfaFilterer) FilterDeposit(opts *bind.FilterOpts, transactionnId []string, sender []common.Address) (*FcfaDepositIterator, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "Deposit", transactionnIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FcfaDepositIterator{contract: _Fcfa.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xd327b35e36b3981157588978d60961f5c09dc2926008abb81dd77b1197a416ed.
//
// Solidity: event Deposit(string indexed transactionnId, address indexed sender, uint256 amount)
func (_Fcfa *FcfaFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *FcfaDeposit, transactionnId []string, sender []common.Address) (event.Subscription, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "Deposit", transactionnIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaDeposit)
				if err := _Fcfa.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xd327b35e36b3981157588978d60961f5c09dc2926008abb81dd77b1197a416ed.
//
// Solidity: event Deposit(string indexed transactionnId, address indexed sender, uint256 amount)
func (_Fcfa *FcfaFilterer) ParseDeposit(log types.Log) (*FcfaDeposit, error) {
	event := new(FcfaDeposit)
	if err := _Fcfa.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Fcfa contract.
type FcfaInitializedIterator struct {
	Event *FcfaInitialized // Event containing the contract specifics and raw log

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
func (it *FcfaInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaInitialized)
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
		it.Event = new(FcfaInitialized)
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
func (it *FcfaInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaInitialized represents a Initialized event raised by the Fcfa contract.
type FcfaInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Fcfa *FcfaFilterer) FilterInitialized(opts *bind.FilterOpts) (*FcfaInitializedIterator, error) {

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FcfaInitializedIterator{contract: _Fcfa.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Fcfa *FcfaFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FcfaInitialized) (event.Subscription, error) {

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaInitialized)
				if err := _Fcfa.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Fcfa *FcfaFilterer) ParseInitialized(log types.Log) (*FcfaInitialized, error) {
	event := new(FcfaInitialized)
	if err := _Fcfa.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Fcfa contract.
type FcfaPausedIterator struct {
	Event *FcfaPaused // Event containing the contract specifics and raw log

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
func (it *FcfaPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaPaused)
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
		it.Event = new(FcfaPaused)
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
func (it *FcfaPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaPaused represents a Paused event raised by the Fcfa contract.
type FcfaPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Fcfa *FcfaFilterer) FilterPaused(opts *bind.FilterOpts) (*FcfaPausedIterator, error) {

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &FcfaPausedIterator{contract: _Fcfa.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Fcfa *FcfaFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *FcfaPaused) (event.Subscription, error) {

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaPaused)
				if err := _Fcfa.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Fcfa *FcfaFilterer) ParsePaused(log types.Log) (*FcfaPaused, error) {
	event := new(FcfaPaused)
	if err := _Fcfa.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaRemoveKYCFlagIterator is returned from FilterRemoveKYCFlag and is used to iterate over the raw logs and unpacked data for RemoveKYCFlag events raised by the Fcfa contract.
type FcfaRemoveKYCFlagIterator struct {
	Event *FcfaRemoveKYCFlag // Event containing the contract specifics and raw log

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
func (it *FcfaRemoveKYCFlagIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaRemoveKYCFlag)
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
		it.Event = new(FcfaRemoveKYCFlag)
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
func (it *FcfaRemoveKYCFlagIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaRemoveKYCFlagIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaRemoveKYCFlag represents a RemoveKYCFlag event raised by the Fcfa contract.
type FcfaRemoveKYCFlag struct {
	Account common.Address
	Remover common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRemoveKYCFlag is a free log retrieval operation binding the contract event 0xfb456b249009c196b58086480400da57f56b6214a5bccad64b4f952e909969e7.
//
// Solidity: event RemoveKYCFlag(address indexed account, address indexed remover)
func (_Fcfa *FcfaFilterer) FilterRemoveKYCFlag(opts *bind.FilterOpts, account []common.Address, remover []common.Address) (*FcfaRemoveKYCFlagIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "RemoveKYCFlag", accountRule, removerRule)
	if err != nil {
		return nil, err
	}
	return &FcfaRemoveKYCFlagIterator{contract: _Fcfa.contract, event: "RemoveKYCFlag", logs: logs, sub: sub}, nil
}

// WatchRemoveKYCFlag is a free log subscription operation binding the contract event 0xfb456b249009c196b58086480400da57f56b6214a5bccad64b4f952e909969e7.
//
// Solidity: event RemoveKYCFlag(address indexed account, address indexed remover)
func (_Fcfa *FcfaFilterer) WatchRemoveKYCFlag(opts *bind.WatchOpts, sink chan<- *FcfaRemoveKYCFlag, account []common.Address, remover []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var removerRule []interface{}
	for _, removerItem := range remover {
		removerRule = append(removerRule, removerItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "RemoveKYCFlag", accountRule, removerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaRemoveKYCFlag)
				if err := _Fcfa.contract.UnpackLog(event, "RemoveKYCFlag", log); err != nil {
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

// ParseRemoveKYCFlag is a log parse operation binding the contract event 0xfb456b249009c196b58086480400da57f56b6214a5bccad64b4f952e909969e7.
//
// Solidity: event RemoveKYCFlag(address indexed account, address indexed remover)
func (_Fcfa *FcfaFilterer) ParseRemoveKYCFlag(log types.Log) (*FcfaRemoveKYCFlag, error) {
	event := new(FcfaRemoveKYCFlag)
	if err := _Fcfa.contract.UnpackLog(event, "RemoveKYCFlag", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Fcfa contract.
type FcfaRoleAdminChangedIterator struct {
	Event *FcfaRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *FcfaRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaRoleAdminChanged)
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
		it.Event = new(FcfaRoleAdminChanged)
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
func (it *FcfaRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaRoleAdminChanged represents a RoleAdminChanged event raised by the Fcfa contract.
type FcfaRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Fcfa *FcfaFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*FcfaRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &FcfaRoleAdminChangedIterator{contract: _Fcfa.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Fcfa *FcfaFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *FcfaRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaRoleAdminChanged)
				if err := _Fcfa.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Fcfa *FcfaFilterer) ParseRoleAdminChanged(log types.Log) (*FcfaRoleAdminChanged, error) {
	event := new(FcfaRoleAdminChanged)
	if err := _Fcfa.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Fcfa contract.
type FcfaRoleGrantedIterator struct {
	Event *FcfaRoleGranted // Event containing the contract specifics and raw log

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
func (it *FcfaRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaRoleGranted)
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
		it.Event = new(FcfaRoleGranted)
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
func (it *FcfaRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaRoleGranted represents a RoleGranted event raised by the Fcfa contract.
type FcfaRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fcfa *FcfaFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FcfaRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FcfaRoleGrantedIterator{contract: _Fcfa.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fcfa *FcfaFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *FcfaRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaRoleGranted)
				if err := _Fcfa.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fcfa *FcfaFilterer) ParseRoleGranted(log types.Log) (*FcfaRoleGranted, error) {
	event := new(FcfaRoleGranted)
	if err := _Fcfa.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Fcfa contract.
type FcfaRoleRevokedIterator struct {
	Event *FcfaRoleRevoked // Event containing the contract specifics and raw log

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
func (it *FcfaRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaRoleRevoked)
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
		it.Event = new(FcfaRoleRevoked)
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
func (it *FcfaRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaRoleRevoked represents a RoleRevoked event raised by the Fcfa contract.
type FcfaRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fcfa *FcfaFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*FcfaRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FcfaRoleRevokedIterator{contract: _Fcfa.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fcfa *FcfaFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *FcfaRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaRoleRevoked)
				if err := _Fcfa.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Fcfa *FcfaFilterer) ParseRoleRevoked(log types.Log) (*FcfaRoleRevoked, error) {
	event := new(FcfaRoleRevoked)
	if err := _Fcfa.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Fcfa contract.
type FcfaTransferIterator struct {
	Event *FcfaTransfer // Event containing the contract specifics and raw log

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
func (it *FcfaTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaTransfer)
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
		it.Event = new(FcfaTransfer)
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
func (it *FcfaTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaTransfer represents a Transfer event raised by the Fcfa contract.
type FcfaTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Fcfa *FcfaFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*FcfaTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &FcfaTransferIterator{contract: _Fcfa.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Fcfa *FcfaFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FcfaTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaTransfer)
				if err := _Fcfa.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Fcfa *FcfaFilterer) ParseTransfer(log types.Log) (*FcfaTransfer, error) {
	event := new(FcfaTransfer)
	if err := _Fcfa.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Fcfa contract.
type FcfaUnpausedIterator struct {
	Event *FcfaUnpaused // Event containing the contract specifics and raw log

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
func (it *FcfaUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaUnpaused)
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
		it.Event = new(FcfaUnpaused)
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
func (it *FcfaUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaUnpaused represents a Unpaused event raised by the Fcfa contract.
type FcfaUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Fcfa *FcfaFilterer) FilterUnpaused(opts *bind.FilterOpts) (*FcfaUnpausedIterator, error) {

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &FcfaUnpausedIterator{contract: _Fcfa.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Fcfa *FcfaFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *FcfaUnpaused) (event.Subscription, error) {

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaUnpaused)
				if err := _Fcfa.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Fcfa *FcfaFilterer) ParseUnpaused(log types.Log) (*FcfaUnpaused, error) {
	event := new(FcfaUnpaused)
	if err := _Fcfa.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Fcfa contract.
type FcfaUpgradedIterator struct {
	Event *FcfaUpgraded // Event containing the contract specifics and raw log

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
func (it *FcfaUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaUpgraded)
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
		it.Event = new(FcfaUpgraded)
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
func (it *FcfaUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaUpgraded represents a Upgraded event raised by the Fcfa contract.
type FcfaUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fcfa *FcfaFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*FcfaUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &FcfaUpgradedIterator{contract: _Fcfa.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fcfa *FcfaFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *FcfaUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaUpgraded)
				if err := _Fcfa.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Fcfa *FcfaFilterer) ParseUpgraded(log types.Log) (*FcfaUpgraded, error) {
	event := new(FcfaUpgraded)
	if err := _Fcfa.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FcfaWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Fcfa contract.
type FcfaWithdrawIterator struct {
	Event *FcfaWithdraw // Event containing the contract specifics and raw log

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
func (it *FcfaWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FcfaWithdraw)
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
		it.Event = new(FcfaWithdraw)
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
func (it *FcfaWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FcfaWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FcfaWithdraw represents a Withdraw event raised by the Fcfa contract.
type FcfaWithdraw struct {
	TransactionnId common.Hash
	Sender         common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x7a5e5c901f9da945e2028bd646eb1f6842dff416e862cfa502e0d3408635143a.
//
// Solidity: event Withdraw(string indexed transactionnId, address indexed sender, uint256 amount)
func (_Fcfa *FcfaFilterer) FilterWithdraw(opts *bind.FilterOpts, transactionnId []string, sender []common.Address) (*FcfaWithdrawIterator, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fcfa.contract.FilterLogs(opts, "Withdraw", transactionnIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &FcfaWithdrawIterator{contract: _Fcfa.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x7a5e5c901f9da945e2028bd646eb1f6842dff416e862cfa502e0d3408635143a.
//
// Solidity: event Withdraw(string indexed transactionnId, address indexed sender, uint256 amount)
func (_Fcfa *FcfaFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FcfaWithdraw, transactionnId []string, sender []common.Address) (event.Subscription, error) {

	var transactionnIdRule []interface{}
	for _, transactionnIdItem := range transactionnId {
		transactionnIdRule = append(transactionnIdRule, transactionnIdItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Fcfa.contract.WatchLogs(opts, "Withdraw", transactionnIdRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FcfaWithdraw)
				if err := _Fcfa.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x7a5e5c901f9da945e2028bd646eb1f6842dff416e862cfa502e0d3408635143a.
//
// Solidity: event Withdraw(string indexed transactionnId, address indexed sender, uint256 amount)
func (_Fcfa *FcfaFilterer) ParseWithdraw(log types.Log) (*FcfaWithdraw, error) {
	event := new(FcfaWithdraw)
	if err := _Fcfa.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
