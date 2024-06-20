package main

import (
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"log"
	"net/http"
	"os"
	"post-api/internal/db"
	"post-api/internal/graph"
	"post-api/internal/repository/postgres"
	"post-api/internal/resolver"
	"post-api/internal/service"

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
	comRepo := postgres.NewCommentsRepository(pool)
	postRepo := postgres.NewPostRepository(pool)
	serv := service.NewService(postRepo, comRepo)

	res := resolver.NewResolver(serv)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: res}))
	srv.AddTransport(&transport.Websocket{})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
