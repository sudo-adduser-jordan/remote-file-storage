package cmd

import "main/database"

func main() {
	database.ConnectToDatabase()
	// MigrateDatabase()
	// SetupRoutes()
}
