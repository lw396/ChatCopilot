package redis

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient interface {
	// String
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	Get(ctx context.Context, key string, target interface{}) (bool, error)
	Del(ctx context.Context, key string) error

	// Set
	SAdd(ctx context.Context, key string, members ...interface{}) error
	SMembers(ctx context.Context, key string, target interface{}) (bool, error)
	SRem(ctx context.Context, key string, members ...interface{}) error
	SUpdate(ctx context.Context, key string, members ...interface{}) error
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
		return
	}
	return true, r.packer.Unmarshal(data, target)
}

func (r *redisClient) Del(ctx context.Context, key string) error {
	return r.rc.Del(ctx, key).Err()
}

func (r *redisClient) SAdd(ctx context.Context, key string, members ...interface{}) error {
	dataList := make([]interface{}, len(members))
	for i, member := range members {
		data, err := r.packer.Marshal(member)
		if err != nil {
			return err
		}
		dataList[i] = data
	}

	if err := r.rc.SAdd(ctx, key, dataList...).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisClient) SMembers(ctx context.Context, key string, target interface{}) (found bool, err error) {
	val, err := r.rc.SMembers(ctx, key).Result()
	if err != nil {
		if err == redis.Nil {
			return
		}
		return
	}

	data := "[" + strings.Join(val, ",") + "]"
	if err = r.packer.UnmarshalFromString(data, target); err != nil {
		return
	}
	return true, nil
}

func (r *redisClient) SRem(ctx context.Context, key string, members ...interface{}) error {
	dataList := make([]interface{}, len(members))
	for i, member := range members {
		data, err := r.packer.Marshal(member)
		if err != nil {
			return err
		}
		dataList[i] = data
	}

	if err := r.rc.SRem(ctx, key, dataList).Err(); err != nil {
		return err
	}
	return nil
}

func (r *redisClient) SUpdate(ctx context.Context, key string, members ...interface{}) error {
	numMembers := len(members)
	if numMembers%2 != 0 {
		return errors.New("the number of members must be even")
	}

	oldMembers := make([]interface{}, numMembers/2)
	newMembers := make([]interface{}, numMembers/2)
	for i, member := range members {
		data, err := r.packer.Marshal(member)
		if err != nil {
			return err
		}
		if i%2 == 0 {
			oldMembers[i/2] = data
		} else {
			newMembers[i/2] = data
		}
	}

	pipe := r.rc.TxPipeline()
	pipe.SRem(ctx, key, oldMembers...)
	pipe.SAdd(ctx, key, newMembers...)
	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}
