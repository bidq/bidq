package main

import (
	"log"
	"os"
	"time"

	bidq "github.com/bidq/core"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func setup() {
	var host string
	var port uint
	var timeout uint
	app.Name = "BidQ CLI"
	app.Usage = "A simple CLI running BidQ"
	app.Author = "oakfang (Alon Niv)"
	app.Version = "0.2.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "host, s",
			Value:       "127.0.0.1",
			Usage:       "host",
			Destination: &host,
		},
		cli.UintFlag{
			Name:        "port, p",
			Value:       8888,
			Usage:       "port",
			Destination: &port,
		},
		cli.UintFlag{
			Name:        "timeout, t",
			Value:       5000,
			Usage:       "queue timeout",
			Destination: &timeout,
		},
	}
	app.Action = func(c *cli.Context) {
		bidq.Start(host, uint16(port), time.Duration(timeout)*time.Millisecond)
	}
}

func main() {
	setup()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
