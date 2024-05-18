package routes

import (
	"database/sql"
	handler "gugu/handlers/accessHandler"
	"gugu/middlewares"

	"github.com/gorilla/mux"
)

func initializeAccessHandler(db *sql.DB) *handler.AccessHandler {
	accessHandler := &handler.AccessHandler{DB: db}

	return accessHandler
}

func SetupAccessRoutes(router *mux.Router, db *sql.DB) {
	accessHandler := initializeAccessHandler(db)

	router.HandleFunc("/access/login", accessHandler.LoginHandler).Methods("POST")
	router.HandleFunc("/access/logout", middlewares.AuthMiddleware(db, accessHandler.LogoutHandler)).Methods("DELETE")
}
