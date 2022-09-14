package testing

import (
	"context"
	"encoding/base64"
	"os"
	"server/agency"
	"server/agent"
	"server/enterprise"
	firebase "server/firebase"
	"server/graph/model"
	"server/middleware"
	"server/qrcode"
	"server/serverutils"
	"server/transaction"
	"server/user"
	"testing"
	"time"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func getDefaultEnterprise(enterprises []*model.Enterprise) *model.Enterprise {
	for _, company := range enterprises {
		if company.DefaultEnterprise == true {
			return company
		}
	}

	output := model.Enterprise{}
	return &output
}

func TestEnterprise(t *testing.T) {
	handlr := serverutils.StartServer()
	options := client.Option(func(bd *client.Request) {
		auth := bd.HTTP.Header.Get("Authorization")
		ctx := context.WithValue(context.TODO(), middleware.TokenCtxKey, auth)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	})

	c := client.New(handlr)
	auth_uid_1 := gofakeit.UUID()
	auth_uid_2 := gofakeit.UUID()
	var _db_user1 *model.User
	// var _db_user2 *model.User
	var _agencies struct{ RetrieveAllAgnecies []*model.Agency }
	var _enterprise model.Enterprise
	var _transaction model.Paiement
	t.Run("create new user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		response, err := user.CreateUserTest(c, options, idToken)
		require.Empty(t, err)
		_db_user1 = response.CreateUser.User
	})
	t.Run("create second user and top up and validate the topup", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		_, err := user.CreateUserTest(c, options, idToken)
		require.Empty(t, err)
		// _db_user2 = response

		agencies, err := agency.LoadAlAgenciesTest(c, options, idToken)
		require.Empty(t, err)
		_agencies = agencies
		input := model.TopUpInput{
			Amount:          10000,
			Agency:          *_agencies.RetrieveAllAgnecies[0].ID,
			Token:           customToken,
			Destination:     "",
			DestinationUser: "",
		}
		result, err := transaction.InitiaTopupTest(c, options, idToken, user.UserPinCode, input)
		require.Empty(t, err)
		require.NotEmpty(t, result)

		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), user.User_adminUIDTEST)
		idToken2 := firebase.GetIdToken(customToken2)
		isOk, err := agent.ConfirmTopUpTest(c, options, idToken2, result.AddTopUp, model.PaymentTypeTopup, user.Admin_pinTEST, os.Getenv("DEFAULT_CURRENCY"))
		require.Empty(t, err)
		require.True(t, isOk.ConfirmTransactionAgent)
	})
	t.Run("should enterprise creation", func(t *testing.T) {
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
		require.Empty(t, err)
		require.NotEmpty(t, company)
		_enterprise = company.CreateEnterprise
		require.Equal(t, *company.CreateEnterprise.Name, input.Name)
		require.Equal(t, *company.CreateEnterprise.Rccm, input.Rccm)
		require.Equal(t, *company.CreateEnterprise.Email, input.Email)
		require.Equal(t, *company.CreateEnterprise.Person.FirstName, input.Person.FirstName)
		require.Equal(t, *company.CreateEnterprise.Person.LastName, input.Person.LastName)
		require.Equal(t, *company.CreateEnterprise.Person.Address, input.Person.Address)
		require.Equal(t, company.CreateEnterprise.Person.State, input.Person.State)
		require.Equal(t, *company.CreateEnterprise.TransactionLibele, input.TransactionLibele)
		require.Equal(t, *company.CreateEnterprise.AbregedLibele, input.AbregedLibele)
		require.Equal(t, *company.CreateEnterprise.Sector, input.ActivitySector)
		require.Equal(t, *company.CreateEnterprise.Phone, input.Phone)
	})
	t.Run("should load all enterprise for a user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		_enterprise, err := enterprise.GetAllEnterpriseTest(c, options, idToken)
		require.Empty(t, err)
		require.NotNil(t, _enterprise)
		require.Equal(t, len(_enterprise.GetAllUserEnterprise), 1)
	})

	t.Run("should update EnterpriseType", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		_type := "INDIVIDUAL"
		_country := "CG"
		compagnies, err := enterprise.UpdateEnterpriseTypeTest(c, options, idToken, _enterprise.ID, _type, _country)
		require.Empty(t, err)
		require.NotNil(t, compagnies)
		require.Equal(t, *compagnies.UpdateEnterpriseType[0].Type, _type)
		require.Equal(t, *compagnies.UpdateEnterpriseType[0].Country, _country)

	})

	t.Run("should update PersonnalInformation", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		first_name := gofakeit.UUID() + "first_name"
		last_name := gofakeit.UUID() + "last_name"
		email := gofakeit.Email()
		address := gofakeit.UUID() + "address"
		city := gofakeit.UUID() + "city"
		state := gofakeit.UUID() + "state"
		zip := "25000"
		compagnies, err := enterprise.UpdatePersonnalInformationTest(c, options, idToken, _enterprise.ID, first_name, last_name, email, address, city, state, zip)
		require.Empty(t, err)
		require.NotEmpty(t, compagnies)
		require.Equal(t, *compagnies.UpdatePersonnalInformation[0].Person.FirstName, first_name)
		require.Equal(t, *compagnies.UpdatePersonnalInformation[0].Person.LastName, last_name)
		require.Equal(t, *compagnies.UpdatePersonnalInformation[0].Person.Email, email)
		require.Equal(t, *compagnies.UpdatePersonnalInformation[0].Person.Address, address)
		require.Equal(t, *compagnies.UpdatePersonnalInformation[0].Person.City, city)
		require.Equal(t, *compagnies.UpdatePersonnalInformation[0].Person.State, state)
		require.Equal(t, *compagnies.UpdatePersonnalInformation[0].Person.Zip, zip)

	})

	t.Run("should update EnterpriseInformation", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		rccm := gofakeit.UUID() + "rccm"
		sector := gofakeit.UUID() + "sector"
		website := gofakeit.UUID() + "website.com"
		description := gofakeit.UUID() + "description"
		compagnies, err := enterprise.UpdateEnterpriseInformationTest(c, options, idToken, _enterprise.ID, rccm, sector, website, description)
		require.Empty(t, err)
		require.NotNil(t, compagnies)
		require.Equal(t, *compagnies.UpdateEnterpriseInformation[0].Rccm, rccm)
		require.Equal(t, *compagnies.UpdateEnterpriseInformation[0].Sector, sector)
		require.Equal(t, *compagnies.UpdateEnterpriseInformation[0].Website, website)
		require.Equal(t, *compagnies.UpdateEnterpriseInformation[0].Description, description)

	})

	t.Run("should update ExecutionInformation", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)

		sellingPyshicalGoods := true
		selfShipping := true
		shippingDelay := gofakeit.UUID() + "shippingDelay"

		compagnies, err := enterprise.UpdateExecutionInformationTest(c, options, idToken, _enterprise.ID, sellingPyshicalGoods, selfShipping, shippingDelay)
		require.Empty(t, err)
		require.NotNil(t, compagnies)
		require.Equal(t, *compagnies.UpdateExecutionInformation[0].SellingPhysicalGoods, sellingPyshicalGoods)
		require.Equal(t, *compagnies.UpdateExecutionInformation[0].SelfShippingProduct, selfShipping)
		require.Equal(t, *compagnies.UpdateExecutionInformation[0].ShippingDelay, shippingDelay)
	})

	t.Run("should update PublicInformation", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)

		name := gofakeit.UUID() + "name"
		libelle := gofakeit.UUID() + "libelle"
		libelleAbreged := gofakeit.UUID() + "libelleAbreged"
		email := gofakeit.Email()
		phone := gofakeit.Phone()

		compagnies, err := enterprise.UpdatePublicInformationTest(c, options, idToken, _enterprise.ID, name, libelle, libelleAbreged, email, phone)
		require.Empty(t, err)
		require.NotNil(t, compagnies)
		require.Equal(t, *compagnies.UpdatePublicInformation[0].Name, name)
		require.Equal(t, *compagnies.UpdatePublicInformation[0].TransactionLibele, libelle)
		require.Equal(t, *compagnies.UpdatePublicInformation[0].AbregedLibele, libelleAbreged)
		require.Equal(t, *compagnies.UpdatePublicInformation[0].Email, email)
		require.Equal(t, *compagnies.UpdatePublicInformation[0].Phone, phone)
	})

	t.Run("should get enterprise ethereum balance", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		balance, err := enterprise.GetEnterpriseBalanceTest(c, options, idToken, _enterprise.ID)
		require.Empty(t, err)
		require.Equal(t, balance.GetEnterpriseBalance, float64(0))
	})
	t.Run("should get getProfilNetChartDataTest", func(t *testing.T) {
		now := time.Now()
		yesterday := now.AddDate(0, -1, 0)
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		charteData, err := enterprise.GetProfilNetChartDataTest(c, options, idToken, _enterprise.ID, now.UTC().Format(time.RFC3339), yesterday.UTC().Format(time.RFC3339))
		require.Empty(t, err)
		require.NotNil(t, charteData)
		require.Equal(t, *charteData.GetProfilNetChartData.CurrentTotal, float64(0))
		require.Equal(t, *charteData.GetProfilNetChartData.FormerTotal, float64(0))
		require.Equal(t, len(charteData.GetProfilNetChartData.Chart), int(7))
		require.Equal(t, *charteData.GetProfilNetChartData.PourcentageDifference, float64(0))
		require.True(t, *charteData.GetProfilNetChartData.IsPositive)
	})

	t.Run("should get getProfilBrutChartDataTest", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)

		now := time.Now()
		yesterday := now.AddDate(0, -1, 0)
		charteData, err := enterprise.GetProfilBrutChartDataTest(c, options, idToken, _enterprise.ID, now.UTC().Format(time.RFC3339), yesterday.UTC().Format(time.RFC3339))
		require.Empty(t, err)
		require.NotNil(t, charteData)
		require.Equal(t, *charteData.GetProfilBrutChartData.CurrentTotal, float64(0))
		require.Equal(t, *charteData.GetProfilBrutChartData.FormerTotal, float64(0))
		require.Equal(t, len(charteData.GetProfilBrutChartData.Chart), int(7))
		require.Equal(t, *charteData.GetProfilBrutChartData.PourcentageDifference, float64(0))
		require.True(t, *charteData.GetProfilBrutChartData.IsPositive)
	})
	t.Run("should get getProfilNonCarpturedChartDataTest", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)

		now := time.Now()
		yesterday := now.AddDate(0, -1, 0)
		charteData, err := enterprise.GetProfilNonCarpturedChartDataTest(c, options, idToken, _enterprise.ID, now.UTC().Format(time.RFC3339), yesterday.UTC().Format(time.RFC3339))
		require.Empty(t, err)
		require.NotNil(t, charteData)
		require.Equal(t, *charteData.GetProfilNonCarpturedChartData.CurrentTotal, float64(0))
		require.Equal(t, *charteData.GetProfilNonCarpturedChartData.FormerTotal, float64(0))
		require.Equal(t, len(charteData.GetProfilNonCarpturedChartData.Chart), int(7))
		require.Equal(t, *charteData.GetProfilNonCarpturedChartData.PourcentageDifference, float64(0))
		require.True(t, *charteData.GetProfilNonCarpturedChartData.IsPositive)
	})

	t.Run("should get pdf", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		pdf, err := enterprise.GetPdfTest(c, options, idToken, _enterprise.ID)
		require.Empty(t, err)
		require.NotEmpty(t, pdf.GetPdf)
	})

	t.Run("should getAllTransactionByEnterpriseIdTest", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)

		now := time.Now()
		yesterday := now.AddDate(0, -1, 0)
		transactions, err := enterprise.GetAllTransactionByEnterpriseIdTest(c, options, idToken, _enterprise.ID, now.UTC().Format(time.RFC3339), yesterday.UTC().Format(time.RFC3339), 5, 0)
		require.Empty(t, err)
		require.Equal(t, len(transactions.GetAllTransactionByEnterpriseId.Transactions), 0)
	})
	t.Run("should getSuccessFullTransactionByEnterpriseIdTest", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		now := time.Now()
		yesterday := now.AddDate(0, -1, 0)
		transactions, err := enterprise.GetSuccessFullTransactionByEnterpriseIdTest(c, options, idToken, _enterprise.ID, now.UTC().Format(time.RFC3339), yesterday.UTC().Format(time.RFC3339), 5, 0)
		require.Empty(t, err)
		require.Equal(t, len(transactions.GetSuccessFullTransactionByEnterpriseId.Transactions), 0)
	})
	t.Run("should getRefundedTransactionByEnterpriseIdTest", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		now := time.Now()
		yesterday := now.AddDate(0, -1, 0)
		transactions, err := enterprise.GetRefundedTransactionByEnterpriseIdTest(c, options, idToken, _enterprise.ID, now.UTC().Format(time.RFC3339), yesterday.UTC().Format(time.RFC3339), 5, 0)
		require.Empty(t, err)
		require.Equal(t, len(transactions.GetRefundedTransactionByEnterpriseId.Transactions), 0)
	})

	t.Run("should getNonCapturedTransactionByEnterpriseIdTest", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		now := time.Now()
		yesterday := now.AddDate(0, -1, 0)
		transactions, err := enterprise.GetNonCapturedTransactionByEnterpriseIdTest(c, options, idToken, _enterprise.ID, now.UTC().Format(time.RFC3339), yesterday.UTC().Format(time.RFC3339), 5, 0)
		require.Empty(t, err)
		require.Equal(t, len(transactions.GetNonCapturedTransactionByEnterpriseId.Transactions), 0)
	})
	t.Run("should getFailedTransactionByEnterpriseIdTest", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		now := time.Now()
		yesterday := now.AddDate(0, -1, 0)
		transactions, err := enterprise.GetFailedTransactionByEnterpriseIdTest(c, options, idToken, _enterprise.ID, now.UTC().Format(time.RFC3339), yesterday.UTC().Format(time.RFC3339), 5, 0)
		require.Empty(t, err)
		require.Equal(t, len(transactions.GetFailedTransactionByEnterpriseId.Transactions), 0)
	})
	t.Run("should recreate publishableKey", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)

		compagnies, err := enterprise.RecreateTestPublishableKeyTest(c, options, idToken, user.UserPinCode, _enterprise.ID)
		require.Empty(t, err)
		require.NotEqual(t, compagnies.RecreateEnterprisePublishableKey[0].PublishableKey, _enterprise.PublishableKey)
	})
	t.Run("should recreate privateKey", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)

		compagnies, err := enterprise.RecreateTestPrivate_keyTest(c, options, idToken, user.UserPinCode, _enterprise.ID)
		require.Empty(t, err)
		require.NotEqual(t, compagnies.RecreateEnterprisePrivateKey[0].PrivateKey, _enterprise.PrivateKey)
	})

	t.Run("should remove enterprise", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		compagnies, err := enterprise.RemoveEnterpriseTest(c, options, idToken, user.UserPinCode, _enterprise.ID)
		require.Empty(t, err)
		require.Equal(t, len(compagnies.RemoveEnterprise), 0)
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
		require.Empty(t, err)
		require.NotNil(t, company)
		_enterprise = company.CreateEnterprise
	})

	t.Run("should get new commerce transaction", func(t *testing.T) {
		auth := _enterprise.PublishableKey + ":" + _enterprise.PrivateKey
		basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))
		_paiment, err := transaction.AuthenticateForPaymentTest(c, options, basicAuth, 100, "")
		require.Empty(t, err)
		_transaction = _paiment.AuthenticateForPayment
	})

	t.Run("should load commerce transaction and pay", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		_paiement, err := qrcode.GetQrOwnerPaymentTest(c, options, idToken, _transaction.ID)
		require.Empty(t, err)

		passedPayment, err := enterprise.PayUnConfirmedTransactionTest(c, options, idToken, _enterprise.ID, user.UserPinCode, _paiement.GetQrOwner.ID)
		require.Empty(t, err)
		require.NotNil(t, passedPayment)

		wallet, err := user.LoadBalanceTest(c, options, idToken)
		require.Empty(t, err)
		require.Equal(t, *wallet.LoadBalance.Amount, float64(9900))

		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken2 := firebase.GetIdToken(customToken2)
		enterpriseBalance, err := enterprise.GetEnterpriseBalanceTest(c, options, idToken2, _enterprise.ID)
		require.Empty(t, err)
		require.Equal(t, enterpriseBalance.GetEnterpriseBalance, float64(100))
	})
	t.Run("should get new commerce transaction again", func(t *testing.T) {
		auth := _enterprise.PublishableKey + ":" + _enterprise.PrivateKey
		basicAuth := base64.StdEncoding.EncodeToString([]byte(auth))
		_paiment, err := transaction.AuthenticateForPaymentTest(c, options, basicAuth, 200, "")
		require.Empty(t, err)
		_transaction = _paiment.AuthenticateForPayment
	})

	t.Run("should load commerce transaction and pay again", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		_paiement, err := qrcode.GetQrOwnerPaymentTest(c, options, idToken, _transaction.ID)
		require.Empty(t, err)

		passedPayment, err := enterprise.PayUnConfirmedTransactionTest(c, options, idToken, _enterprise.ID, user.UserPinCode, _paiement.GetQrOwner.ID)
		require.Empty(t, err)
		require.NotNil(t, passedPayment)

		wallet, err := user.LoadBalanceTest(c, options, idToken)
		require.Empty(t, err)
		require.Equal(t, *wallet.LoadBalance.Amount, float64(9700))

		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken2 := firebase.GetIdToken(customToken2)
		enterpriseBalance, err := enterprise.GetEnterpriseBalanceTest(c, options, idToken2, _enterprise.ID)
		require.Empty(t, err)
		require.Equal(t, enterpriseBalance.GetEnterpriseBalance, float64(298))
	})

	t.Run("should refund the transaction", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)

		success, err := enterprise.RefundTransactionTest(c, options, idToken, _enterprise.ID, user.UserPinCode, _transaction.ID)
		require.Empty(t, err)
		require.True(t, success.RefundTransaction)

		customToken2, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken2 := firebase.GetIdToken(customToken2)
		_wallet, err := user.LoadBalanceTest(c, options, idToken2)
		require.Empty(t, err)
		require.Equal(t, *_wallet.LoadBalance.Amount, float64(9900))

		_wallet_enterprise, err := enterprise.GetEnterpriseBalanceTest(c, options, idToken, _enterprise.ID)
		require.Empty(t, err)
		require.Equal(t, _wallet_enterprise.GetEnterpriseBalance, float64(99))
	})

	t.Run("should send money from enterprise to user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		success, err := enterprise.TransferMoneyEnterpriseTest(c, options, idToken, _enterprise.ID, user.UserPinCode, _db_user1.Keypair.PublicKey, 100)
		require.Empty(t, err)
		require.True(t, success.TransferMoneyEnterprise)

		_wallet, err := user.LoadBalanceTest(c, options, idToken)
		require.Empty(t, err)
		require.Equal(t, *_wallet.LoadBalance.Amount, float64(9900))

		_wallet_enterprise, err := enterprise.GetEnterpriseBalanceTest(c, options, idToken, _enterprise.ID)
		require.Empty(t, err)
		require.Equal(t, _wallet_enterprise.GetEnterpriseBalance, float64(48.5))
	})

	t.Run("should enterprise creation", func(t *testing.T) {
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
		require.Empty(t, err)
		require.NotNil(t, company)
		compagnies, err := enterprise.GetAllEnterpriseTest(c, options, idToken)
		require.Empty(t, err)
		require.Equal(t, getDefaultEnterprise(compagnies.GetAllUserEnterprise).ID, company.CreateEnterprise.ID)
	})

	t.Run("update default enterprise", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		compagnies, err := enterprise.ChangeDefaultEnterpriseTest(c, options, idToken, _enterprise.ID)
		require.Empty(t, err)
		require.Equal(t, getDefaultEnterprise(compagnies.ChangeDefaultEnterprise).ID, _enterprise.ID)
	})
}
