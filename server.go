package main

import (
	"log"
	"net/http"
	"os"

	"github.com/retr0h/gqle/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/retr0h/gqle/pkg/config"
	"github.com/retr0h/gqle/pkg/database"
	"github.com/sirupsen/logrus"
)

const defaultPort = "8080"

func main() {
	logger := &logrus.Logger{
		Out:       os.Stderr,
		Formatter: &logrus.JSONFormatter{PrettyPrint: true},
		Hooks:     make(logrus.LevelHooks),
		Level:     logrus.InfoLevel,
	}
	config := config.New(logger)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	logger.WithFields(logrus.Fields{
		"port": port,
	}).Info("Serving GQL API")

	db := database.New(
		config,
		logger,
	)
	db.Connect()
	db.Migrate()

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{
			Database: db.Database,
			Config:   config,
		},
	}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
