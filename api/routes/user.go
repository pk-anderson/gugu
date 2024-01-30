package routes

import (
	"database/sql"
	handler "gugu/handlers/userHandler"
	repository "gugu/repositories/userRepository"
	service "gugu/services/user"
	"net/http"

	"github.com/gorilla/mux"
)

func initializeUserHandler(db *sql.DB) *handler.UserHandler {
	userRepo := &repository.UserRepository{DB: db}
	userService := &service.UserService{UserRepository: userRepo}
	userHandler := &handler.UserHandler{UserService: userService}

	return userHandler
}

// setupRoutes é uma função auxiliar para configurar as rotas
func SetupUserRoutes(db *sql.DB) {
	userHandler := initializeUserHandler(db)

	router := mux.NewRouter()

	// create user
	router.HandleFunc("/users", userHandler.CreateUserHandler).Methods("POST")

	http.Handle("/", router)
}
