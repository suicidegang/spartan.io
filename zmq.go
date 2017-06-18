package main

import (
	zmq4 "github.com/pebbe/zmq4"
	"log"
	"encoding/json"
)

func zmq(port string) <-chan Event {
	events := make(chan Event)

	// Pull zmq socket server
	receiver, err := zmq4.NewSocket(zmq4.PULL)
	if err != nil {
		panic(err)
	}

	err = receiver.Bind("tcp://*:" + port)
	if err != nil {
		panic(err)
	}

	log.Println("Started zmq pull server at tcp://*:" + port)

	go func() {
		for {
			var event Event

			// Block until new message is received.
			msg, _ := receiver.Recv(0)

			if err := json.Unmarshal([]byte(msg), &event); err != nil {
				continue
			}

			// Once message is unmarshaled send it back to processing channel
			events <- event

			log.Println("Broadcasted message to " + event.Room)
		}

		close(events)
	}()

	return events
}
