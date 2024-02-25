package service

import (
	"github.com/lw396/WeComCopilot/internal/repository"
	"github.com/lw396/WeComCopilot/pkg/log"
	"github.com/lw396/WeComCopilot/pkg/redis"
	"github.com/lw396/WeComCopilot/pkg/wechat"

	"go.opentelemetry.io/otel/trace"
)

type Service struct {
	*options
}

type TokenConfig struct {
	Secret     string
	ExpireSecs int
}

type options struct {
	rep    repository.Repository
	logger log.Logger
	tracer trace.Tracer
	token  *TokenConfig
	redis  redis.RedisClient
	wechat wechat.WeChatClient
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

func WithJWT(jc *TokenConfig) Option {
	return func(o *options) {
		o.token = jc
	}
}

func WithRedis(rc redis.RedisClient) Option {
	return func(o *options) {
		o.redis = rc
	}
}

func WithWeChat(wc wechat.WeChatClient) Option {
	return func(o *options) {
		o.wechat = wc
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
