package serverutils

import (
	"context"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"server/enterprise"
	"server/graph"
	"server/graph/generated"
	"server/graph/model"
	"server/middleware"
	"server/user"
	"server/utils"
	"strings"
	"time"

	firebase "server/firebase"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
)

const defaultPort = "8080"

func init() {
	utils.LoadEnv()
}

func contains(s []*model.Role, str *model.Role) bool {
	for _, v := range s {
		if v.String() == str.String() {
			return true
		}
	}
	return false
}

func StartServer() *handler.Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = func(ctx context.Context, obj interface{}, next graphql.Resolver, requires *model.Role) (res interface{}, err error) {
		token := middleware.ForContext(ctx)
		if token == "" {
			return nil, fmt.Errorf("access denied, please connect to retrieve this information")
		}

		bearer := "Bearer"
		_token := token[len(bearer)+1:]

		authToken, err := firebase.Connect().VerifyIdToken(ctx, _token)
		if err != nil {
			return nil, err
		}

		_user, err := user.GetUserByFirebaseId(authToken.UID)
		if err != nil {
			return nil, err
		}

		if requires.String() == model.RoleOwner.String() {
			switch v := obj.(type) {
			case *model.UserCreated:
				uid := v.User.FirebaseUID
				if *uid == authToken.UID {
					return v.User, nil
				}
				return nil, fmt.Errorf("access denied, please connect to retrieve this information")
			case *model.User:
				uid := v.FirebaseUID
				if *uid == authToken.UID {
					return v, nil
				}
				return nil, fmt.Errorf("access denied, please connect to retrieve this information")
			case nil:
				if *_user.FirebaseUID == authToken.UID {
					return _user, nil
				}
				return nil, fmt.Errorf("access denied, please connect to retrieve this information")

			case *model.Enterprise:
				creator := v.Creator
				if _user.ID == creator {
					return v, nil
				}
				return nil, fmt.Errorf("access denied, you are not the creator of this enteeerprise")

			case []*model.Enterprise:
				var _all []*model.Enterprise
				for i := 0; i < len(v); i++ {
					creator := v[i].Creator
					if creator == _user.ID {
						_all = append(_all, v[i])
					}
				}

				return _all, nil

			}
		}

		contain := contains(_user.Permissions, requires)

		if !contain {
			return nil, fmt.Errorf("access denied")
		}

		newctx := context.WithValue(ctx, utils.UserCtxKey, _user)

		return next(newctx)

	}

	c.Directives.AuthCommerce = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		token := middleware.ForContextBasic(ctx)
		if token == "" {
			return nil, fmt.Errorf("access denied")
		}

		bearer := "Basic"
		_token := token[len(bearer)+1:]
		_new_string := strings.Split(_token, ":")
		_enterprise, err := enterprise.KeysExist(ctx, _new_string[0], _new_string[1])
		if err != nil || _enterprise == nil {
			return nil, fmt.Errorf("access denied")
		}

		newctx := context.WithValue(ctx, utils.EnterpriseCtxKey, _enterprise)
		return next(newctx)
	}

	c.Directives.AuthWithPin = func(ctx context.Context, obj interface{}, next graphql.Resolver, requires *model.Role) (res interface{}, err error) {
		token := middleware.ForContext(ctx)
		if token == "" {
			return nil, fmt.Errorf("access denied")
		}

		bearer := "Bearer"
		_token := token[len(bearer)+1:]

		authToken, err := firebase.Connect().VerifyIdToken(ctx, _token)
		if err != nil {
			return nil, err
		}

		_user, err := user.GetUserByFirebaseId(authToken.UID)
		if err != nil {
			return nil, err
		}

		contain := contains(_user.Permissions, requires)

		if !contain {
			return nil, fmt.Errorf("access denied")
		}

		newctx := context.WithValue(ctx, utils.UserCtxKey, _user)
		return next(newctx)

	}

	c.Directives.TokenVerify = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		token := middleware.ForContext(ctx)
		if token == "" {
			return nil, fmt.Errorf("access denied")
		}

		bearer := "Bearer"
		_token := token[len(bearer)+1:]

		authToken, err := firebase.Connect().VerifyIdToken(ctx, _token)
		if err != nil {
			return nil, err
		}

		newctx := context.WithValue(ctx, utils.UIDCtxKey, authToken.UID)
		return next(newctx)
	}

	c.Directives.AuthCommerce = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		token := middleware.ForContext(ctx)
		if token == "" {
			return nil, fmt.Errorf("access denied")
		}

		bearer := "Basic"
		_token := token[len(bearer)+1:]
		_trim := strings.Trim(_token, " ")
		decoded, err := base64.StdEncoding.DecodeString(_trim)
		if err != nil {
			return nil, err
		}

		apis := strings.Split(string(decoded[:]), ":")

		_enterprise, err := enterprise.KeysExist(ctx, apis[0], apis[1])
		if err != nil {
			return nil, err
		}

		if _enterprise == nil {
			return nil, fmt.Errorf("access denied")
		}

		newctx := context.WithValue(ctx, utils.EnterpriseCtxKey, _enterprise)
		return next(newctx)

	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	return srv
}
