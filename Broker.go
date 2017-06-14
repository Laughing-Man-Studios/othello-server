package main

import (
	"encoding/json"
	"log"
)

type event struct {
	Type string
	Data interface{}
}

//SterilizedEvent is a struct used to pass through the subscription channel in order
//to send to an event down the event stream
type SterilizedEvent struct {
	Type string
	Data string
}

type broker struct {
	subscribers map[chan SterilizedEvent]bool
}

var b = &broker{
	subscribers: make(map[chan SterilizedEvent]bool),
}

//Subscribe is a method that allows you to get a channel from the broker
func Subscribe() chan SterilizedEvent {
	ch := make(chan SterilizedEvent)
	b.subscribers[ch] = true
	return ch
}

//Unsubscribe is a method that removes your channel from the broker
func Unsubscribe(ch chan SterilizedEvent) {
	delete(b.subscribers, ch)
}

//Publish is a method that publishes a method on all channels
func Publish(evt SterilizedEvent) {
	for ch := range b.subscribers {
		ch <- evt
	}
}

func publishEvent(evt event) {
	var strEvt SterilizedEvent
	eventJSON, err := json.Marshal(evt.Data)
	if err != nil {
		log.Fatal(err)
	}
	strEvt.Type = evt.Type
	strEvt.Data = string(eventJSON)
	Publish(strEvt)
}
