package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/chack-check/organizations-service/infrastructure/api/graph"
	"github.com/chack-check/organizations-service/infrastructure/api/middlewares"
	"github.com/chack-check/organizations-service/infrastructure/database"
	"github.com/chack-check/organizations-service/infrastructure/settings"
	"github.com/go-chi/chi"
)

func RunApi() {
	// defer rabbit.EventsRabbitConnection.Close()
	// defer redisdb.RedisConnection.Close()

	database.SetupMigrations()

	router := chi.NewRouter()

	router.Use(middlewares.UserMiddleware)
	router.Use(middlewares.CorsMiddleware)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/api/v1/organizations", playground.Handler("GraphQL playground", "/api/v1/organizations/query"))
	router.Handle("/api/v1/organizations/query", srv)

	listen := fmt.Sprintf(":%d", settings.Settings.APP_PORT)
	log.Fatal(http.ListenAndServe(listen, router))
}
