package service

import (
	"github.com/lw396/WeComCopilot/internal/repository"
	"github.com/lw396/WeComCopilot/pkg/log"
	"github.com/lw396/WeComCopilot/pkg/redis"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	*options
}

type JWTConfig struct {
	Secret          string
	ExpireSecs      int
	DefaultPassword string
}

type AdminConfig struct {
	Username string
	Password string
}

type TaskConfig struct {
	Interval uint8
	Crontab  string
}

type options struct {
	rep    repository.Repository
	logger log.Logger
	tracer trace.Tracer
	redis  redis.RedisClient
	sqlite repository.SQLiteClient
	jwt    *JWTConfig
	admin  *AdminConfig
	task   *TaskConfig
	path   string
}

type Option func(*options)

func WithRepository(rep repository.Repository) Option {
	return func(o *options) {
		o.rep = rep
	}
}

func WithLogger(logger log.Logger) Option {
	return func(o *options) {
		o.logger = logger
	}
}

func WithTracer(tracer trace.Tracer) Option {
	return func(o *options) {
		o.tracer = tracer
	}
}

func WithSQLite(s repository.SQLiteClient) Option {
	return func(o *options) {
		o.sqlite = s
	}
}

func WithRedis(rc redis.RedisClient) Option {
	return func(o *options) {
		o.redis = rc
	}
}

func WithJWT(jwt *JWTConfig) Option {
	return func(o *options) {
		o.jwt = jwt
	}
}

func WithAdmin(admin *AdminConfig) Option {
	return func(o *options) {
		o.admin = admin
	}
}

func WithTask(task *TaskConfig) Option {
	return func(o *options) {
		o.task = task
	}
}

func WithFilePath(path string) Option {
	return func(o *options) {
		o.path = path
	}
}

func New(opts ...Option) *Service {
	o := &options{}
	for _, opt := range opts {
		opt(o)
	}
	return &Service{
		options: o,
	}
}
