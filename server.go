package main

import (
	"log"
	"net/http"
	"os"
	"context"
	// "debug"
	"errors"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/amar-jay/todolist_v2/src/resolver"
	"github.com/amar-jay/todolist_v2/src/generated"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	 srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{}}))
	 srv.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
	//	 send this panic somewhere
	 	log.Print(err)
	 	// debug.PrintStack()
	 	return errors.New("user message on panic")
	 })


	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	 http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
