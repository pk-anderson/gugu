package routes

import (
	"database/sql"
	handler "gugu/handlers/userHandler"
	repository "gugu/repositories/userRepository"
	service "gugu/services/user"

	"github.com/gorilla/mux"
)

func initializeUserHandler(db *sql.DB) *handler.UserHandler {
	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{UserRepository: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	return userHandler
}

func SetupUserRoutes(router *mux.Router, db *sql.DB) {
	userHandler := initializeUserHandler(db)

	router.HandleFunc("/user/create", userHandler.CreateUserHandler).Methods("POST")
	router.HandleFunc("/user/list", userHandler.ListUsersHandler).Methods("GET")
}
