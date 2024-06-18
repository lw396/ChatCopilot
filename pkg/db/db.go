package db

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/lw396/WeComCopilot/pkg/log"

	"github.com/lw396/WeComCopilot/pkg/cache"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type options struct {
	driver        string
	dsn           string
	slowThreshold time.Duration
	isDebug       bool
	logger        log.Logger
	cacheStore    cache.CacheStore
	idGenerator   IDGenerator
	tracer        trace.Tracer
}

func defaultOptions() *options {
	return &options{
		driver:        "mysql",
		dsn:           "root:secret@tcp(127.0.0.1:3306)/WeChatCopilot?charset=utf8&parseTime=true&loc=UTC",
		slowThreshold: time.Millisecond * 100,
		logger:        log.NewConsoleLogger("DB"),
		cacheStore:    cache.DefaultStore(),
		tracer:        otel.Tracer("github.com/lw396/ChatCopilot/pkg/db"),
	}
}

type Option func(*options)

func WithDriver(driver string) Option {
	return func(o *options) {
		o.driver = driver
	}
}

func WithDSN(dsn string) Option {
	return func(o *options) {
		o.dsn = dsn
	}
}

func WithLogger(l log.Logger) Option {
	return func(o *options) {
		o.logger = l
	}
}

func WithSlowThreshold(slowThreshold time.Duration) Option {
	return func(o *options) {
		o.slowThreshold = slowThreshold
	}
}

func WithDebug(isDebug bool) Option {
	return func(o *options) {
		o.isDebug = isDebug
	}
}

func WithCacheStore(store cache.CacheStore) Option {
	return func(o *options) {
		o.cacheStore = store
	}
}

func WithIDGenerator(idGenerator IDGenerator) Option {
	return func(o *options) {
		o.idGenerator = idGenerator
	}
}

func WithTracer(tracer trace.Tracer) Option {
	return func(o *options) {
		o.tracer = tracer
	}
}

type IDGenerator interface {
	ID() int64
}

func New(opts ...Option) (*gorm.DB, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	logger := &gormLogger{
		l:             o.logger,
		slowThreshold: o.slowThreshold,
		debug:         o.isDebug,
	}

	var (
		db  *gorm.DB
		err error
	)

	switch o.driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(o.dsn), &gorm.Config{
			Logger: logger,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
			},
		})
	default:
		return nil, errors.New("unsupported database driver")
	}

	if err != nil {
		return nil, err
	}

	if o.cacheStore != nil {
		cachePlugin := NewCachePlugin(o.cacheStore)
		_ = db.Use(cachePlugin)
	}

	if o.tracer != nil {
		tracePlugin := NewTracePlugin(o.tracer)
		_ = db.Use(tracePlugin)
	}

	if o.idGenerator != nil {
		if err := db.Callback().Create().Before("gorm:before_create").Register("set_id", func(db *gorm.DB) {
			setID(db.Statement.Context, db, o.idGenerator)
		}); err != nil {
			return nil, err
		}
	}

	return db, nil
}

func setID(ctx context.Context, db *gorm.DB, idGenerator IDGenerator) {
	switch db.Statement.ReflectValue.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < db.Statement.ReflectValue.Len(); i++ {
			rv := db.Statement.ReflectValue.Index(i)
			if reflect.Indirect(rv).Kind() != reflect.Struct {
				break
			}

			field := db.Statement.Schema.PrioritizedPrimaryField
			if field != nil {
				if _, isZero := field.ValueOf(ctx, rv); isZero {
					if err := field.Set(ctx, rv, idGenerator.ID()); err != nil {
						_ = db.AddError(err)
					}
				}
			}
		}
	case reflect.Struct:
		field := db.Statement.Schema.PrioritizedPrimaryField
		if field != nil {
			if _, isZero := field.ValueOf(ctx, db.Statement.ReflectValue); isZero {
				if err := field.Set(ctx, db.Statement.ReflectValue, idGenerator.ID()); err != nil {
					_ = db.AddError(err)
				}
			}
		}
	}
}
