package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	// Import the postgres driver
	_ "github.com/jackc/pgx/v5/stdlib"
	// _ "github.com/lib/pq"
)

type Input struct {
	DSN             string
	SetMaxIdeConns  int
	SetMaxOpenConns int
}

func CreatePostgreSQL(input *Input) *sqlx.DB {
	db, err := sqlx.Connect("postgres", input.DSN)
	if err != nil {
		panic(err.Error())
	}

	db.SetMaxIdleConns(input.SetMaxIdeConns)
	db.SetMaxOpenConns(input.SetMaxOpenConns)

	// Ping the database to ensure the connection is working
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(err.Error())
	}

	return db
}
