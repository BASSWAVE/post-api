package main

import (
	"log"
	"net/http"
	"os"
	"post-api/internal/db"
	"post-api/internal/graph"
	postgres2 "post-api/internal/repository/postgres"
	"post-api/internal/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	pool, err := db.NewPool()
	if err != nil {
		log.Fatal(err)
	}
	comRepo := postgres2.NewCommentsRepository(pool)
	postRepo := postgres2.NewPostRepository(pool)

	res := resolver.NewResolver(postRepo, comRepo)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: res}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
