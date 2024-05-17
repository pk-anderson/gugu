package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	userInterface "gugu/interfaces/user"
	"gugu/services/user"
	"io"
	"net/http"
)

type UserHandler struct {
	DB *sql.DB
}

// CREATE USER

type CreateUserRequest struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Bio        string `json:"bio"`
	ProfilePic []byte `json:"profile_pic,omitempty"`
}

type CreateUserResponse struct {
	UserId string `json:"userId"`
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "error on parsing multipart form", http.StatusBadRequest)
		return
	}

	request := CreateUserRequest{
		Username: r.FormValue("username"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Bio:      r.FormValue("bio"),
	}

	file, _, err := r.FormFile("profile_pic")
	if err == nil {
		defer file.Close()
		request.ProfilePic, err = io.ReadAll(file)
		if err != nil {
			http.Error(w, "error on reading profile picture", http.StatusInternalServerError)
			return
		}
	} else if err != http.ErrMissingFile {
		http.Error(w, "error on parsing profile picture", http.StatusBadRequest)
		return
	}

	service := user.NewService(h.DB)

	userId, err := service.CreateUser(
		request.Username,
		request.Email,
		request.Password,
		request.Bio,
		request.ProfilePic,
	)
	if err != nil {
		http.Error(w, fmt.Sprintf("error on creating user: %s", err), http.StatusInternalServerError)
		return
	}

	response := CreateUserResponse{UserId: userId}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error on encoding response: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

// LIST USERS

type ListUsersResponse struct {
	Users []userInterface.User `json:"users"`
}

func (h *UserHandler) ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	service := user.NewService(h.DB)

	users, err := service.ListUsers()
	if err != nil {
		http.Error(w, fmt.Sprintf("error on listing users: %s", err), http.StatusInternalServerError)
		return
	}

	response := ListUsersResponse{
		Users: users,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error on encoding response: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
