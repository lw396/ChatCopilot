package cmd

import (
	"context"

	"github.com/urfave/cli/v3"
)

var ctx *Context

var App = cli.Command{
	Name:  "github.com/lw396/WeComCopilot",
	Usage: "微信机器人消息转发服务",
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
	Before: func(c context.Context, cmd *cli.Command) (err error) {
		if err = apiApp(c, cmd); err != nil {
			return
		}
		return
	},
	After: func(c context.Context, cmd *cli.Command) (err error) {
		if err = scheduleApp(c, cmd); err != nil {
			return err
		}
		return nil
	},
}
