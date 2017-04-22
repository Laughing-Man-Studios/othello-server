package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

//NewRouter returns a new mux router with each given route and handler to go with
func NewRouter() *mux.Router {
	router := mux.NewRouter()
	for _, route := range routes {
		var handler http.Handler
		handler = Logger(route.HandlerFunc, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	for _, route := range static {

		router.
			PathPrefix(route.Pattern).
			Handler(http.FileServer(http.Dir(route.Dir)))
	}

	return router
}