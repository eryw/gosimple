package gosimple

import (
	"github.com/gorilla/mux"
	"net/http"
)

func Dispatch() {
	router := mux.NewRouter()
	router.HandleFunc("/", homeHandler).Name("home")
	router.HandleFunc("/about", aboutHandler).Name("about")
	router.HandleFunc("/{article}", pagesHandler).Name("article")

	http.Handle("/", router)
}
