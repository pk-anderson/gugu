package handler

import (
	"encoding/json"
	"fmt"
	"gugu/services/user"
	"net/http"
)

type UserHandler struct {
	UserService *user.UserService
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserResponse struct {
	UUID string `json:"uuid"`
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, fmt.Sprintf("error on decoding request: %s", err), http.StatusBadRequest)
		return
	}

	uuid, err := h.UserService.CreateUser(request.Username, request.Email, request.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("error on creating user: %s", err), http.StatusInternalServerError)
		return
	}

	response := CreateUserResponse{UUID: uuid}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error on encoding response: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	w.Write(jsonResponse)
}
