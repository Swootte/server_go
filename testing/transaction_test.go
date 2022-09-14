package testing

import (
	"context"
	"fmt"
	"os"
	"server/agency"
	"server/agent"
	firebase "server/firebase"
	"server/graph/model"
	"server/middleware"
	"server/serverutils"
	"server/transaction"
	"server/user"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestTransaction(t *testing.T) {
	handlr := serverutils.StartServer()
	options := client.Option(func(bd *client.Request) {
		auth := bd.HTTP.Header.Get("Authorization")
		ctx := context.WithValue(context.TODO(), middleware.TokenCtxKey, auth)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	})

	c := client.New(handlr)
	auth_uid_1 := gofakeit.UUID()
	auth_uid_2 := gofakeit.UUID()
	var _user struct{ CreateUser *model.UserCreated }
	var _user2 struct{ CreateUser *model.UserCreated }
	var _agencies struct{ RetrieveAllAgnecies []*model.Agency }
	var _topupID struct{ AddTopUp string }
	t.Run("should create user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		response, err := user.CreateUserTest(c, options, idToken)
		require.Nil(t, err)
		_user = *response
	})

	t.Run("should create user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		response, err := user.CreateUserTest(c, options, idToken)
		require.Nil(t, err)
		_user2 = *response
	})

	t.Run("should load agencies", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		agencies, err := agency.LoadAlAgenciesTest(c, options, idToken)
		require.Nil(t, err)
		_agencies = agencies
	})
	t.Run("should top up | USER", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		input := model.TopUpInput{
			Amount:          100,
			Agency:          *_agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _user.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _user.CreateUser.User.ID,
		}
		id, err := transaction.InitiaTopupTest(c, options, idToken, user.UserPinCode, input)
		require.Nil(t, err)
		_topupID = id
	})
	t.Run("should cancel top up | USER", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		success, err := transaction.CancelTopUpUserTest(c, options, idToken, _topupID.AddTopUp, model.PaymentTypeTopup, user.UserPinCode)
		require.Nil(t, err)
		require.True(t, success.CancelTransactionUser)
	})
	t.Run("should initiate withdraw | USER", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		input := model.WithdrawInput{
			Amount:          100,
			Agency:          *_agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _user.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _user.CreateUser.User.ID,
		}
		_, err := transaction.InitiateWithdrawTest(c, options, idToken, user.UserPinCode, input)
		require.NotEmpty(t, err)
	})

	t.Run("should initiate a top up and validate by agent", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		input := model.TopUpInput{
			Amount:          10000,
			Agency:          *_agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _user.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _user.CreateUser.User.ID,
		}
		id, err := transaction.InitiaTopupTest(c, options, idToken, user.UserPinCode, input)
		require.Nil(t, err)

		customTokenAdmin, _ := firebase.Connect().CreateCustomToken(context.Background(), user.User_adminUIDTEST)
		idTokenAdmin := firebase.GetIdToken(customTokenAdmin)
		success, err := agent.ConfirmTopUpTest(c, options, idTokenAdmin, id.AddTopUp, model.PaymentTypeTopup, user.Admin_pinTEST, os.Getenv("DEFAULT_CURRENCY"))
		require.Nil(t, err)
		require.True(t, success.ConfirmTransactionAgent)

		wallet, err := user.LoadBalanceTest(c, options, idToken)
		require.Nil(t, err)
		require.Equal(t, *wallet.LoadBalance.Amount, float64(10000))
	})

	t.Run("initiate payment", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		success, err := transaction.SendMoneyToPayeeTest(c, options, idToken, _user2.CreateUser.User.Keypair.PublicKey, os.Getenv("DEFAULT_CURRENCY"), 100, user.UserPinCode, _user2.CreateUser.User.ID)
		require.Nil(t, err)
		require.True(t, success.CreateTransfer)
		_wallet1, err := user.LoadBalanceTest(c, options, idToken)
		require.Nil(t, err)

		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken2 := firebase.GetIdToken(customToken2)
		_wallet2, err := user.LoadBalanceTest(c, options, idToken2)
		require.Nil(t, err)
		require.Equal(t, *_wallet1.LoadBalance.Amount, float64(9899))
		require.Equal(t, *_wallet2.LoadBalance.Amount, float64(100))
	})

	t.Run("should withdraw and validate By User", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		input := model.WithdrawInput{
			Amount:          1000,
			Agency:          *_agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _user.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _user.CreateUser.User.ID,
		}
		id, err := transaction.InitiateWithdrawTest(c, options, idToken, user.UserPinCode, input)
		require.Nil(t, err)

		customTokenAdmin, _ := firebase.Connect().CreateCustomToken(context.Background(), user.User_adminUIDTEST)
		idTokenAdmin := firebase.GetIdToken(customTokenAdmin)
		success, err := agent.ConfirmWithdrawTest(c, options, idTokenAdmin, id.AddWithDraw, model.PaymentTypeWithdraw, user.Admin_pinTEST, os.Getenv("DEFAULT_CURRENCY"))
		require.Nil(t, err)
		require.True(t, success.ConfirmTransactionAgent)

		_wallet, err := user.LoadBalanceTest(c, options, idToken)
		require.Nil(t, err)
		require.Equal(t, *_wallet.LoadBalance.Amount, float64(8899))
	})

	t.Run("should load all transaction", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		_transactions, err := transaction.LoadAllUserTransactionTest(c, options, idToken)
		require.Nil(t, err)
		require.NotEmpty(t, _transactions.GetActivity)
		fmt.Println(_transactions.GetActivity[0])
	})

}
