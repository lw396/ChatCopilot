package db

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/lw396/ChatCopilot/pkg/log"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type gormLogger struct {
	l             log.Logger
	slowThreshold time.Duration
	debug         bool
	level         logger.LogLevel
}

func (l *gormLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.level = level
	return l
}

func (l gormLogger) Info(ctx context.Context, format string, v ...interface{}) {
	if l.level >= logger.Info {
		l.l.WithContext(ctx).Debugf(format, v...)
	}
}

func (l gormLogger) Warn(ctx context.Context, format string, v ...interface{}) {
	if l.level >= logger.Warn {
		l.l.WithContext(ctx).Warnf(format, v...)
	}
}

func (l gormLogger) Error(ctx context.Context, format string, v ...interface{}) {
	if l.level >= logger.Error {
		l.l.WithContext(ctx).Errorf(format, v...)
	}
}

func (l gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	var msg, sql string
	var rows int64
	elapsed := time.Since(begin)
	switch {
	case err != nil && !errors.Is(err, gorm.ErrRecordNotFound):
		sql, rows = fc()
		msg = err.Error()
	case l.slowThreshold != 0 && elapsed >= l.slowThreshold:
		sql, rows = fc()
		msg = fmt.Sprintf("slow sql >= %s", l.slowThreshold.String())
	case l.debug:
		sql, rows = fc()
	}

	if sql == "" {
		return
	}

	rs := "-"
	if rows != -1 {
		rs = strconv.FormatInt(rows, 10)
	}
	l.Info(ctx, "%s %s [%s] [rows:%s] %s", utils.FileWithLineNum(), msg, elapsed.String(), rs, sql)
}
