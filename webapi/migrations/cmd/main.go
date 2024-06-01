package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Connect("postgres", os.Getenv("DB"))
	if err != nil {
		panic(err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		panic(err.Error())
	}

	if err := goose.UpContext(context.Background(), db.DB, "/app/sql"); err != nil {
		fmt.Println("Migration failed: ", err)
		panic(err)

	}

	fmt.Println("Successfully migrated database")
	os.Exit(0)
}
