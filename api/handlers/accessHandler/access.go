package accesshandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gugu/services/access"
	"gugu/utils"
	"net/http"
)

type AccessHandler struct {
	DB *sql.DB
}

// LOGIN
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func (h *AccessHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	service := access.NewService(h.DB)
	token, err := service.Login(req.Email, req.Password)
	if err != nil {
		if err.Error() == "invalid credentials" {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := LoginResponse{Token: token}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error on encoding response: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

// LOGOUT

func (h *AccessHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	token, ok := r.Context().Value(utils.TokenString).(string)
	if !ok {
		http.Error(w, "Token not found in context", http.StatusInternalServerError)
		return
	}

	service := access.NewService(h.DB)
	err := service.Logout(token)
	if err != nil {
		http.Error(w, "Error revoking token", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
}
