package handler

import (
	"html/template"
	"modulo/modelo"
	"net/http"
	"path/filepath"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	// Caminho para encontrar o seu HTML na Vercel
	caminhoTemplate := filepath.Join("templates", "index.html")

	temp, err := template.ParseFiles(caminhoTemplate)
	if err != nil {
		http.Error(w, "Erro ao carregar o site: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Puxa os dados da academia lá da sua pasta modelo
	dados := modelo.GetPageData()

	// Mostra o HTML na tela
	temp.Execute(w, dados)
}
