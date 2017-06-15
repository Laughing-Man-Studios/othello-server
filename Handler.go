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

type moveResponse struct {
	Valid bool
}

type newGameResponse struct {
	Full   bool
	Player int
}

func newGame(w http.ResponseWriter, r *http.Request) {
	numOfPlayer := len(b.subscribers)
	var response = newGameResponse{
		true,
		0,
	}
	if numOfPlayer == 0 {
		response.Full = false
		response.Player = 1
	} else if numOfPlayer == 1 {
		response.Full = false
		response.Player = 2
		// TODO: This event needs to happen when both players have connected to the eventHandler.
		// Need to figure out how to ensure that both are connected (maybe move logic to Broker?)
		var eventData = event{
			"start",
			startData{
				1,
			},
		}
		defer publish(eventData)
		setupGame()
	}
	printResponse(w, response)
}

func move(w http.ResponseWriter, r *http.Request) {
	var move moveData
	var response = moveResponse{
		true,
	}
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
			response.Valid = false
		} else {
			move.Turn = theGame.turn
			hasMove, board := findPotentialMoves(theGame.board, move.Turn)
			if hasMove {
				move.Board = board
			}
			var eventData = event{
				"move",
				move,
			}
			defer publish(eventData)
		}

	} else {
		response.Valid = false
	}
	printResponse(w, response)
}

func events(w http.ResponseWriter, r *http.Request) {
	f := w.(http.Flusher)
	ch := subscribe()
	defer unsubscribe(ch)

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	f.Flush()

	cn := w.(http.CloseNotifier)

	for {
		select {
		case m := <-ch:
			data, err := json.Marshal(m.Data)
			if err != nil {
				log.Fatal(err)
			}

			evt := fmt.Sprintf("event: %s\n", m.Type)
			msg := fmt.Sprintf("data: %s\n\n", data)

			fmt.Fprint(w, evt)
			fmt.Fprintln(w, msg)
			f.Flush()
		case <-cn.CloseNotify():
			fmt.Println("Connection Close")
			return
		}
	}
}

func printResponse(w http.ResponseWriter, response interface{}) {
	eventJSON, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprint(w, string(eventJSON))
}
