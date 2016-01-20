package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "jtool"
	app.Usage = "A linux utility command line tool!"

	app.Commands = []cli.Command{
		{
			Name:      "as",
			Aliases:     []string{"as"},
			Usage:     "application server operation",
			Subcommands: []cli.Command{
				{
					Name:  "start",
					Usage: "Start the application server",
					Action: func(c *cli.Context) {
						println("Starting...")
					},
				},
				{
					Name:  "stop",
					Usage: "Stop the application server",
					Action: func(c *cli.Context) {
						println("Stopping...")
					},
				},
			},
		},
	}

	app.Run(os.Args)
}