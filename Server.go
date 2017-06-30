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
	log.Fatal(http.ListenAndServe(":"+port, router))
}
