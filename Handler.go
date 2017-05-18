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

type moveData struct {
	Row    int
	Col    int
	Player int
}

func newGame(w http.ResponseWriter, r *http.Request) {
	numOfPlayer := len(b.subscribers)
	if numOfPlayer > 1 {
		fmt.Fprint(w, 0)
	} else {
		fmt.Fprint(w, numOfPlayer+1)
	}
	setupGame()
}

func move(w http.ResponseWriter, r *http.Request) {
	var move moveData
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

		if !movePiece(move) {
			fmt.Fprintf(w, "Player %v: Invalid Move", move.Player)
		} else {
			moveJSON, err := json.Marshal(move)
			if err != nil {
				log.Fatal(err)
			}

			jsonStr := string(moveJSON)
			Publish(jsonStr)

			fmt.Fprintf(w, "Player %v: Move Initiated", move.Player)
		}

	}
}

func events(w http.ResponseWriter, r *http.Request) {
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
