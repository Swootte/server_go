package finance

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"server/errors"
	"server/ethereum"
	"server/fcfa"
	"server/utils"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	solsha "github.com/miguelmota/go-solidity-sha3"
)

var _client ethclient.Client
var FCFAIntance *fcfa.Fcfa

func init() {
	utils.LoadEnv()
	client, err := ethclient.Dial(os.Getenv("CHAIN_LINK"))
	if err != nil {
		panic(err)
	}
	_client = *client
	_instance, err := loadContractInstance(os.Getenv("DEFAULT_CURRENCY"))
	if err != nil {
		panic(err)
	}

	FCFAIntance = _instance
}

func FromWei(amount big.Int) *big.Float {
	decimals, _ := FCFAIntance.Decimals(nil)
	fbalance := new(big.Float)
	fbalance.SetString(amount.String())
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(int(decimals))))
}

func FromWeiFloat(amount float64) *big.Float {
	decimals, _ := FCFAIntance.Decimals(nil)
	fbalance := new(big.Float)
	fbalance.SetFloat64(amount)
	return new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(int(decimals))))
}

func ToWei(amount float64) *big.Int {
	decimals, _ := FCFAIntance.Decimals(nil)
	fbalance := big.NewFloat(amount)
	result := new(big.Float).Mul(fbalance, big.NewFloat(math.Pow10(int(decimals))))
	output, _ := result.Int(big.NewInt(0))
	return output
}

func ToWeiFromInt(amount int) *big.Int {
	amountFloat := float64(int64(amount))
	decimals, _ := FCFAIntance.Decimals(nil)
	fbalance := big.NewFloat(amountFloat)
	result := new(big.Float).Mul(fbalance, big.NewFloat(math.Pow10(int(decimals))))
	output, _ := result.Int(big.NewInt(0))
	return output
}

func CalculateFee(amount float64) float64 {
	_amount := big.NewFloat(amount)
	new_amount := new(big.Float).Mul(_amount, new(big.Float).Quo(big.NewFloat(1), big.NewFloat(100)))
	_new_amount, _ := new_amount.Float64()
	return _new_amount
}

func LoadTokenSupply() (*big.Int, error) {
	return FCFAIntance.TotalSupply(&bind.CallOpts{
		From: common.HexToAddress(os.Getenv("CHAIN_ADDRESS")),
	})
}

func CreateAccount() (string, string, bool) {
	privateKey := ethereum.CreateWallet()
	privateKeyBytes := crypto.FromECDSA(privateKey)
	privateKeyDB := hexutil.Encode(privateKeyBytes)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", "", ok
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return privateKeyDB, address, false
}

func getModSigner() (*ecdsa.PrivateKey, error) {
	key, err := crypto.HexToECDSA(os.Getenv("CHAIN_PRIVATE_KEY")[2:])
	if err != nil {
		return nil, err
	}
	return key, nil
}

func getSigner(privateKey string) (*ecdsa.PrivateKey, error) {
	key, err := crypto.HexToECDSA(privateKey[2:])
	if err != nil {
		return nil, err
	}
	return key, nil
}

func EthSign(prefixed []byte, signer *ecdsa.PrivateKey) ([]byte, error) {
	sig, err := crypto.Sign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	if sig[64] < 27 {
		sig[64] += 27
	}

	return sig, nil
}

func ethereumBalance(publicKey string) *big.Int {
	account := common.HexToAddress(publicKey)
	balance, err := _client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	return balance
}

func loadContractInstance(contractAdress string) (*fcfa.Fcfa, error) {
	address := common.HexToAddress(contractAdress)
	instance, err := fcfa.NewFcfa(address, &_client)
	if err != nil {
		return nil, err
	}
	return instance, err
}

func GetBalanceOnContractInstance(publicKey string) (*big.Float, error) {
	address := common.HexToAddress(publicKey)
	balance, err := FCFAIntance.BalanceOf(&bind.CallOpts{}, address)
	if err != nil {
		return nil, err
	}

	return FromWei(*balance), nil
}

func getGas() *big.Int {
	gasPrice, _ := _client.SuggestGasPrice(context.Background())
	return gasPrice
}

func getNonce() uint64 {
	address := common.HexToAddress(os.Getenv("CHAIN_ADDRESS"))
	nonce, _ := _client.PendingNonceAt(context.Background(), address)
	return nonce

}

func getTransactionOpt() (*bind.TransactOpts, error) {
	__amdin_adress := common.HexToAddress(os.Getenv("CHAIN_ADDRESS"))
	mod, err := getModSigner()
	if err != nil {
		return nil, err
	}

	chainId, errs := _client.ChainID(context.Background())
	if errs != nil {
		return nil, errs
	}

	modsign, err := bind.NewKeyedTransactorWithChainID(mod, chainId)

	if err != nil {
		return nil, err
	}

	return &bind.TransactOpts{
		From:     __amdin_adress,
		Nonce:    big.NewInt(int64(getNonce())),
		Signer:   modsign.Signer,
		GasPrice: getGas(),
		Context:  context.Background(),
		NoSend:   false,
	}, nil
}

func Transfer(payer string, payee string, amount *big.Int, privateKey string) (*types.Receipt, error) {
	from := common.HexToAddress(payer)
	to := common.HexToAddress(payee)
	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}
	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.Address(to),
		solsha.Uint256(amount),
		solsha.Address(from),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.Transfer(transactionOps, to, amount, from, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}
	return bind.WaitMined(context.Background(), &_client, tx)
}

