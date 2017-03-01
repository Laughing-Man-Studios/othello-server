package main

import (
	"fmt"
	"net/http"
)

//NewGame is a handler for starting a new Game of Othello
func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Started New Game!")
}
