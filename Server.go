package main

import (
	"log"
	"net/http"
	"os"
    "fmt"
)

func main() {
	router := newRouter()
	port, exist := os.LookupEnv("PORT")
	if !(exist) {
		port = "8080"
	}
    fmt.Printf("%v\n", port)
	go ping()
	log.Fatal(http.ListenAndServe(":"+port, router))
}
