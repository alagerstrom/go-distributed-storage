package command

import (
	"go-distributed-storage/client"
	"go-distributed-storage/logger"
	"go-distributed-storage/server"
	"go-distributed-storage/storage"
	"log"
	"os"

	"github.com/urfave/cli"
)

func Start() {
	var app = cli.NewApp()
	var url string
	app.Name = "Go Distributed Storage CLI"
	app.Usage = "Manage your distributed storage"
	app.Author = "Andreas Lagerström"
	app.Version = "1.0.0"

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:        "url",
			Value:       "http://localhost:8080",
			Usage:       "Url to use when sending requests to server",
			Destination: &url,
		},
	}
	app.Commands = []cli.Command{
		{
			Name:        "server",
			Description: "Server options",
			Subcommands: cli.Commands{
				{
					Name:        "start",
					Description: "Start a server",
					Usage:       "[port] [name]",
					Action: func(c *cli.Context) error {
						port := c.Args().First()
						name := c.Args().Get(1)
						logger.Log("Starting server on port", port, "with name", name)
						s := server.New(storage.New(), port, name)
						s.Start()
						return nil
					},
				},
			},
		},
		{
			Name:  "client",
			Usage: "Client options",
			Subcommands: cli.Commands{
				{
					Name:  "put",
					Usage: "[key] [value]",
					Action: func(c *cli.Context) error {
						client.Put(url, c.Args().First(), c.Args().Get(1))
						return nil
					},
				},
				{
					Name:  "get",
					Usage: "[key]",
					Action: func(c *cli.Context) error {
						client.Get(url, c.Args().First())
						return nil
					},
				},
				{
					Name:  "delete",
					Usage: "[key]",
					Action: func(c *cli.Context) error {
						client.Delete(url, c.Args().First())
						return nil
					},
				},
				{
					Name:  "list",
					Usage: "Send a request to a server to list entries",
					Action: func(c *cli.Context) {
						client.List(url)
					},
				},
				{
					Name:        "connect",
					Usage:       "[url]",
					Description: "Connect a server to another server",
					Action: func(c *cli.Context) {
						client.Connect(url, c.Args().First())
					},
				},
				{
					Name:        "ping",
					Description: "Ping a server",
					Action: func(c *cli.Context) {
						client.Ping(url)
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
