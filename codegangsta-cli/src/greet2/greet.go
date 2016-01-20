package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"
	app.Action = func(ctx *cli.Context) {
		args := ctx.Args()

		if len(args) == 0 {
			println("Hello, World!")
		} else {
			for _, arg := range args {
				println("Hello, ", arg, " !")
			}
		}
	}

	app.Run(os.Args)
}