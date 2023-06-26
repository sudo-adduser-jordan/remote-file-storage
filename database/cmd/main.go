package main

import "main/database"

func main() {
	database.ConnectToDatabase()
	database.MigrateDatabase()
	// SetupRoutes()

}
