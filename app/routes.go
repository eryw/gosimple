package app

import (
	"github.com/gorilla/mux"
)

func Routes(router *mux.Router) {
	router.HandleFunc("/", HomeController).Name("home")
	router.HandleFunc("/about", AboutController).Name("about")
	router.HandleFunc("/{article}", ArticleController).Name("article")
}
