package database

import (
	"context"
	"fmt"
	"log"
	"main/configs"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sudo-adduser-jordan/Toolchain/Go/styles"
)

var connection *pgxpool.Pool

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

	var err error
	connection, err = pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Unable to connection to database: %v\n",
			err,
		)
		os.Exit(1)
	}
	fmt.Print("	-----> ")
	fmt.Println(styles.GreenText("Connect to database."))
}

func MigrateDatabase() {
	_, err := connection.Exec(context.Background(),
		CREATE_USER_TABLE,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("	-----> ")
	fmt.Println(styles.BlueText("Database migrated."))
}
