package main

import (
	"github.com/googollee/go-socket.io"
	"github.com/gorilla/context"

	"log"
)

func socketIo(events <-chan Event, auth Auth) *socketio.Server {
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	// History channel.
	history, lastMessages := history()

	server.On("connection", func(so socketio.Socket) {
		defer context.Clear(so.Request())

		// Join public channels.
		so.Join("feed")
		so.Join("post")
		so.Join("general")
		so.Join("chat")
		so.Join("user")

		// Get token & perform user retrieval
		token := so.Request().URL.Query().Get("token")
		user, err := auth.User(token)
		if err != nil {
			log.Printf("[err] Could not authenticate user: %v\n", err)
			return
		}

		so.On("chat send", func(channel, message string) {
			if len(message) == 0 || len(channel) == 0 {
				return
			}

			m := userMessage(user, message)
			packed := list(m)

			log.Printf("Incoming chat: %s - %s\n", channel, message)

			// Broadcast to other clients in the pool.
			so.BroadcastTo("chat", "chat "+channel, packed)

			// Send to the history channel.
			history <- PackedMessage{channel, m}
		})

		so.On("chat update-me", func(channel string) {
			messages := lastMessages(channel)
			if len(messages) > 0 {
				packed := list(messages...)
				so.Emit("chat "+channel, packed)
			}
		})

		so.On("disconnection", func() {
			log.Println("Client disconnected.")
			user = nil
		})
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
	})

	// Events channel consumer.
	// Other incoming messages interfaces will interact through
	// the events channel.
	go func() {
		for event := range events {
			server.BroadcastTo(event.Room, event.Event, event.Message)
		}
	}()

	log.Printf("Started socket.io server.")
	return server
}
