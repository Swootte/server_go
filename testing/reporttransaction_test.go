package testing

import (
	"context"
	"server/agency"
	firebase "server/firebase"
	"server/graph/model"
	"server/middleware"
	"server/reporttransaction"
	"server/serverutils"
	"server/transaction"
	"server/user"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestReportTransaction(t *testing.T) {
	handlr := serverutils.StartServer()
	options := client.Option(func(bd *client.Request) {
		auth := bd.HTTP.Header.Get("Authorization")
		ctx := context.WithValue(context.TODO(), middleware.TokenCtxKey, auth)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	})

	c := client.New(handlr)
	auth_uid_1 := gofakeit.UUID()
	_user := struct{ CreateUser *model.UserCreated }{}
	var _agency *model.Agency
	t.Run("should create a user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		response, err := user.CreateUserTest(c, options, idToken)
		require.Empty(t, err)
		_user = *response
	})

	t.Run("should load agencies", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		agencies, err := agency.LoadAlAgenciesTest(c, options, idToken)
		require.Empty(t, err)
		_agency = agencies.RetrieveAllAgnecies[0]
	})

	t.Run("should create a transaction and report the transaction", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		input := model.TopUpInput{
			Amount:          0,
			Agency:          *_agency.ID,
			Token:           customToken,
			Destination:     _user.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _user.CreateUser.User.ID,
		}
		id, err := transaction.InitiaTopupTest(c, options, idToken, user.UserPinCode, input)
		require.Empty(t, err)

		success, err := reporttransaction.ReportTransactionTest(c, options, idToken, id.AddTopUp, "message")
		require.Empty(t, err)
		require.True(t, success.ReportTransaction)
	})

}
