package helpers
import (
	"net/http"
	"html/template"
)

var templates = template.Must(template.ParseFiles("edit.html", "view.html", "header.tmpl", "editModal.tmpl"))

func RenderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
