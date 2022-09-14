package dataloaders

import (
	"context"
	"net/http"
	"server/utils"

	"github.com/graph-gophers/dataloader"
)

func init() {
	utils.LoadEnv()
}

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	UserLoader       *dataloader.Loader
	EnterpriseLoader *dataloader.Loader
	AgencyLoader     *dataloader.Loader
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders() *Loaders {
	loaders := &Loaders{
		UserLoader:       dataloader.NewBatchedLoader(GetUsers),
		EnterpriseLoader: dataloader.NewBatchedLoader(GetEnterprises),
		AgencyLoader:     dataloader.NewBatchedLoader(GetAgencies),
	}
	return loaders
}

// Middleware injects data loaders into the context
func Middleware(loader *Loaders, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loader)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}
