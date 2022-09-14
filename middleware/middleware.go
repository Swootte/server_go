package middleware

import (
	"context"
	"net/http"
)

// A private key for context that only this package can access. This is important
// to prevent collisions between different context uses
var TokenCtxKey = &contextKey{"token"}
var BasicTokenCtxKey = &contextKey{"basicToken"}
var RemoteAddressCTxKey = &contextKey{"RemoteAddress"}

type contextKey struct {
	name string
}

func Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), TokenCtxKey, header)
			_ctx := context.WithValue(ctx, RemoteAddressCTxKey, ip)
			r = r.WithContext(_ctx)
			next.ServeHTTP(w, r)
		})
	}
}

func BasicAuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := r.RemoteAddr
			header := r.Header.Get("Authorization")
			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), BasicTokenCtxKey, header)
			_ctx := context.WithValue(ctx, RemoteAddressCTxKey, ip)
			r = r.WithContext(_ctx)
			next.ServeHTTP(w, r)

		})
	}
}

func ForContextBasic(ctx context.Context) string {
	raw, _ := ctx.Value(BasicTokenCtxKey).(string)
	return raw
}

func ForContext(ctx context.Context) string {
	raw, _ := ctx.Value(TokenCtxKey).(string)
	return raw
}

func ForRemoteAddressContext(ctx context.Context) string {
	raw, _ := ctx.Value(RemoteAddressCTxKey).(string)
	return raw
}
