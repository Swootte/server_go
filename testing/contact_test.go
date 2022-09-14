package testing

import (
	"context"
	"server/contact"
	firebase "server/firebase"
	"server/graph/model"
	"server/middleware"
	"server/serverutils"
	"server/user"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestContact(t *testing.T) {
	handlr := serverutils.StartServer()
	options := client.Option(func(bd *client.Request) {
		auth := bd.HTTP.Header.Get("Authorization")
		ctx := context.WithValue(context.TODO(), middleware.TokenCtxKey, auth)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	})

	c := client.New(handlr)
	auth_uid_1 := gofakeit.UUID()
	auth_uid_2 := gofakeit.UUID()
	auth_uid_3 := gofakeit.UUID()
	auth_uid_4 := gofakeit.UUID()
	auth_uid_5 := gofakeit.UUID()
	auth_uid_6 := gofakeit.UUID()
	var _contacts []model.UserSmall
	t.Run("should create user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		_, err := user.CreateUserTest(c, options, idToken)
		require.Nil(t, err)
	})
	t.Run("should create user to search 0", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_2)
		idToken := firebase.GetIdToken(customToken)
		_, err := user.CreateUserTest(c, options, idToken)
		require.Nil(t, err)
	})
	t.Run("should create user to search 1", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_3)
		idToken := firebase.GetIdToken(customToken)
		_, err := user.CreateUserTest(c, options, idToken)
		require.Nil(t, err)
	})
	t.Run("should create user to search 2", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_4)
		idToken := firebase.GetIdToken(customToken)
		_, err := user.CreateUserTest(c, options, idToken)
		require.Nil(t, err)
	})
	t.Run("should create user to search 3", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_5)
		idToken := firebase.GetIdToken(customToken)
		_, err := user.CreateUserTest(c, options, idToken)
		require.Nil(t, err)
	})
	t.Run("should create user to search 4", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_6)
		idToken := firebase.GetIdToken(customToken)
		_, err := user.CreateUserTest(c, options, idToken)
		require.Nil(t, err)
	})
	t.Run("loadAllContacts", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		contacts, err := contact.LoadAllContactsTest(c, options, idToken, "")
		require.Nil(t, err)
		require.Equal(t, len(contacts.GetAllUserContact), 0)
	})
	t.Run("should search contact not added", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		contacts, err := contact.LoadNonAddedContactTest(c, options, idToken, "")
		require.Nil(t, err)
		require.Equal(t, len(contacts.GetAllContactNotAdded), 15)
		_contacts = contacts.GetAllContactNotAdded
	})

	t.Run("addNewcontact", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		success, err := contact.AddContactTest(c, options, idToken, _contacts[0].ID)
		require.Nil(t, err)
		require.True(t, success.CreateContact)

	})
	t.Run("should search contact not added already in my contact", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		contacts, err := contact.LoadAllContactsTest(c, options, idToken, *_contacts[0].LastName)
		require.Nil(t, err)
		require.Equal(t, len(contacts.GetAllUserContact), 1)
		_contacts = contacts.GetAllUserContact
	})

	t.Run("removeContact", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		success, err := contact.RemoveContactTest(c, options, idToken, _contacts[0].ID)
		require.Nil(t, err)
		require.True(t, success.RemoveContact)
		contacts, err := contact.LoadAllContactsTest(c, options, idToken, "")
		require.Nil(t, err)
		require.Equal(t, len(contacts.GetAllUserContact), 0)
	})

	t.Run("should search contact not added", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		contacts, err := contact.LoadNonAddedContactTest(c, options, idToken, *_contacts[0].FirstName)
		require.Nil(t, err)
		require.Equal(t, len(contacts.GetAllContactNotAdded), 1)
		_contacts = contacts.GetAllContactNotAdded
	})

}
