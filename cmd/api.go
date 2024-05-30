package main

import (
	"context"

	"github.com/lw396/WeComCopilot/api"
	"github.com/lw396/WeComCopilot/internal/repository/gorm"
	"github.com/lw396/WeComCopilot/service"
	"github.com/urfave/cli/v3"
)

var apiCmd = &cli.Command{
	Name:  "api",
	Usage: "启动API服务",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   6978,
			Usage:   "端口号",
		},
	},
	Before: func(c context.Context, cmd *cli.Command) (err error) {
		ctx, err = buildContext(cmd, "app")
		if err != nil {
			return err
		}
		return nil
	},
	Action: func(c context.Context, cmd *cli.Command) error {
		db, err := ctx.buildDB()
		if err != nil {
			return err
		}

		redis, err := ctx.buildRedis()
		if err != nil {
			return err
		}

		service := service.New(
			service.WithRepository(gorm.New(db)),
			service.WithLogger(ctx.buildLogger("API")),
			service.WithSQLite(ctx.buildSQLite()),
			service.WithRedis(redis),
			service.WithJWT(ctx.buildJWT()),
			service.WithAdmin(ctx.buildAdmin()),
			service.WithFilePath(ctx.buildFilePath()),
		)

		port := cmd.Int("port")
		api := api.New(api.Config{
			App:  service,
			Port: port,
		})

		return api.Run()
	},
}
