package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
	"time"
)

type ServerOptions struct {
	Port       string
	ZmqPort    string
	MgoAddress string
	MgoDB string
	JwtSecret  string
	StackImpactName string
	StackImpactKey string
}

func server(config ServerOptions) {
	db, err := mongo(config.MgoAddress, config.MgoDB)
	if err != nil {
		log.Fatal(err)
	}

	// ZMQ server computes an events channel.
	auth := Auth{db, config.JwtSecret}
	events := zmq(config.ZmqPort)

	mux := http.NewServeMux()
	mux.Handle("/socket.io/", socketIo(events, auth))
	handler := cors.New(cors.Options{
		AllowCredentials: true,
	}).Handler(mux)

	s := &http.Server{
		Addr:           ":" + config.Port,
		Handler:        handler,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
