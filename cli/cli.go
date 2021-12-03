package cli

import (
	"github.com/urfave/cli"
	"go-distributed-storage/client"
	"go-distributed-storage/logger"
	"go-distributed-storage/server"
	"go-distributed-storage/storage"
	"log"
	"os"
)

func Start() {
	var app = cli.NewApp()
	var port string
	var url string
	app.Name = "Go Distributed Storage CLI"
	app.Usage = "Manage your distributed storage"
	app.Author = "Andreas Lagerström"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "port",
			Value:       "8080",
			Usage:       "Port to use when starting server",
			Destination: &port,
		},
		&cli.StringFlag{
			Name:        "url",
			Value:       "http://localhost:8080",
			Usage:       "Url to use when sending requests to server",
			Destination: &url,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "Start a server",
			Action: func(c *cli.Context) {
				startServer(port)
			},
		},
		{
			Name:  "list",
			Usage: "Send a request to a server to list entries",
			Action: func(c *cli.Context) {
				listEntries(url)
			},
		},
		{
			Name:  "ping",
			Usage: "Ping a server",
			Action: func(c *cli.Context) {
				ping(url)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func ping(url string) {
	client.Ping(url)
}

func listEntries(url string) {
	client.List(url)
}

func startServer(port string) {
	logger.Log("Starting server on port", port)
	s := server.New(storage.New(), port)
	s.Start()
}
