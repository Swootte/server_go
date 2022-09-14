package utils

import (
	"context"
	"server/graph/model"
)

var UIDCtxKey = &contextKey{"uid"}
var UserCtxKey = &contextKey{"user"}
var EnterpriseCtxKey = &contextKey{"enterprise"}

type contextKey struct {
	name string
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContextUser(ctx context.Context, key *contextKey) string {
	raw, _ := ctx.Value(key).(string)
	return raw
}

func ForContextDBUser(ctx context.Context, key *contextKey) *model.User {
	raw, _ := ctx.Value(key).(*model.User)
	return raw
}

func ForContextDBEnnterprise(ctx context.Context, key *contextKey) *model.Enterprise {
	raw, _ := ctx.Value(key).(*model.Enterprise)
	return raw
}
