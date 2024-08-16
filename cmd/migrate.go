package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/lw396/ChatCopilot/internal/repository/gorm"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/urfave/cli/v3"
)

var migrateTpl string = `
-- +migrate Up
	-- Do something here

-- +migrate Down
	-- Undo something here
`

var migrateCmd = &cli.Command{
	Name:  "migrate",
	Usage: "同步表结构到数据库",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "dir",
			Aliases: []string{"d"},
			Value:   "migration",
			Usage:   "迁移文件目录",
		},
	},
	Before: func(c context.Context, cmd *cli.Command) (err error) {
		ctx, err = buildContext(cmd, "migrate")
		if err != nil {
			return
		}
		return nil
	},
	Commands: []*cli.Command{migrateCreateCmd, migrateUpCmd, migrateDownCmd},
}

var migrateCreateCmd = &cli.Command{
	Name:      "create",
	UsageText: "create name [commad options]",
	Usage:     "创建数据库迁移文件",
	Action: func(c context.Context, cmd *cli.Command) error {
		if cmd.NArg() == 0 {
			return errors.New("缺少迁移文件名称")
		}
		name := cmd.Args().First()
		dir := cmd.String("dir")

		version := time.Now().Format("20060102150405")
		matches, err := filepath.Glob(filepath.Join(dir, fmt.Sprintf("%s_*.sql", version)))
		if err != nil {
			return err
		}
		if len(matches) > 0 {
			return fmt.Errorf("重复的迁移文件版本: %s", version)
		}

		filename := filepath.Join(dir, fmt.Sprintf("%s_%s.sql", version, name))

		f, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0666)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err := f.WriteString(migrateTpl); err != nil {
			return err
		}

		fmt.Printf("成功创建迁移文件%s.\n", filename)
		return nil
	},
}

var migrateUpCmd = &cli.Command{
	Name:      "up",
	UsageText: "up",
	Usage:     "运行迁移文件",
	Flags: []cli.Flag{
		&cli.UintFlag{
			Name:    "limit",
			Aliases: []string{"l"},
			Value:   0,
			Usage:   "迁移步数 (0表示迁移到最新的文件)",
		},
	},
	Action: func(c context.Context, cmd *cli.Command) error {
		req, err := ctx.buildDB()
		if err != nil {
			return err
		}

		dir := cmd.String("dir")
		step := cmd.Int("limit")

		db := gorm.New(req)
		num, err := db.Migrate(dir, migrate.Up, int(step))
		if err != nil {
			return err
		}

		fmt.Printf("运行 %d 步迁移完成\n", num)
		return nil
	},
}

var migrateDownCmd = &cli.Command{
	Name:      "down",
	UsageText: "down",
	Usage:     "回滚迁移文件",
	Flags: []cli.Flag{
		&cli.UintFlag{
			Name:    "limit",
			Aliases: []string{"l"},
			Value:   0,
			Usage:   "回滚步数 (0表示回滚到初始状态)",
		},
	},
	Action: func(c context.Context, cmd *cli.Command) error {
		req, err := ctx.buildDB()
		if err != nil {
			return err
		}

		dir := cmd.String("dir")
		step := cmd.Int("limit")

		db := gorm.New(req)
		num, err := db.Migrate(dir, migrate.Down, int(step))
		if err != nil {
			return err
		}

		fmt.Printf("运行 %d 步回滚完成\n", num)
		return nil
	},
}
