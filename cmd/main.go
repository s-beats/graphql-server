package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/s-beats/graphql-todo/domain/repository"
	"github.com/s-beats/graphql-todo/handler"
	"github.com/s-beats/graphql-todo/infra/rdb"
	"github.com/s-beats/graphql-todo/usecase"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg(("Error loading .env file"))
	}
}

// FIXME:
func logMiddleware(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
		log.Debug().Str("path", r.URL.String())
	}
}

func main() {
	db, err := rdb.NewDB()
	if err != nil {
		log.Fatal().Err(err)
	}

	taskRepo := repository.NewTask(db)
	taskUsecase := usecase.NewTask(taskRepo)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", logMiddleware(handler.GraphQLHandler(taskUsecase)))

	address := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Fatal().Err(http.ListenAndServe(address+":"+port, nil)).Send()
}
