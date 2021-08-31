package graph

import (
	"database/sql"

	redis "github.com/go-redis/redis/v8"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB    *sql.DB
	Redis *redis.Client
}
