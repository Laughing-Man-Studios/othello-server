package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

//MoveData tells where a peice moved from and to, and the player who made the move
type MoveData struct {
	Row    int
	Col    int
	Player int
}

//NewGame is a handler for starting a new Game of Othello
func NewGame(w http.ResponseWriter, r *http.Request) {
	numOfPlayer := len(b.subscribers)
	if numOfPlayer > 1 {
		fmt.Fprint(w, 0)
	} else {
		fmt.Fprint(w, numOfPlayer+1)
	}
	SetupGame()
}

//Move is a handler for logging a move made by a player
func Move(w http.ResponseWriter, r *http.Request) {
	var move MoveData
	vars := mux.Vars(r)
	player, err := strconv.Atoi(vars["player"])
	if err != nil {
		log.Fatal(err)
	}
	if player > 0 {
		move.Player = player
		err = r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}

		decoder := schema.NewDecoder()
		decoder.Decode(&move, r.PostForm)

		moveJSON, err := json.Marshal(move)
		if err != nil {
			log.Fatal(err)
		}

		jsonStr := string(moveJSON)
		Publish(jsonStr)

		fmt.Fprintf(w, "Player %v: Move Initiated", move.Player)
	}
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

	cn := w.(http.CloseNotifier)

	for {
		select {
		case m := <-ch:
			msg := fmt.Sprintf("data: %s\n\n", m)
			fmt.Fprintln(w, msg)
			f.Flush()
		case <-cn.CloseNotify():
			fmt.Println("Connection Close")
			return
		}
	}
}
