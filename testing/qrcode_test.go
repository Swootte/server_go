package testing

import (
	"context"
	"encoding/base64"
	"server/enterprise"
	firebase "server/firebase"
	"server/graph/model"
	"server/middleware"
	"server/qrcode"
	"server/serverutils"
	"server/transaction"
	"server/user"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestQRCode(t *testing.T) {
	handlr := serverutils.StartServer()
	options := client.Option(func(bd *client.Request) {
		auth := bd.HTTP.Header.Get("Authorization")
		ctx := context.WithValue(context.TODO(), middleware.TokenCtxKey, auth)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	})

	c := client.New(handlr)
	auth_uid_1 := gofakeit.UUID()
	auth_uid_2 := gofakeit.UUID()
	// var _user model.User
	var _user2 struct{ CreateUser *model.UserCreated }
	var _enterprise model.Enterprise
	var _payment model.Paiement
	t.Run("create new user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		_, err := user.CreateUserTest(c, options, idToken)
		require.Nil(t, err)
		// _user = *response
	})
	t.Run("create user 2", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		response, err := user.CreateUserTest(c, options, idToken)
		require.Nil(t, err)
		_user2 = *response
	})

	t.Run("should getQRcode | User", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		__user, err := qrcode.GetQrOwnerUserTest(c, options, idToken, _user2.CreateUser.User.Keypair.PublicKey)
		require.Nil(t, err)
		require.Equal(t, __user.GetQrOwner.FirstName, _user2.CreateUser.User.FirstName)
		require.Equal(t, __user.GetQrOwner.LastName, _user2.CreateUser.User.LastName)
		require.Equal(t, __user.GetQrOwner.Keypair.PublicKey, _user2.CreateUser.User.Keypair.PublicKey)
	})

	t.Run("should create another enterprise", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		img_ := gofakeit.ImageURL(300, 400)
		url_ := gofakeit.URL()
		description_ := gofakeit.LoremIpsumWord()
		input := model.EnterpriseInput{
			Name:    gofakeit.Company(),
			Country: gofakeit.Country(),
			Address: &model.AdressInput{
				Title: &gofakeit.Address().Address,
				Location: &model.LocationInput{
					Latitude:  new(float64),
					Longitude: new(float64),
				},
				IsChosed: new(bool),
				City:     &gofakeit.Address().City,
			},
			LogoURL:        &img_,
			Website:        &url_,
			Type:           "INDIVIDUAL",
			Rccm:           gofakeit.UUID() + "rccm",
			ActivitySector: "Agroalimentaire",
			Person: &model.PersonInput{
				FirstName: gofakeit.FirstName(),
				LastName:  gofakeit.LastName(),
				Email:     gofakeit.Email(),
				Address:   gofakeit.Address().Address,
				City:      gofakeit.City(),
				Zip:       &gofakeit.Address().Zip,
				State:     &gofakeit.Address().State,
			},
			Description:          &description_,
			SellingPhysicalGoods: new(bool),
			SelfShippingProduct:  new(bool),
			ShippingDelay:        new(string),
			TransactionLibele:    gofakeit.CompanySuffix(),
			AbregedLibele:        gofakeit.CompanySuffix(),
			Phone:                gofakeit.Phone(),
			Email:                gofakeit.Email(),
		}
		company, err := enterprise.CreateEnterpriseTest(c, options, idToken, input)
		require.Nil(t, err)
		_enterprise = company.CreateEnterprise
	})

	t.Run("should getQRcode | Enterprise", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		company, err := qrcode.GetQrOwnerTest(c, options, idToken, _enterprise.WalletPublicKey)
		require.Nil(t, err)
		require.Equal(t, _enterprise.Name, company.GetQrOwner.Name)
		require.Equal(t, _enterprise.ID, company.GetQrOwner.ID)
	})
	t.Run("should get new commerce transaction", func(t *testing.T) {
		auth := _enterprise.PublishableKey + ":" + _enterprise.PrivateKey
		basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))
		payment, err := transaction.AuthenticateForPaymentTest(c, options, basicAuth, 100, "blabla")
		require.Nil(t, err)
		require.NotNil(t, payment.AuthenticateForPayment.ID)
		_payment = payment.AuthenticateForPayment
	})

	t.Run("should getQRcode | Paiement", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		payment, err := qrcode.GetQrOwnerPaymentTest(c, options, idToken, _payment.ID)
		require.Nil(t, err)
		require.NotNil(t, payment)

	})
}
