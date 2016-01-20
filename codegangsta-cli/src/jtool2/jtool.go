package main

import (
	"os"
	"github.com/codegangsta/cli"
	"strconv"
)

func main() {
	app := cli.NewApp()
	app.Name = "jtool"
	app.Usage = "A Linux Utility Command Line Tool!"

	app.Commands = []cli.Command{
		{
			Name:      "as",
			Aliases:     []string{"as"},
			Usage:     "application server operation",
			Flags: [] cli.Flag{
				cli.BoolFlag{
					Name: "start, 1",
					Usage: "Start the application server",
				},
				cli.BoolFlag{
					Name: "stop, 0",
					Usage: "Stop the application server",
				},
				cli.BoolFlag{
					Name: "restart, 2",
					Usage: "Restart the application server",
				},
			},
			Action: func(ctx *cli.Context) {
				for _, v := range ctx.Args() {
					println(v)
				}

				start := ctx.Bool("start")
				stop := ctx.Bool("stop")

				println("start=" + strconv.FormatBool(start))
				println("stop=" + strconv.FormatBool(stop))
			},
		},
	}

	app.Run(os.Args)
}