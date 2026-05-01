package controle

import (
	"html/template"
	"modulo/modelo"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	dados := modelo.GetPageData()
	temp.ExecuteTemplate(w, "index.html", dados)
}
