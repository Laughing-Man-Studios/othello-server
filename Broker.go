package main

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

func subscribe() chan event {
	ch := make(chan event)
	b.subscribers[ch] = true
	return ch
}

func unsubscribe(ch chan event) {
	delete(b.subscribers, ch)
}

func publish(evt event) {
	for ch := range b.subscribers {
		ch <- evt
	}
}
