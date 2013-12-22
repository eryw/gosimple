package gosimple

import (
	"github.com/eryw/gosimple/app"
	"github.com/gorilla/mux"
	"net/http"
)

func Bootstrap() (err error) {
	dispatch()
	return nil
}

func dispatch() {
	router := mux.NewRouter()
	app.Routes(router)
	http.Handle("/", router)
}
