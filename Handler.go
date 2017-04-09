package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Square contains the row and col of a designated square on the board
type Square struct {
	Row int
	Col int
}

//MoveData tells where a peice moved from and to, and the player who made the move
type MoveData struct {
	Old    Square
	New    Square
	Player int
}

//NewGame is a handler for starting a new Game of Othello
func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Started New Game!")
}

//Move is a handler for logging a move made by a player
func Move(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	for decoder.More() {
		var u MoveData
		err := decoder.Decode(&u)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Fprintln(w, "Move Initiated")
}

//Events is the handler for connecting with the SSE broker
func Events(w http.ResponseWriter, r *http.Request) {
	f := w.(http.Flusher)
	ch := Subscribe()
	defer Unsubscribe(ch)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	fmt.Fprintln(w, "Connected To EventHandler!")
	f.Flush()

	for {
		m := <-ch
		msg := fmt.Sprintf("data: %s\n\n", m)
		fmt.Fprintln(w, msg)
		f.Flush()
	}
}
