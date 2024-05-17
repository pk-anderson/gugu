package routes

import (
	"database/sql"
	handler "gugu/handlers/accessHandler"

	"github.com/gorilla/mux"
)

func initializeAccessHandler(db *sql.DB) *handler.AccessHandler {
	accessHandler := &handler.AccessHandler{DB: db}

	return accessHandler
}

func SetupAccessRoutes(router *mux.Router, db *sql.DB) {
	accessHandler := initializeAccessHandler(db)

	router.HandleFunc("/access/login", accessHandler.LoginHandler).Methods("POST")
}
