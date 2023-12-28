package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	// String
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, target interface{}) (found bool, err error)
	Del(ctx context.Context, key string) error

	// Set
	SMembers(ctx context.Context, key string) ([]string, error)
}

type redisClient struct {
	rc     *redis.Client
	packer Packer
}

func NewClient(opts ...Option) (RedisClient, error) {
	o := defaultOptions()
	for _, opt := range opts {
		opt(o)
	}

	rc := redis.NewClient(&redis.Options{
		Addr:       fmt.Sprintf("%s:%d", o.host, o.port),
		Username:   o.user,
		Password:   o.password,
		DB:         o.db,
		MaxRetries: 3,
	})
	if err := rc.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	if o.tracer != nil {
		rc.AddHook(NewTraceHook(o.tracer))
	}

	return &redisClient{
		rc:     rc,
		packer: o.packer,
	}, nil
}

func (r *redisClient) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	data, err := r.packer.Marshal(value)
	if err != nil {
		return err
	}
	return r.rc.Set(ctx, key, data, expiration).Err()
}

func (r *redisClient) Get(ctx context.Context, key string, target interface{}) (found bool, err error) {
	data, err := r.rc.Get(ctx, key).Bytes()
	if err != nil {
		if err == redis.Nil {
			return false, nil
		}
		return false, err
	}
	return true, r.packer.Unmarshal(data, target)
}

func (r *redisClient) Del(ctx context.Context, key string) error {
	return r.rc.Del(ctx, key).Err()
}

func (r *redisClient) SMembers(ctx context.Context, key string) ([]string, error) {
	return r.rc.SMembers(ctx, key).Result()
}
