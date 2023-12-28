package cache

import (
	"context"
	"errors"
	"reflect"
	"sync"
	"time"

	"golang.org/x/sync/singleflight"
)

var singleSetter singleflight.Group

type CacheStore interface {
	Get(ctx context.Context, key string, target interface{}) (bool, error)
	Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error
}

type AgingData struct {
	value     interface{}
	expiredAt time.Time
}

type defaultStore struct {
	m *sync.Map
}

func DefaultStore() *defaultStore {
	return &defaultStore{
		m: &sync.Map{},
	}
}

func (s *defaultStore) Set(ctx context.Context, key string, val interface{}, expiration time.Duration) error {
	s.m.Store(key, AgingData{
		value:     val,
		expiredAt: time.Now().Add(expiration),
	})
	return nil
}

func (s *defaultStore) Get(ctx context.Context, key string, target interface{}) (bool, error) {
	val, ok := s.m.Load(key)
	if !ok {
		return false, nil
	}
	ad := val.(AgingData)
	if ad.expiredAt.Before(time.Now()) {
		s.m.Delete(key)
		return false, nil
	}

	if reflect.ValueOf(target).Type().Kind() != reflect.Ptr {
		return false, errors.New("target must be a pointer")
	}

	tv := reflect.ValueOf(target).Elem()
	tv.Set(reflect.ValueOf(ad.value))
	return true, nil
}

type Cacher struct {
	store      CacheStore
	key        string
	expiration time.Duration
	fetch      bool
	defaultGet func(ctx context.Context) (interface{}, error)
}

func WithCache(store CacheStore, key string, expiration time.Duration) *Cacher {
	return &Cacher{
		store:      store,
		key:        key,
		expiration: expiration,
	}
}

func (c *Cacher) Default(defaultGet func(ctx context.Context) (interface{}, error)) *Cacher {
	c.defaultGet = defaultGet
	return c
}

func (c *Cacher) Refresh(fetch bool) *Cacher {
	c.fetch = fetch
	return c
}

func (c *Cacher) Get(ctx context.Context, target interface{}) error {
	if !c.fetch {
		found, err := c.store.Get(ctx, c.key, target)
		if err != nil {
			return err
		}
		if found {
			return nil
		}
	}

	data, err, _ := singleSetter.Do(c.key, func() (interface{}, error) {
		val, err := c.defaultGet(ctx)
		if err != nil {
			return nil, err
		}

		if err := c.store.Set(ctx, c.key, val, c.expiration); err != nil {
			return nil, err
		}

		return val, nil
	})

	if err != nil {
		return err
	}

	if reflect.ValueOf(target).Type().Kind() != reflect.Ptr {
		return errors.New("target must be a pointer")
	}

	tv := reflect.ValueOf(target).Elem()
	tv.Set(reflect.ValueOf(data))
	return nil
}
