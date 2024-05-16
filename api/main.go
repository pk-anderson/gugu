package main

import (
	"fmt"
	"log"
	"net/http"

	"gugu/routes"
	"gugu/server"

	"github.com/gorilla/mux"
)

func main() {
	db, err := server.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()
	http.Handle("/", r)

	routes.SetupUserRoutes(r, db)

	fmt.Println("Server running on port: 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
