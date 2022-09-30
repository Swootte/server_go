package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"server/dataloaders"
	"server/middleware"
	"server/serverutils"
	"server/user"
	"server/utils"

	"github.com/99designs/gqlgen/graphql"
	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

const defaultPort = "8080"

func init() {
	utils.LoadEnv()
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	router := chi.NewRouter()

	router.Use(middleware.Middleware())
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{},
		AllowedMethods:   []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Access-Control-Allow-Credentials", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "Apollographql-Client-Name", "Authorization", "Content-Type", "Mode"},
		AllowCredentials: true,
	}).Handler)

	router.Use(middleware.BasicAuthMiddleware())
	srv := serverutils.StartServer()
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		graphql.GetOperationContext(ctx).DisableIntrospection = true
		return next(ctx)
	})

	dataloadersSrv := dataloaders.Middleware(dataloaders.NewLoaders(), srv)
	//router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", dataloadersSrv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	user.MigrateAllUsersWallet()
}
