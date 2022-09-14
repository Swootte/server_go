package testing

import (
	"context"
	"os"
	"testing"

	"server/agency"
	"server/graph/model"
	"server/middleware"
	"server/serverutils"
	"server/transaction"
	"server/user"
	"server/utils"

	firebase "server/firebase"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func init() {
	utils.LoadEnv()
}

func TestUser(t *testing.T) {
	handlr := serverutils.StartServer()
	options := client.Option(func(bd *client.Request) {
		auth := bd.HTTP.Header.Get("Authorization")
		ctx := context.WithValue(context.TODO(), middleware.TokenCtxKey, auth)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	})

	c := client.New(handlr)
	auth_uid := gofakeit.UUID()
	var _user struct{ CreateUser *model.UserCreated }
	var _agencies struct{ RetrieveAllAgnecies []*model.Agency }
	t.Run("create user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid)
		idToken := firebase.GetIdToken(customToken)
		response, err := user.CreateUserTest(c, options, idToken)
		require.NoError(t, err)
		_user = *response
	})

	t.Run("should load balance", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid)
		idToken := firebase.GetIdToken(customToken)
		wallet, err := user.LoadBalanceTest(c, options, idToken)
		require.NoError(t, err)
		_result := wallet.LoadBalance.Amount
		_output := float64(0)
		require.Equal(t, _result, &_output)
	})

	t.Run("update FCM Token", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid)
		idToken := firebase.GetIdToken(customToken)
		fcmToken := gofakeit.UUID()
		success, err := user.UpdateFcmTokenTest(c, options, idToken, fcmToken)
		require.NoError(t, err)
		require.True(t, success.UpdateFcmToken)

		data, err := user.UserExistTest(c, options, idToken)
		require.NoError(t, err)
		_result := data.UsersExist.FcmToken
		_output := fcmToken
		require.Equal(t, *_result, _output)

	})
	t.Run("should load agencies", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid)
		idToken := firebase.GetIdToken(customToken)
		agencies, err := agency.LoadAlAgenciesTest(c, options, idToken)
		require.NoError(t, err)
		require.NotEqual(t, 0, len(agencies.RetrieveAllAgnecies))
		_agencies = agencies
	})

	t.Run("Change pin code", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid)
		idToken := firebase.GetIdToken(customToken)
		newPin := gofakeit.Password(true, true, false, false, false, 10)
		success, err := user.ChangePinCodeTest(c, options, idToken, newPin)
		require.NoError(t, err)
		require.True(t, success.ChangePinCode)
		input := model.TopUpInput{
			Amount:          1000,
			Agency:          *_agencies.RetrieveAllAgnecies[0].ID,
			Token:           os.Getenv("DEFAULT_CURRENCY"),
			Destination:     _user.CreateUser.User.Keypair.PublicKey,
			DestinationUser: _user.CreateUser.User.ID,
		}
		id, err := transaction.InitiaTopupTest(c, options, idToken, newPin, input)
		require.NoError(t, err)
		require.NotEmpty(t, id)
	})
	t.Run("change profil picture", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid)
		idToken := firebase.GetIdToken(customToken)
		url := gofakeit.URL()
		success, err := user.UpdateProfilPictureTest(c, options, idToken, url)
		require.NoError(t, err)
		require.True(t, success.UpdateProfilePicture)
		data, err := user.UserExistTest(c, options, idToken)
		require.Nil(t, err)
		require.Equal(t, data.UsersExist.PhotoURL, &url)
	})
	t.Run("deleteUser", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid)
		idToken := firebase.GetIdToken(customToken)
		success, err := user.DeleteUserTest(c, options, idToken)
		require.NoError(t, err)
		require.True(t, success.DeleteUser)

	})

}
