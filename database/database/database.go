package database

import (
	"context"
	"fmt"
	"log"
	"main/configs"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/sudo-adduser-jordan/Toolchain/Go/styles"
)

var connection *pgx.Conn

type User struct {
	// int32
	UserID   int
	Username string
	Password string
}

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
	connection, err = pgx.Connect(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Unable to connection to database: %v\n",
			err,
		)
		os.Exit(1)
	}
	fmt.Print("	-----> ")
	fmt.Println(styles.GreenText("Connected to database."))
}

func MigrateDatabase() {
	_, err := connection.Exec(context.Background(),
		CREATE_USER_TABLE,
	)

	if err != nil {
		log.Fatal(err)
	}

	// Create("user1", "pass")
	// Read("user1")
	// Update("user2", "admin", "user1")
	// Read("user2")
	// Delete("user2")

	fmt.Print("	-----> ")
	fmt.Println(styles.BlueText("Database migrated."))
}

func Create(username string, password string) {
	_, err := connection.Exec(context.Background(),
		INSERT_USER,
		username,
		password,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func Read(username string) {
	row, err := connection.Query(context.Background(),
		SELECT_USER,
		username,
	)
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		user := &User{}
		err = row.Scan(&user)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", user.Username)
	}
}

func Update(
	username string,
	new_username string,
	new_password string,
) {
	_, err := connection.Exec(context.Background(),
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
	_, err := connection.Exec(context.Background(),
		DELETE_USER,
		username,
	)

	if err != nil {
		log.Fatal(err)
	}
}

// ===================================== //

func CreateUser(username string, password string) bool {
	_, err := connection.Exec(context.Background(),
		INSERT_USER,
		username,
		password,
	)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func ReadUser(username string) bool {
	row, err := connection.Query(context.Background(),
		SELECT_USER,
		username,
	)
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		user := &User{}
		err = row.Scan(&user)
		if err != nil {
			log.Fatal(err)
		}
		if user.Username != "" {
			return true
		}
	}
	return false
}

func ReadPassword(username string) string {
	row, err := connection.Query(context.Background(),
		SELECT_USER,
		username,
	)

	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		user := &User{}
		err = row.Scan(&user)
		if err != nil {
			log.Fatal(err)
		}
		return user.Password
	}
	return "Read password error."
}
