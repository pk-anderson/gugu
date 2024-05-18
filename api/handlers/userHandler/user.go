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

func extractCreateUserRequest(r *http.Request) (CreateUserRequest, error) {
	var request CreateUserRequest
	request.Username = r.FormValue("username")
	request.Email = r.FormValue("email")
	request.Password = r.FormValue("password")
	request.Bio = r.FormValue("bio")

	file, _, err := r.FormFile("profile_pic")
	if err == nil {
		defer file.Close()
		request.ProfilePic, err = io.ReadAll(file)
		if err != nil {
			return CreateUserRequest{}, fmt.Errorf("error on reading profile picture: %s", err)
		}
	} else if err != http.ErrMissingFile {
		return CreateUserRequest{}, fmt.Errorf("error on parsing profile picture: %s", err)
	}

	return request, nil
}

func (h *UserHandler) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "error on parsing multipart form", http.StatusBadRequest)
		return
	}

	request, err := extractCreateUserRequest(r)
	if err != nil {
		http.Error(w, fmt.Sprintf("error on extracting request: %s", err), http.StatusBadRequest)
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

	response := ListUsersResponse{Users: users}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error on encoding response: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
