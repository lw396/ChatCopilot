package main

import (
	"context"
	"log"
	"os"

	"github.com/urfave/cli/v3"
)

var ctx *Context

var app = cli.Command{
	Name:  "github.com/lw396/WeComCopilot",
	Usage: "微信消息转存储",
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
	Commands: []*cli.Command{apiCmd, scheduleCmd},
}

func main() {
	if err := app.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
