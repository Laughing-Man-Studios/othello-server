package main

import (
	"fmt"
	"time"
)

type event struct {
	Type string
	Data interface{}
}

type broker struct {
	subscribers map[chan event]bool
}

var b = &broker{
	subscribers: make(map[chan event]bool),
}

func subscribe() (bool, chan event) {
	ch := make(chan event)
	if len(b.subscribers) <= 2 {
		b.subscribers[ch] = true
		return true, ch
	}
	return false, ch

}

func unsubscribe(ch chan event) {
	var eventData = event{
		"left",
		"",
	}
	delete(b.subscribers, ch)
	publish(eventData)
}

func publish(evt event) {
	for ch := range b.subscribers {
		ch <- evt
	}
}

func waitForSubscribers(numOfSubscribers int) bool {
	timerChan := time.NewTimer(time.Second * 2).C
	for {
		select {
		case <-timerChan:
			return false
		default:
			if len(b.subscribers) == numOfSubscribers {
				return true
			}
		}
	}
}

func ping() {
	tick := time.Tick(50 * time.Second)
	pingEvent := event{
		"ping",
		"ping",
	}
	fmt.Println("inPing")

	for {
		select {
		case <-tick:
			publish(pingEvent)
			fmt.Println("ping")
		}
	}
}
