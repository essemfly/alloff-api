package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi"
	"github.com/lessbutter/alloff-api/api/apiServer"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/resolver"
	"github.com/lessbutter/alloff-api/cmd"
	"github.com/rs/cors"
)

var (
	GitInfo   = "no info"
	BuildTime = "no datetime"
	Env       = "dev"
)

func main() {
	fmt.Println("Git commit information: ", GitInfo)
	fmt.Println("Build date, time: ", BuildTime)

	conf := cmd.SetBaseConfig(Env)

	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(conf.PORT)
	}

	router := chi.NewRouter()
	router.Use(middleware.Middleware())
	router.Use(cors.AllowAll().Handler)

	srv := handler.NewDefaultServer(apiServer.NewExecutableSchema(apiServer.Config{Resolvers: &resolver.Resolver{}}))
	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})
	gqlgenHandler := playground.Handler("Outlet API", "/query")

	router.Handle("/", sentryHandler.Handle(gqlgenHandler))
	router.Handle("/query", sentryHandler.Handle(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
