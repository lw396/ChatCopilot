package log

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	Info(v ...interface{})
	Infof(format string, v ...interface{})
	Infow(msg string, keysAndValues ...interface{})
	Debug(v ...interface{})
	Debugf(format string, v ...interface{})
	Debugw(msg string, keysAndValues ...interface{})
	Warn(v ...interface{})
	Warnf(format string, v ...interface{})
	Warnw(msg string, keysAndValues ...interface{})
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorw(msg string, keysAndValues ...interface{})
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicw(msg string, keysAndValues ...interface{})
	DPanic(v ...interface{})
	DPanicf(format string, v ...interface{})
	DPanicw(msg string, keysAndValues ...interface{})
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalw(msg string, keysAndValues ...interface{})
	WithContext(ctx context.Context) Logger
	SetScope(scope string) Logger
}

type DefaultLogger struct {
	*zap.SugaredLogger
}

func (l DefaultLogger) WithContext(ctx context.Context) Logger {
	span := trace.SpanFromContext(ctx)
	traceID := span.SpanContext().TraceID()
	if traceID.IsValid() {
		return DefaultLogger{
			SugaredLogger: l.SugaredLogger.With(zap.String("trace", traceID.String())),
		}
	}
	return l
}

func (l DefaultLogger) SetScope(scope string) Logger {
	return DefaultLogger{
		l.SugaredLogger.Named(scope),
	}
}

type Config struct {
	App    string
	Scope  string
	LogDir string
	Debug  bool
	MaxAge int
}

func NewConsoleLogger(scope string, opts ...zap.Option) DefaultLogger {
	debugPriority := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= zapcore.DebugLevel
	})
	consoleDebug := zapcore.Lock(os.Stdout)
	cfg := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	consoleEncoder := zapcore.NewConsoleEncoder(cfg)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleDebug, debugPriority),
	)

	options := []zap.Option{
		zap.AddCaller(), zap.AddStacktrace(zap.DPanicLevel),
	}

	options = append(options, opts...)

	logger := zap.New(core, options...)

	sugar := logger.Sugar().Named(scope)

	return DefaultLogger{
		sugar,
	}
}

func NewLogger(c Config, opts ...zap.Option) DefaultLogger {
	errPriority := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= zapcore.ErrorLevel
	})
	infoPriority := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= zapcore.InfoLevel
	})
	debugPriority := zap.LevelEnablerFunc(func(lv zapcore.Level) bool {
		return lv >= zapcore.DebugLevel
	})

	infoWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(c.LogDir, fmt.Sprintf("%s.log", c.App)),
		MaxSize:    128,
		MaxBackups: 3,
		MaxAge:     c.MaxAge,
		Compress:   true,
	})

	errWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(c.LogDir, fmt.Sprintf("%s-err.log", c.App)),
		MaxSize:    128,
		MaxBackups: 3,
		MaxAge:     c.MaxAge,
		Compress:   true,
	})
	consoleDebug := zapcore.Lock(os.Stdout)

	cfg := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}
	var options []zap.Option
	options = append(options, zap.AddCaller(), zap.AddStacktrace(zap.DPanicLevel))
	if c.Debug {
		options = append(options, zap.Development())
	}
	options = append(options, opts...)

	consoleEncoder := zapcore.NewConsoleEncoder(cfg)
	jsonEncoder := zapcore.NewJSONEncoder(cfg)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleDebug, debugPriority),
		zapcore.NewCore(jsonEncoder, infoWriter, infoPriority),
		zapcore.NewCore(jsonEncoder, errWriter, errPriority),
	)

	logger := zap.New(core, options...)

	name := c.Scope
	if name == "" {
		name = c.App
	}
	sugar := logger.Sugar().Named(name)

	return DefaultLogger{
		sugar,
	}
}
