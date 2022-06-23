package main

import (
	"log"
	"net/http"

	chimiddleware "github.com/go-chi/chi/middleware"
	"github.com/spf13/viper"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi"
	"github.com/lessbutter/alloff-api/api/apiServer"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/resolver"
	"github.com/lessbutter/alloff-api/cmd"
	"github.com/lessbutter/alloff-api/cmd/apiServer/webhooks"
	"github.com/rs/cors"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func main() {
	cmd.SetBaseConfig()
	port := viper.GetString("PORT")

	router := chi.NewRouter()
	router.Use(chimiddleware.RequestID)
	router.Use(middleware.Middleware())
	router.Use(cors.AllowAll().Handler)

	srv := handler.NewDefaultServer(apiServer.NewExecutableSchema(apiServer.Config{Resolvers: &resolver.Resolver{}}))
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})
	gqlgenHandler := playground.Handler("Alloff API", "/query")

	router.Handle("/", sentryHandler.Handle(gqlgenHandler))
	router.Handle("/query", sentryHandler.Handle(srv))
	router.Handle("/webhook/payment", sentryHandler.Handle(webhooks.Handler(webhooks.IamportHandler)))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
