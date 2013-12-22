package app

import (
	"github.com/eryw/gosimple/view"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type aboutData struct {
	Whatever  string
	Something string
	Nothing   string
}

func AboutController(w http.ResponseWriter, r *http.Request) {
	log.Printf("oh no B")
}

type homeData struct {
	Whatever  string
	Something string
	Nothing   string
}

func HomeController(w http.ResponseWriter, r *http.Request) {
	log.Printf("oh no C")
}

type articleData struct {
	Whatever  string
	Something string
	Nothing   string
}

func ArticleController(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	data := &articleData{Whatever: "my variable"}
	view.Render(w, params["article"], data)
}
