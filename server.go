package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/s-beats/graphql-server/graph"
	"github.com/s-beats/graphql-server/graph/generated"
)

const defaultPort = ":8080"
const defaultAddress = "127.0.0.1"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = defaultAddress
	}

	db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/database")
	if err != nil {
		log.Fatal("OpenError: ", err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("PingError: ", err)
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	// extention handler
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)
		log.Printf("around: %s %s", oc.OperationName, oc.RawQuery)
		return next(ctx)
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://%s%s/ for GraphQL playground", address, port)
	log.Fatal(http.ListenAndServe(address+port, nil))
}
