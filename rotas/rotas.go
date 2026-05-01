package rotas

import (
	"modulo/controle"
	"net/http"
)

func CarregaRotas() {
	// Rota principal
	http.HandleFunc("/", controle.Index)

	// Servidor de arquivos estáticos (CSS, JS, Imagens)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}
