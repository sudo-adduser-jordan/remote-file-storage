package database

import (
	"context"
	"fmt"
	"log"
	"main/configs"
	"os"

	"github.com/jackc/pgx/v5"
)

var conn *pgx.Conn

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
	conn, err = pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Unable to connection to database: %v\n",
			err,
		)
		os.Exit(1)
	}
}

func MigrateDatabase() {
	_, err := conn.Exec(context.Background(),
		CREATE_USER_TABLE,
	)

	if err != nil {
		log.Fatal(err)
	}
	// Create("user1", "pass")
	// Update("user2", "admin", "user1")
	// Delete("user2")
}

func Create(username string, password string) {
	_, err := conn.Exec(context.Background(),
		INSERT_USER,
		username,
		password,
	)

	if err != nil {
		log.Fatal(err)
	}

}

func Read(username string) string {
	_, err := conn.Exec(context.Background(),
		SELECT_USER,
		username,
	)

	if err != nil {
		log.Fatal(err)
	}
	return ""
}

func Update(
	username string,
	new_username string,
	new_password string,
) {
	_, err := conn.Exec(context.Background(),
		UPDATE_USER,
		username,
		new_username,
		new_password,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func Delete(username string) {
	_, err := conn.Exec(context.Background(),
		DELETE_USER,
		username,
	)

	if err != nil {
		log.Fatal(err)
	}
}
