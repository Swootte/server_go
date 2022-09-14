package testing

import (
	"context"
	firebase "server/firebase"
	"server/middleware"
	"server/notification"
	"server/serverutils"
	"server/user"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/require"
)

func TestNotification(t *testing.T) {
	handlr := serverutils.StartServer()
	options := client.Option(func(bd *client.Request) {
		auth := bd.HTTP.Header.Get("Authorization")
		ctx := context.WithValue(context.TODO(), middleware.TokenCtxKey, auth)
		bd.HTTP = bd.HTTP.WithContext(ctx)
	})

	c := client.New(handlr)
	auth_uid_1 := gofakeit.UUID()

	t.Run("should create a user", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		_, err := user.CreateUserTest(c, options, idToken)
		require.Empty(t, err)
	})
	t.Run("should load all notification", func(t *testing.T) {
		customToken, _ := firebase.Connect().CreateCustomToken(context.Background(), auth_uid_1)
		idToken := firebase.GetIdToken(customToken)
		_notifications, err := notification.LoadAllNotificationTest(c, options, idToken)
		require.Empty(t, err)
		require.Equal(t, len(_notifications.LoadNotification), 0)
	})
}
