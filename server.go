package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/rs/zerolog/log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/s-beats/graphql-server/graph"
	"github.com/s-beats/graphql-server/graph/generated"
)

const defaultPort = ":8080"
const defaultAddress = "127.0.0.1"

func logMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Debug().Str("path", r.URL.String())
	}
}

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
		log.Fatal().Err(err).Msg("Openerror")
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Msg("PingError")
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{DB: db}}))
	// extention handler
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		oc := graphql.GetOperationContext(ctx)

		// format query
		rp := regexp.MustCompile(`\n *| {2,}`)
		q := rp.ReplaceAllString(oc.RawQuery, " ")
		// trim right space
		q = strings.TrimRight(q, " ")

		log.Debug().Str("query", q).Send()
		return next(ctx)
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", logMiddleware(srv))

	log.Printf("connect to http://%s%s/ for GraphQL playground", address, port)
	log.Fatal().Err(http.ListenAndServe(address+port, nil))
}
