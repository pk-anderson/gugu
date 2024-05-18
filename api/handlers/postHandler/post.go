package posthandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gugu/services/post"
	"gugu/utils"
	"net/http"
)

type PostHandler struct {
	DB *sql.DB
}

// CREATE POST
type CreatePostRequest struct {
	Content string `json:"content"`
}

type CreatePostResponse struct {
	PostId string `json:"postId"`
}

func (h *PostHandler) CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	userId, ok := r.Context().Value(utils.UserIDKey).(string)
	if !ok {
		http.Error(w, "user ID not found in context", http.StatusInternalServerError)
		return
	}

	var request CreatePostRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "error decoding request", http.StatusBadRequest)
		return
	}

	service := post.NewService(h.DB)
	postId, err := service.CreatePost(userId, request.Content)
	if err != nil {
		http.Error(w, "error creating post", http.StatusInternalServerError)
		return
	}

	response := CreatePostResponse{PostId: postId}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("error on encoding response: %s", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}
