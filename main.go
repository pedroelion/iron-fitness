package main

import (
	"log"
	"modulo/rotas"
	"net/http"
)

func main() {
	// Carrega as rotas
	rotas.CarregaRotas()

	log.Println("Servidor rodando na porta 8000. Acesse: http://localhost:8000")
	http.ListenAndServe(":8000", nil)
}
