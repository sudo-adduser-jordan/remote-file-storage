package database

import (
	"context"
	"fmt"
	"main/configs"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var dbpool *pgxpool.Pool

func ConnectToDatabase() {

	// "postgres://username:password@localhost:5432/database_name"
	DATABASE_URL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		configs.Env("DB_USER"),
		configs.Env("DB_PASSWORD"),
		configs.Env("DB_HOST"),
		configs.Env("DB_PORT"),
		configs.Env("DB_NAME"),
	)

	dbpool, err := pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Unable to create connection pool: %v\n",
			err,
		)
		os.Exit(1)
	}
	defer dbpool.Close()

}

func MigrateDatabase() {

}

func Create() {}

func Read() {}

func Update() {}

func Delete() {}
