package routes

import (
	"database/sql"
	handler "gugu/handlers/userHandler"

	"github.com/gorilla/mux"
)

func initializeUserHandler(db *sql.DB) *handler.UserHandler {
	userHandler := &handler.UserHandler{DB: db}

	return userHandler
}

func SetupUserRoutes(router *mux.Router, db *sql.DB) {
	userHandler := initializeUserHandler(db)

	router.HandleFunc("/user/create", userHandler.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/list", userHandler.ListUsersHandler).Methods("GET")
}
