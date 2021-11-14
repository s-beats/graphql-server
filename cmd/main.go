package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"

	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/go-sql-driver/mysql"
	"github.com/s-beats/graphql-todo/handler"
	"github.com/s-beats/graphql-todo/repository"
	"github.com/s-beats/graphql-todo/usecase"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
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
	dbHost := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	dbUser := os.Getenv("DATABASE_USER")
	dbPass := os.Getenv("DATABASE_PASS")
	dbName := os.Getenv("DATABASE_NAME")
	dsnName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsnName)
	if err != nil {
		log.Fatal().Err(err).Str("dsn-name:", dsnName).Msg("Open mysql error")
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal().Err(err).Str("dsn-name:", dsnName).Msg("Ping mysql error")
	}
	bunDB := bun.NewDB(db, mysqldialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	taskRepo := repository.NewTask(bunDB)
	taskUsecase := usecase.NewTask(taskRepo)

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", logMiddleware(handler.GraphQLHandler(taskUsecase)))

	address := os.Getenv("HOST")
	port := os.Getenv("PORT")
	log.Fatal().Err(http.ListenAndServe(address+":"+port, nil)).Send()
}
