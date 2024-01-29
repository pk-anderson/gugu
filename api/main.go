package main

import (
	"log"

	"gugu/server"
)

func main() {
	// Inicializa o banco de dados
	db, err := server.InitializeDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Inicia o servidor chamando a função do pacote "server"
	if err := server.InitServer(); err != nil {
		log.Fatal(err)
	}
}
