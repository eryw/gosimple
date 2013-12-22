package view

import (
	"github.com/eryw/gosimple/config"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var templatesPath string

func init() {
	templatesPath = strings.Replace(config.TemplatesPath, "\\", "/", -1)
}

func Render(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = getFullTemplatePath(tmpl)
	if _, err := os.Stat(tmpl); err != nil {
		http.Error(w, "404 Page Not Found", http.StatusNotFound)
		return
	}
	if t, err := template.ParseFiles(tmpl); err != nil {
		log.Printf(err.Error())
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
	} else {
		t.Execute(w, data)
	}
}

func getRealFileName(path string) string {
	path = strings.Replace(path, "\\", "/", -1)
	last := strings.LastIndex(path, "/")
	return path[last+1:]
}

func getFullTemplatePath(tmpl string) string {
	tmpl = getRealFileName(tmpl)
	return filepath.FromSlash(filepath.Join(templatesPath, tmpl+".html"))
}
