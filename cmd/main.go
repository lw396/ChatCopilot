package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

var ctx *Context

var app = cli.App{
	Name: "github.com/lw396/WeComCopilot",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "config-dir",
			Aliases: []string{"c"},
			Value:   "config",
			Usage:   "配置文件存放目录",
		},
		&cli.StringFlag{
			Name:  "log-dir",
			Value: "logs",
			Usage: "日志文件存放目录",
		},
		&cli.UintFlag{
			Name:  "pod-id",
			Value: 0,
			Usage: "副本ID",
		},
		&cli.BoolFlag{
			Name:  "print-config",
			Value: false,
			Usage: "输出配置文件信息",
		},
	},
	Commands: []*cli.Command{apiCmd},
}

func main() {
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