func AskWithdraw(amount *big.Int, privateKey string, transactionId string, payer string) (*types.Receipt, error) {
	from := common.HexToAddress(payer)
	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.String(transactionId),
		solsha.Uint256(amount),
		solsha.Address(from),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.Withdraw(transactionOps, transactionId, amount, from, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}

	return bind.WaitMined(context.Background(), &_client, tx)
}

func ApproveWithdraw(transactionId string, privateKey string, approver string) (*types.Receipt, error) {
	_approver := common.HexToAddress(approver)
	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.String(transactionId),
		solsha.Address(_approver),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.ApproveWithdraw(transactionOps, transactionId, _approver, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}

	return bind.WaitMined(context.Background(), &_client, tx)
}

func Deposit(transactionId string, amount *big.Int, privateKey string, depositorAddress string) (*types.Receipt, error) {
	_depositor := common.HexToAddress(depositorAddress)
	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.String(transactionId),
		solsha.Uint256(amount),
		solsha.Address(_depositor),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.Deposit(transactionOps, transactionId, amount, _depositor, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}

	return bind.WaitMined(context.Background(), &_client, tx)
}

func ApproveDeposit(depositId string, privateKey string, approver string) (*types.Receipt, error) {
	_approver := common.HexToAddress(approver)
	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.String(depositId),
		solsha.Address(_approver),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))
	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.ApproveDeposit(transactionOps, depositId, _approver, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}

	return bind.WaitMined(context.Background(), &_client, tx)
}

func CancelWithDraw(transactionId string, privateKey string, canceller string) (*types.Receipt, error) {
	_canceller := common.HexToAddress(canceller)
	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.String(transactionId),
		solsha.Address(_canceller),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.CancelWithdraw(transactionOps, transactionId, _canceller, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}
	return bind.WaitMined(context.Background(), &_client, tx)
}

func CancelDeposit(transactionId string, privateKey string, canceller string) (*types.Receipt, error) {
	_canceller := common.HexToAddress(canceller)
	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.String(transactionId),
		solsha.Address(_canceller),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.CancelDeposit(transactionOps, transactionId, _canceller, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}
	return bind.WaitMined(context.Background(), &_client, tx)
}

func GiveRole(account string, role string, privateKey string, giver string) (*types.Receipt, error) {
	_account := common.HexToAddress(account)
	_giver := common.HexToAddress(giver)
	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.Address(_account),
		solsha.String(role),
		solsha.Address(_giver),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.GiveRoleToAddress(transactionOps, _account, role, _giver, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}
	return bind.WaitMined(context.Background(), &_client, tx)
}

func RemoveRole(account string, role string, privateKey string, remover string) (*types.Receipt, error) {
	_account := common.HexToAddress(account)
	_remover := common.HexToAddress(remover)
	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.Address(_account),
		solsha.String(role),
		solsha.Address(_remover),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.RemoveroleFromAddress(transactionOps, _account, role, _remover, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}

	return bind.WaitMined(context.Background(), &_client, tx)
}

func CommercePay(transactionId string, payee string, amount *big.Int, privateKey string, payer string) (*types.Receipt, error) {
	_payee := common.HexToAddress(payee)
	_payer := common.HexToAddress(payer)

	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.String(transactionId),
		solsha.Address(_payee),
		solsha.Uint256(amount),
		solsha.Address(_payer),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.Commercepay(transactionOps, transactionId, _payee, amount, _payer, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}
	return bind.WaitMined(context.Background(), &_client, tx)
}

func AddUnConfirmedTransaction(transactionId string, to string, amount *big.Int) (*types.Receipt, error) {
	transactionOps, err := getTransactionOpt()
	_payee := common.HexToAddress(to)
	if err != nil {
		return nil, err
	}
	tx, err := FCFAIntance.AddUnconfirmedPayment(transactionOps, transactionId, _payee, amount)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}
	return bind.WaitMined(context.Background(), &_client, tx)
}

func ConfirmUnformedTransaction(transactionId string, privateKey string, confirmer string) (*types.Receipt, error) {
	_payer := common.HexToAddress(confirmer)

	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.String(transactionId),
		solsha.Address(_payer),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.ConfirmUnconfirmedPayment(transactionOps, transactionId, _payer, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}
	return bind.WaitMined(context.Background(), &_client, tx)

}

func RefundCommercePay(transactionId string, privateKey string, refunder string) (*types.Receipt, error) {
	_refunder := common.HexToAddress(refunder)
	transactionOps, err := getTransactionOpt()
	if err != nil {
		return nil, err
	}

	signer, err := getSigner(privateKey)
	if err != nil {
		return nil, err
	}

	hash := solsha.SoliditySHA3(
		solsha.String(transactionId),
		solsha.Address(_refunder),
	)

	prefixed := solsha.SoliditySHA3(solsha.String("\x19Ethereum Signed Message:\n32"), solsha.Bytes32(hash))

	sig, err := EthSign(prefixed, signer)
	if err != nil {
		return nil, err
	}

	tx, err := FCFAIntance.Refund(transactionOps, transactionId, _refunder, sig)
	if err != nil {
		f := strings.Replace(err.Error(), "execution reverted: ", "", -1)
		_f := strings.TrimSpace(f)
		return nil, fmt.Errorf(errors.GetEthereumError(_f))
	}

	return bind.WaitMined(context.Background(), &_client, tx)
}
