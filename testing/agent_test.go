package testing

import (
	"context"
	"os"
	"server/agency"
	"server/agent"
	firebase "server/firebase"
	"server/graph/model"
	"server/middleware"
	"server/serverutils"
	"server/transaction"
	"server/user"
	"server/utils"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func init() {
	utils.LoadEnv()
}

func TestAgent(t *testing.T) {
	handlr := serverutils.StartServer()
	options := client.Option(func(bd *client.Request) {
		auth := bd.HTTP.Header.Get("Authorization")
		ctx := context.WithValue(context.TODO(), middleware.TokenCtxKey, auth)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	})

	c := client.New(handlr)
	auth_uid_1 := gofakeit.UUID()
	auth_uid_2 := gofakeit.UUID()

	var _db_user1 struct{ CreateUser *model.UserCreated }
	var _db_user2 struct{ CreateUser *model.UserCreated }
	var agencies struct{ RetrieveAllAgnecies []*model.Agency }

	t.Run("it should create user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		response, err := user.CreateUserTest(c, options, idToken)
		require.Empty(t, err)
		_db_user1 = *response
	})
	t.Run("it should create a second user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		response, err := user.CreateUserTest(c, options, idToken)
		require.Empty(t, err)
		_db_user2 = *response
	})
	t.Run("it should grant role", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), user.User_adminUIDTEST)
		idToken := firebase.GetIdToken(customToken)
		success, err := agent.GrantRoleTest(c, options, idToken, _db_user1.CreateUser.User.ID, model.RoleAgentRole.String(), user.Admin_pinTEST, os.Getenv("DEFAULT_CURRENCY"))
		require.Empty(t, err)
		require.True(t, success.AssignRole)
	})
	t.Run("should load agencies", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		result, err := agency.LoadAlAgenciesTest(c, options, idToken)
		require.Empty(t, err)
		require.NotEmpty(t, result)
		agencies = result

	})
	t.Run("user 2 should top up and use1 should validate", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		input := model.TopUpInput{
			Amount:          1000,
			Agency:          *agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _db_user2.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _db_user2.CreateUser.User.ID,
		}
		result, err := transaction.InitiaTopupTest(c, options, idToken, user.UserPinCode, input)
		require.Empty(t, err)
		require.NotEmpty(t, result)
		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken2 := firebase.GetIdToken(customToken2)
		isOk, err := agent.ConfirmTopUpTest(c, options, idToken2, result.AddTopUp, model.PaymentTypeTopup, user.UserPinCode, os.Getenv("DEFAULT_CURRENCY"))
		require.Empty(t, err)
		require.True(t, isOk.ConfirmTransactionAgent)
	})
	t.Run("user 2 should withdraw and user 1 should validate", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		input := model.WithdrawInput{
			Amount:          100,
			Agency:          *agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _db_user2.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _db_user2.CreateUser.User.ID,
		}
		result, err := transaction.InitiateWithdrawTest(c, options, idToken, user.UserPinCode, input)
		require.Empty(t, err)
		require.NotEmpty(t, result)

		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken2 := firebase.GetIdToken(customToken2)
		isOk, err := agent.ConfirmWithdrawTest(c, options, idToken2, result.AddWithDraw, model.PaymentTypeWithdraw, user.UserPinCode, os.Getenv("DEFAULT_CURRENCY"))
		require.Empty(t, err)
		require.True(t, isOk.ConfirmTransactionAgent)
	})

	t.Run("user 2 should top up and user1 should cancel", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		input := model.TopUpInput{
			Amount:          100,
			Agency:          *agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _db_user2.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _db_user2.CreateUser.User.ID,
		}
		result, err := transaction.InitiaTopupTest(c, options, idToken, user.UserPinCode, input)
		require.Empty(t, err)
		require.NotEmpty(t, result)

		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken2 := firebase.GetIdToken(customToken2)
		isOk, err := agent.CancelTopUpTest(c, options, idToken2, result.AddTopUp, model.PaymentTypeTopup.String(), user.UserPinCode)
		require.Empty(t, err)
		require.True(t, isOk.CancelTransactionAgent)
	})

	t.Run("user 2 should withdraw and user1 should cancel", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		input := model.WithdrawInput{
			Amount:          100,
			Agency:          *agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _db_user2.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _db_user2.CreateUser.User.ID,
		}

		result, err := transaction.InitiateWithdrawTest(c, options, idToken, user.UserPinCode, input)
		require.Empty(t, err)
		require.NotEmpty(t, result)

		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken2 := firebase.GetIdToken(customToken2)
		isOk, err := agent.CancelWithDrawTest(c, options, idToken2, result.AddWithDraw, model.PaymentTypeWithdraw.String(), user.UserPinCode)
		require.Empty(t, err)
		require.True(t, isOk.CancelTransactionAgent)
	})
	t.Run("it should remove role and test again", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), user.User_adminUIDTEST)
		idToken := firebase.GetIdToken(customToken)
		result, err := agent.RemoveRoleTest(c, options, idToken, _db_user1.CreateUser.User.ID, model.RoleAgentRole.String(), user.Admin_pinTEST, os.Getenv("DEFAULT_CURRENCY"))
		require.Empty(t, err)
		require.NotEmpty(t, result)
	})

	t.Run("it should throw error when trying to validate topup", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		input := model.TopUpInput{
			Amount:          100,
			Agency:          *agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _db_user2.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _db_user2.CreateUser.User.ID,
		}
		result, err := transaction.InitiaTopupTest(c, options, idToken, user.UserPinCode, input)
		require.Empty(t, err)
		require.NotEmpty(t, result)

		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken2 := firebase.GetIdToken(customToken2)
		isOk, err := agent.ConfirmTopUpTest(c, options, idToken2, result.AddTopUp, model.PaymentTypeTopup, user.UserPinCode, os.Getenv("DEFAULT_CURRENCY"))
		require.False(t, isOk.ConfirmTransactionAgent)
		require.Error(t, err)
	})

	t.Run("it should throw error when trying to validate withdraw", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		input := model.WithdrawInput{
			Amount:          100,
			Agency:          *agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _db_user2.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _db_user2.CreateUser.User.ID,
		}
		result, err := transaction.InitiateWithdrawTest(c, options, idToken, user.UserPinCode, input)
		require.Empty(t, err)
		require.NotEmpty(t, result)

		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken2 := firebase.GetIdToken(customToken2)
		isOk, err := agent.ConfirmWithdrawTest(c, options, idToken2, result.AddWithDraw, model.PaymentTypeWithdraw, user.UserPinCode, os.Getenv("DEFAULT_CURRENCY"))
		require.False(t, isOk.ConfirmTransactionAgent)
		require.Error(t, err)
	})
}
