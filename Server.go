package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var runMode = flag.String("runMode", "deb", "Mode to Run in. 'prod' for production and 'deb' for debug")

func main() {
	router := newRouter()
	port, exist := os.LookupEnv("PORT")

	if !(exist) {
		port = "8080"
	}

	flag.Parse()

	if *runMode == "deb" {
		startCMDGame()
	} else {
		go ping()
		log.Fatal(http.ListenAndServe(":"+port, router))
	}

}
