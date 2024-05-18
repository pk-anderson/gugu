package routes

import (
	"database/sql"
	handler "gugu/handlers/postHandler"
	"gugu/middlewares"

	"github.com/gorilla/mux"
)

func initializePostHandler(db *sql.DB) *handler.PostHandler {
	postHandler := &handler.PostHandler{DB: db}

	return postHandler
}

func SetupPostRoutes(router *mux.Router, db *sql.DB) {
	postHandler := initializePostHandler(db)

	router.HandleFunc("/post/create", middlewares.AuthMiddleware(db, postHandler.CreatePostHandler)).Methods("POST")
}
