package db

import (
	"context"
	"errors"
	"time"

	"gorm.io/gorm"
)

type Helper[T any] struct {
	db *gorm.DB
}

func NewHelper[T any](db *gorm.DB) *Helper[T] {
	return &Helper[T]{db: db}
}

func (h *Helper[T]) clone(db *gorm.DB) *Helper[T] {
	return &Helper[T]{db: db}
}

func (h *Helper[T]) Debug() *Helper[T] {
	return h.clone(h.db.Debug())
}

func (h *Helper[T]) WithCache(key string, expires ...time.Duration) *Helper[T] {
	expire := 10 * time.Second
	if len(expires) > 0 && expires[0] > 0 {
		expire = expires[0]
	}

	if expire > 0 {
		return h.clone(h.db.Set(CacheParamKey, CacheParam{Key: key, Expires: expire}))
	}

	return h
}

func (h *Helper[T]) Transaction(handler func(hx *Helper[T]) error) error {
	return h.db.Transaction(func(tx *gorm.DB) error {
		hx := h.clone(tx)
		return handler(hx)
	})
}

func (h *Helper[T]) Preload(query string, args ...interface{}) *Helper[T] {
	return h.clone(h.db.Preload(query, args...))
}

func (h *Helper[T]) Where(query interface{}, args ...interface{}) *Helper[T] {
	return h.clone(h.db.Where(query, args...))
}

func (h *Helper[T]) Limit(limit int) *Helper[T] {
	if limit > 0 {
		return h.clone(h.db.Limit(limit))
	}
	return h
}

func (h *Helper[T]) Offset(offset int) *Helper[T] {
	return h.clone(h.db.Offset(offset))
}

func (h *Helper[T]) Order(cond string) *Helper[T] {
	return h.clone(h.db.Order(cond))
}

func (h *Helper[T]) First(ctx context.Context) (*T, error) {
	var item T
	if err := h.db.First(&item).Error; err != nil {
		return nil, err
	}
	return &item, nil
}

func (h *Helper[T]) Find(ctx context.Context) ([]*T, error) {
	var items []*T
	if err := h.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (h *Helper[T]) FindWithCount(ctx context.Context) ([]*T, int64, error) {
	var items []*T
	var count int64
	if err := h.db.Find(&items).Offset(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func (h *Helper[T]) Create(ctx context.Context, target *T) error {
	return h.db.Create(target).Error
}

func (h *Helper[T]) Save(ctx context.Context, target *T) error {
	return h.db.Save(target).Error
}

func (h *Helper[T]) Update(ctx context.Context, column string, value interface{}) error {
	return h.db.Update(column, value).Error
}

func (h *Helper[T]) UpdateStruct(ctx context.Context, up T) error {
	return h.db.Updates(up).Error
}

func (h *Helper[T]) Updates(ctx context.Context, up map[string]interface{}) error {
	return h.db.Updates(up).Error
}

func (h *Helper[T]) MustUpdate(ctx context.Context, column string, value interface{}) error {
	result := h.db.Update(column, value)
	return h.must(result)
}

func (h *Helper[T]) MustUpdateStruct(ctx context.Context, up T) error {
	result := h.db.Updates(up)
	return h.must(result)
}

func (h *Helper[T]) MustUpdates(ctx context.Context, up map[string]interface{}) error {
	result := h.db.Updates(up)
	return h.must(result)
}

func (h *Helper[T]) Delete(ctx context.Context, target *T) error {
	return h.db.Delete(target).Error
}

func (h *Helper[T]) MustDelete(ctx context.Context, target *T) error {
	result := h.db.Delete(target)
	return h.must(result)
}

func (h *Helper[T]) must(result *gorm.DB) error {
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("rows not affected")
	}
	return nil
}

func IsRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
