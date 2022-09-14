package finance_test

import (
	"math/big"
	"os"
	"server/finance"
	"server/utils"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func init() {
	utils.LoadEnv()
}

func TestFinance(t *testing.T) {
	var first_privateKey string
	var first_publicKey string

	var second_privateKey string
	var second_publicKey string

	var third_privateKey string
	var third_publicKey string

	topupuid := gofakeit.UUID()
	secondtopupuid := gofakeit.UUID()

	withdrawId := gofakeit.UUID()
	commercepayId := gofakeit.UUID()

	t.Run("create ethereum account", func(t *testing.T) {
		private, public, hasError := finance.CreateAccount()
		require.Equal(t, hasError, false)
		require.NotEmpty(t, public)
		require.NotEmpty(t, private)
		first_privateKey = private
		first_publicKey = public
	})

	t.Run("create a second account", func(t *testing.T) {
		private, public, hasError := finance.CreateAccount()
		require.Equal(t, hasError, false)
		require.NotEmpty(t, public)
		require.NotEmpty(t, private)
		third_privateKey = private
		third_publicKey = public
	})

	t.Run("create a third account", func(t *testing.T) {
		private, public, hasError := finance.CreateAccount()
		require.Equal(t, hasError, false)
		require.NotEmpty(t, public)
		require.NotEmpty(t, private)
		second_privateKey = private
		second_publicKey = public
	})

	t.Run("getBalance of the first user", func(t *testing.T) {
		balance, err := finance.GetBalanceOnContractInstance(first_publicKey)
		require.Empty(t, err)
		require.Equal(t, big.NewFloat(0).String(), balance.String())
	})

	t.Run("getBalance of the second user", func(t *testing.T) {
		balance, err := finance.GetBalanceOnContractInstance(first_publicKey)
		require.Empty(t, err)
		require.Equal(t, big.NewFloat(0).String(), balance.String())
	})

	t.Run("Ask deposit, first account", func(t *testing.T) {
		_amount := finance.ToWei(10000)
		topup, err := finance.Deposit(topupuid, _amount, first_privateKey, first_publicKey)
		require.Empty(t, err)
		require.NotEmpty(t, topup.TxHash)
	})

	t.Run("Ask deposit, second account", func(t *testing.T) {
		_amount := finance.ToWei(10000)
		topup, err := finance.Deposit(secondtopupuid, _amount, second_privateKey, second_publicKey)
		require.Empty(t, err)
		require.NotEmpty(t, topup.TxHash)
	})

	t.Run("should validate topup, first account", func(t *testing.T) {
		receipt, err := finance.ApproveDeposit(topupuid, os.Getenv("CHAIN_PRIVATE_KEY"), os.Getenv("CHAIN_ADDRESS"))
		require.NotEmpty(t, receipt.TxHash)
		require.Empty(t, err)
		balance, err := finance.GetBalanceOnContractInstance(first_publicKey)
		require.Empty(t, err)
		require.Equal(t, balance.String(), big.NewFloat(10000).String())
	})

	t.Run("should ask for withdraw", func(t *testing.T) {
		receipt, err := finance.AskWithdraw(finance.ToWei(100), first_privateKey, withdrawId, first_publicKey)
		require.Empty(t, err)
		require.NotEmpty(t, receipt.TxHash)
	})

	t.Run("should validate withdraw", func(t *testing.T) {
		receipt, err := finance.ApproveWithdraw(withdrawId, os.Getenv("CHAIN_PRIVATE_KEY"), os.Getenv("CHAIN_ADDRESS"))
		require.NotEmpty(t, receipt.TxHash)
		require.Empty(t, err)
		balance, err := finance.GetBalanceOnContractInstance(first_publicKey)
		require.Empty(t, err)
		require.Equal(t, balance.String(), big.NewFloat(9900).String())
	})

	t.Run("should  grant agent role to user and validate a transaction", func(t *testing.T) {
		receipt, err := finance.GiveRole(third_publicKey, "AGENT_ROLE", os.Getenv("CHAIN_PRIVATE_KEY"), os.Getenv("CHAIN_ADDRESS"))
		require.NotEmpty(t, receipt.TxHash)
		require.Empty(t, err)
		txUID := gofakeit.UUID()
		topup, err := finance.Deposit(txUID, finance.ToWei(1000), second_privateKey, second_publicKey)
		require.Empty(t, err)
		require.NotEmpty(t, topup.TxHash)
		receiptDepositApprove, errDepositApprove := finance.ApproveDeposit(txUID, third_privateKey, third_publicKey)
		require.NotEmpty(t, receiptDepositApprove.TxHash)
		require.Empty(t, errDepositApprove)
		balance, err := finance.GetBalanceOnContractInstance(second_publicKey)
		require.Empty(t, err)
		require.Equal(t, balance.String(), big.NewFloat(1000).String())
	})

	t.Run("should remove role and try to  validate a transaction that should result in error", func(t *testing.T) {
		receipt, err := finance.RemoveRole(third_publicKey, "AGENT_ROLE", os.Getenv("CHAIN_PRIVATE_KEY"), os.Getenv("CHAIN_ADDRESS"))
		require.NotEmpty(t, receipt.TxHash)
		require.Empty(t, err)
		txUID := gofakeit.UUID()
		_amount := finance.ToWei(1000)
		topup, err := finance.Deposit(txUID, _amount, second_privateKey, second_publicKey)
		require.Empty(t, err)
		require.NotEmpty(t, topup.TxHash)
		_, errDepositApprove := finance.ApproveDeposit(txUID, third_privateKey, third_publicKey)
		require.NotEmpty(t, errDepositApprove)
	})

	t.Run("should transfer money to a user", func(t *testing.T) {
		amount := finance.ToWei(1000)
		receipt, err := finance.Transfer(first_publicKey, second_publicKey, amount, first_privateKey)
		require.NotEmpty(t, receipt.TxHash)
		require.Empty(t, err)
		balance, err := finance.GetBalanceOnContractInstance(second_publicKey)
		require.Empty(t, err)
		require.Equal(t, balance.String(), big.NewFloat(2000).String())
		balance2, err2 := finance.GetBalanceOnContractInstance(first_publicKey)
		require.Empty(t, err2)
		require.Equal(t, big.NewFloat(8890).String(), balance2.String())
	})

	t.Run("should commercepay", func(t *testing.T) {
		receipt, err := finance.CommercePay(commercepayId, second_publicKey, finance.ToWei(1000), first_privateKey, first_publicKey)
		require.NotEmpty(t, receipt.TxHash)
		require.Empty(t, err)
		balance, err := finance.GetBalanceOnContractInstance(second_publicKey)
		require.Empty(t, err)
		require.Equal(t, big.NewFloat(2990).String(), balance.String())
		balance2, err2 := finance.GetBalanceOnContractInstance(first_publicKey)
		require.Empty(t, err2)
		require.Equal(t, balance2.String(), big.NewFloat(7890).String())
	})

	t.Run("should refund commercepay transaction", func(t *testing.T) {
		receipt, err := finance.RefundCommercePay(commercepayId, second_privateKey, second_publicKey)
		require.NotEmpty(t, receipt.TxHash)
		require.Empty(t, err)
		balance, err := finance.GetBalanceOnContractInstance(second_publicKey)
		require.Empty(t, err)
		require.Equal(t, balance.String(), big.NewFloat(2000).String())
		balance2, err2 := finance.GetBalanceOnContractInstance(first_publicKey)
		require.Empty(t, err2)
		require.Equal(t, balance2.String(), big.NewFloat(8890).String())
	})

	t.Run("should ask topup and cancel it | USER", func(t *testing.T) {
		txUID := gofakeit.UUID()
		topup, err := finance.Deposit(txUID, finance.ToWei(10), second_privateKey, second_publicKey)
		require.NotEmpty(t, topup.TxHash)
		require.Empty(t, err)
		cancelReceipt, errCancel := finance.CancelDeposit(txUID, second_privateKey, second_publicKey)
		require.NotEmpty(t, cancelReceipt.TxHash)
		require.Empty(t, errCancel)
	})

	t.Run("should ask withdraw and cancel it | USER", func(t *testing.T) {
		txUID := gofakeit.UUID()
		topup, err := finance.AskWithdraw(finance.ToWei(10), second_privateKey, txUID, second_publicKey)
		require.NotEmpty(t, topup.TxHash)
		require.Empty(t, err)
		cancelReceipt, err := finance.CancelWithDraw(txUID, second_privateKey, second_publicKey)
		require.NotEmpty(t, cancelReceipt.TxHash)
		require.Empty(t, err)

	})

	t.Run("should ask topup and cancel it | AGENT", func(t *testing.T) {
		txUID := gofakeit.UUID()
		topup, err := finance.Deposit(txUID, finance.ToWei(10), second_privateKey, second_publicKey)
		require.NotEmpty(t, topup.TxHash)
		require.Empty(t, err)
		cancelReceipt, errCancel := finance.CancelDeposit(txUID, os.Getenv("CHAIN_PRIVATE_KEY"), os.Getenv("CHAIN_ADDRESS"))
		require.NotEmpty(t, cancelReceipt.TxHash)
		require.Empty(t, errCancel)
	})

	t.Run("should ask withdraw and cancel it | AGENT", func(t *testing.T) {
		txUID := gofakeit.UUID()
		topup, err := finance.AskWithdraw(finance.ToWei(10), second_privateKey, txUID, second_publicKey)
		require.NotEmpty(t, topup.TxHash)
		require.Empty(t, err)
		cancelReceipt, err := finance.CancelWithDraw(txUID, os.Getenv("CHAIN_PRIVATE_KEY"), os.Getenv("CHAIN_ADDRESS"))
		require.NotEmpty(t, cancelReceipt.TxHash)
		require.Empty(t, err)
	})

	t.Run("should grand managment Role and give Agent Role to other people", func(t *testing.T) {
		receipt, err := finance.GiveRole(third_publicKey, "MANAGMENT_ROLE", os.Getenv("CHAIN_PRIVATE_KEY"), os.Getenv("CHAIN_ADDRESS"))
		require.NotEmpty(t, receipt.TxHash)
		require.Empty(t, err)
		receipt2, err2 := finance.GiveRole(second_publicKey, "AGENT_ROLE", third_privateKey, third_publicKey)
		require.NotEmpty(t, receipt2.TxHash)
		require.Empty(t, err2)
	})
}
