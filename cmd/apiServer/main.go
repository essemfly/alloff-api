package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/go-chi/chi"
	"github.com/lessbutter/alloff-api/api/apiServer"
	"github.com/lessbutter/alloff-api/api/apiServer/middleware"
	"github.com/lessbutter/alloff-api/api/apiServer/resolver"
	"github.com/lessbutter/alloff-api/config"
	"github.com/lessbutter/alloff-api/internal/storage/mongo"
	"github.com/lessbutter/alloff-api/internal/storage/postgres"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func main() {
	conf := config.GetConfiguration()
	log.Println(conf)

	conn := mongo.NewMongoDB(conf)
	conn.RegisterRepos()

	pgconn := postgres.NewPostgresDB(conf)
	pgconn.RegisterRepos()

	// (TODO) Be Refactored
	config.InitIamPort(conf)
	config.InitSlack(conf)

	port := os.Getenv("PORT")
	if port == "" {
		port = strconv.Itoa(conf.PORT)
	}

	router := chi.NewRouter()
	router.Use(middleware.Middleware())

	srv := handler.NewDefaultServer(apiServer.NewExecutableSchema(apiServer.Config{Resolvers: &resolver.Resolver{}}))
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		sentry.CaptureException(e)
		err := graphql.DefaultErrorPresenter(ctx, e)
		return err
	})

	sentryHandler := sentryhttp.New(sentryhttp.Options{
		Repanic: true,
	})
	gqlgenHandler := playground.Handler("Outlet API", "/query")

	router.Handle("/", sentryHandler.Handle(gqlgenHandler))
	router.Handle("/query", sentryHandler.Handle(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
