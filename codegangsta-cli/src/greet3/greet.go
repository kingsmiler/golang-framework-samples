package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "greet"
	app.Usage = "fight the loneliness!"

	var vip bool
	var name string
	app.Flags = [] cli.Flag{
		// 字符串
		cli.StringFlag{
			Name:"name",
			Value: "world",
			Usage: "your name",
			Destination: &name,
		},
		// 布尔值
		cli.BoolFlag{
			Name: "vip",
			Usage: "whether it's vip or not",
			Destination: &vip,
		},
	}
	app.Action = func(ctx *cli.Context) {
		n := ctx.String("name")
		v := ctx.Bool("vip")

		println("Hello, ", n, "(", name, ")!")
		println("Are you vip? ", v, "(", vip, ")!")
	}

	app.Run(os.Args)
}