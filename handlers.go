package gosimple

import (
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"path/filepath"
)

type ViewData struct {
	Title string
	Body  []byte
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	if t, err := template.ParseFiles(filepath.Join(templatesPath, tmpl+".html")); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		t.Execute(w, data)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

}

func pagesHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	filename := params["article"]

	data := struct {
		Whatever string
	}{
		"my variable"}

	renderTemplate(w, filename, data)
}
