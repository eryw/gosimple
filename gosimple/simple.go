package main

import (
	"fmt"
	"github.com/eryw/gosimple"
	"github.com/gorilla/context"
	"github.com/keep94/weblogs"
	"net/http"
	"os"
)

func main() {
	if err := gosimple.Bootstrap(); err != nil {
		fmt.Println("Error starting server!")
		os.Exit(1)
	}
	options := &weblogs.Options{Logger: weblogs.ApacheCommonLogger()}
	handler := context.ClearHandler(weblogs.HandlerWithOptions(http.DefaultServeMux, options))
	http.ListenAndServe(":8000", handler)
}
