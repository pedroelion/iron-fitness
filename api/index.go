package handler

import (
	_ "embed"
	"html/template"
	"modulo/modelo"
	"net/http"
)

//go:embed index.html
var htmlContent string

func Handler(w http.ResponseWriter, r *http.Request) {
	// O Go lê o HTML diretamente da variável embutida
	temp, err := template.New("index").Parse(htmlContent)
	if err != nil {
		http.Error(w, "Erro ao carregar o site: "+err.Error(), http.StatusInternalServerError)
		return
	}

	dados := modelo.GetPageData()
	temp.Execute(w, dados)
}
