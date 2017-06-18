package main

import (
	"gopkg.in/urfave/cli.v1"
	"github.com/stackimpact/stackimpact-go"
	"os"
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
		cli.StringFlag{
			Name:        "stack_name",
			Value:       "spartan-io",
			Usage:       "Stack impact agent name.",
			Destination: &config.StackImpactName,
		},
		cli.StringFlag{
			Name:        "stack_key",
			Usage:       "Stack impact agent key.",
			Destination: &config.StackImpactKey,
		},
	}
	app.Action = func(c *cli.Context) {
		if config.StackImpactKey != "" && config.StackImpactName != "" {
			agent := stackimpact.NewAgent()
			agent.Start(stackimpact.Options{
				AgentKey: config.StackImpactKey,
				AppName: config.StackImpactName,
			})
		}

		server(config)
	}

	app.Run(os.Args)
}
