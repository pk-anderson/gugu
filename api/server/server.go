package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitServer() error {
	r := mux.NewRouter()

	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	fmt.Println("Server running on port: 8080")
	return nil
}
