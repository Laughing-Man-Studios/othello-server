package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	router := newRouter()
	port, exist := os.LookupEnv("PORT")
	if !(exist) {
		port = "8080"
	}
	go ping()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
