package main

import (
	"log"
	"modulo/rotas"
	"net/http"
	"os"
)

func main() {
	rotas.CarregaRotas()

	// Pega a porta fornecida pelo servidor de hospedagem
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" // Usa a 8000 se estiver rodando local no seu PC
	}

	log.Println("Servidor rodando na porta " + port)
	http.ListenAndServe(":"+port, nil)
}
