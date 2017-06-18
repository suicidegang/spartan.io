package main

import (
	"gopkg.in/urfave/cli.v1"
	"os"
	"log"
)

func main() {
	var config ServerOptions

	app := cli.NewApp()
	app.Name = "Spartan.IO"
	app.Usage = "Starts the socket.io spartan service."
	app.Version = "0.0.1"
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name:        "port",
			Value:       "3100",
			Usage:       "Socket.io port",
			Destination: &config.Port,
		},
		cli.StringFlag{
			Name:        "zmq",
			Value:       "5557",
			Usage:       "ZMQ pull port",
			Destination: &config.ZmqPort,
		},
		cli.StringFlag{
			Name:        "mgo",
			Value:       "mongodb://127.0.0.1",
			Usage:       "MongoDB address",
			Destination: &config.MgoAddress,
		},
		cli.StringFlag{
			Name:        "db",
			Value:       "spartangeek",
			Usage:       "Mongo Database Name",
			Destination: &config.MgoDB,
		},
		cli.StringFlag{
			Name:        "jwt_secret",
			Usage:       "JWT secret for sessions.",
			Destination: &config.JwtSecret,
		},
	}
	app.Action = func(c *cli.Context) {
		log.Printf("%+v\n", config)
		server(config)
	}

	app.Run(os.Args)
}
