package main

import (
	"encoding/json"
	"log"
)

type event struct {
	Type string
	Data interface{}
}

type broker struct {
	subscribers map[chan string]bool
}

var b = &broker{
	subscribers: make(map[chan string]bool),
}

//Subscribe is a method that allows you to get a channel from the broker
func Subscribe() chan string {
	ch := make(chan string)
	b.subscribers[ch] = true
	return ch
}

//Unsubscribe is a method that removes your channel from the broker
func Unsubscribe(ch chan string) {
	delete(b.subscribers, ch)
}

//Publish is a method that publishes a method on all channels
func Publish(msg string) {
	for ch := range b.subscribers {
		ch <- msg
	}
}

func publishEvent(evt event) {
	eventJSON, err := json.Marshal(evt)
	if err != nil {
		log.Fatal(err)
	}

	jsonStr := string(eventJSON)
	Publish(jsonStr)
}
