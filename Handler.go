package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

type event struct {
	Type string
	Data interface{}
}

type moveData struct {
	Row    int
	Col    int
	Player int
	Turn   int
}

func newGame(w http.ResponseWriter, r *http.Request) {
	numOfPlayer := len(b.subscribers)
	if numOfPlayer > 1 {
		fmt.Fprint(w, 0)
	} else if numOfPlayer == 0 {
		fmt.Fprint(w, `{"gameStatus": "waiting", "playerNum": 1}`)
	} else {
		fmt.Fprint(w, `{"gameStatus": "started", "playerNum": 2}`)
		var eventData = event{
			"game",
			"started",
		}
		publishEvent(eventData)
		setupGame()
	}
}

func move(w http.ResponseWriter, r *http.Request) {
	var move moveData
	vars := mux.Vars(r)
	player, err := strconv.Atoi(vars["player"])
	if err != nil {
		log.Fatal(err)
	}
	if player > 0 && player == theGame.turn {
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

			move.Turn = theGame.turn
			var eventData = event{
				"move",
				move,
			}
			publishEvent(eventData)

			fmt.Fprintf(w, "Player %v: Move Initiated", move.Player)
		}

	} else {
		fmt.Fprintln(w, "Invalid Player or Player out of turn")
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
