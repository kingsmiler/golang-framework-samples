package main

import (
	"os"
	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	// 应用名
	app.Name = "boom"
	// 使用说明
	app.Usage = "make an explosive entrance"
	// 默认执行事件
	app.Action = func(c *cli.Context) {
		println("boom! I say!")
	}

	// 启动
	app.Run(os.Args)
}