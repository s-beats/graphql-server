package rdb

import (
	"database/sql"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"
)

func NewQueries() (*Queries, error) {
	db, err := sql.Open("mysql", getDataSourceName())
	if err != nil {
		return nil, err
	}
	logger := zerolog.New(
		zerolog.ConsoleWriter{Out: os.Stdout, NoColor: false},
	)
	db = sqldblogger.OpenDriver(
		getDataSourceName(),
		db.Driver(),
		zerologadapter.New(logger),
	)
	return New(db), nil
}

func getDataSourceName() string {
	c := mysql.Config{
		Net:                  "tcp",
		Addr:                 os.Getenv("MYSQL_HOST"),
		User:                 os.Getenv("MYSQL_USER"),
		Passwd:               os.Getenv("MYSQL_PASSWORD"),
		DBName:               os.Getenv("MYSQL_NAME"),
		Loc:                  time.UTC,
		Collation:            "utf8mb4_general_ci",
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	return c.FormatDSN()
}
