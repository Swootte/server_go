package main

import (
	"log"
	"net/http"
	"os"
	"server/dataloaders"
	"server/middleware"
	"server/serverutils"
	"server/utils"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
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

	router.Use(middleware.BasicAuthMiddleware())
	srv := serverutils.StartServer()
	dataloadersSrv := dataloaders.Middleware(dataloaders.NewLoaders(), srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", dataloadersSrv)
	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
