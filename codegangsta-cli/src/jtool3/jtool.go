package main

import (
	"os"
	"github.com/codegangsta/cli"
	"fmt"
)

func main() {
	app := cli.NewApp()
	app.Name = "jtool"
	app.Usage = "A Linux Utility Command Line Tool!"

	var tasks = []string{"cook", "clean", "laundry", "eat", "sleep", "code"}
	app.EnableBashCompletion = true
	app.Commands = []cli.Command{
		{
			Name:  "complete",
			Aliases: []string{"c"},
			Usage: "complete a task on the list",
			Action: func(c *cli.Context) {
				println("completed task: ", c.Args().First())
			},
			BashComplete: func(c *cli.Context) {
				// This will complete if no args are passed
				if len(c.Args()) > 0 {
					return
				}
				for _, t := range tasks {
					fmt.Println(t)
				}
			},
		},
	}

	app.Run(os.Args)
}