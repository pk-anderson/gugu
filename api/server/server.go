package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func InitServer() error {
	// Cria um roteador Gorilla Mux
	r := mux.NewRouter()

	// Adiciona o roteador como manipulador padrão
	http.Handle("/", r)

	// Inicia o servidor na porta 8080
	fmt.Println("Server running on port: 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return err
	}
	return nil
}
