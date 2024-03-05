package main

import (
	"os"

	"github.com/lw396/WeComCopilot/api"
	"github.com/lw396/WeComCopilot/internal/repository/gorm"
	"github.com/lw396/WeComCopilot/pkg/valuer"
	"github.com/lw396/WeComCopilot/service"

	"github.com/urfave/cli/v2"
)

var apiCmd = &cli.Command{
	Name:  "api",
	Usage: "启动API服务",
	Flags: []cli.Flag{
		&cli.UintFlag{
			Name:    "port",
			Aliases: []string{"p"},
			Value:   6978,
			Usage:   "端口号",
		},
	},
	Before: func(c *cli.Context) (err error) {
		ctx, err = buildContext(c, "app")
		if err != nil {
			return err
		}
		return nil
	},
	Action: func(c *cli.Context) error {

		db, err := ctx.buildDB()
		if err != nil {
			return err
		}

		redis, err := ctx.buildRedis()
		if err != nil {
			return err
		}

		tokenKey := valuer.Value("key").Try(
			os.Getenv("TOKEN_KEY"),
			ctx.Section("token").Key("key").String(),
		).String()
		tokenExpire := valuer.Value(3600).Try(
			ctx.Section("token").Key("expire").Int(),
		).Int()

		service := service.New(
			service.WithRepository(gorm.New(db)),
			service.WithRedis(redis),
			service.WithLogger(ctx.buildLogger("API")),
			service.WithJWT(&service.TokenConfig{
				Secret:     tokenKey,
				ExpireSecs: tokenExpire,
			}),
			service.WithSQLite(ctx.buildSQLite()),
		)

		port := c.Int("port")
		api := api.New(api.Config{
			App:  service,
			Port: port,
		})

		return api.Run()
	},
}
