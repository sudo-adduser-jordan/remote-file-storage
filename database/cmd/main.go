package main

import (
	"main/database"
	"main/routes"
)

func main() {
	database.ConnectToDatabase()
	database.MigrateDatabase()
	routes.SetupRoutes()
}
