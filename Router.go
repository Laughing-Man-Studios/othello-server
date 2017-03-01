package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

const srcPath string = "./src/github.com/Laughing-Man-Studios/othello-server/"

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
			Handler(http.FileServer(http.Dir(srcPath + route.Dir)))
	}

	return router
}
