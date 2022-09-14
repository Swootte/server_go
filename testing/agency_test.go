package testing

import (
	"context"
	"server/agency"
	firebase "server/firebase"
	"server/graph/model"
	"server/middleware"
	"server/serverutils"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/stretchr/testify/require"
)

func TestAgency(t *testing.T) {
	handlr := serverutils.StartServer()
	options := client.Option(func(bd *client.Request) {
		auth := bd.HTTP.Header.Get("Authorization")
		ctx := context.WithValue(context.TODO(), middleware.TokenCtxKey, auth)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	})

	c := client.New(handlr)

	customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), "auth_uid")
	idToken := firebase.GetIdToken(customToken)
	t.Run("create agency", func(t *testing.T) {
		input := model.AgencyInpyt{
			Title:   "",
			Address: "",
			City:    "",
			Country: "",
			Token:   "",
		}
		agency, err := agency.CreateAgencyTest(c, options, idToken, "", input)
		require.Empty(t, err)
		require.NotEmpty(t, agency)

	})

	t.Run("load all agency", func(t *testing.T) {
		response, err := agency.LoadAlAgenciesTest(c, options, idToken)
		require.Empty(t, err)
		require.NotEmpty(t, response)
	})
}
