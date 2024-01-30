package main

import (
	"log"

	"gugu/routes"
	"gugu/server"
)

func main() {
	// initialize database
	db, err := server.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// initialize user routes
	routes.SetupUserRoutes(db)

	if err := server.InitServer(); err != nil {
		log.Fatal(err)
	}
}
