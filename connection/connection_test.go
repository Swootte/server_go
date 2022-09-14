package connection_test

import (
	"context"
	"server/connection"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
)

func TestConnection(t *testing.T) {
	t.Run("Test warning of a new connection", func(t *testing.T) {
		ipAdress := gofakeit.IPv4Address()
		connection.WarnOfNewConnectionOccured(context.Background(), ipAdress, gofakeit.UUID())
	})
}
