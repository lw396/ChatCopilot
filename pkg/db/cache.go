package db

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"github.com/lw396/ChatCopilot/pkg/cache"

	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
	"gorm.io/gorm/utils"
)

const (
	PluginCache   = "PLUGIN_CACHE"
	CacheParamKey = "gorm_cache_param"
	CachePrefix   = "_GORM_CACHE_"
)

type contextKey string

const (
	KeyUseCache contextKey = "USE_CACHE"
)

type CacheParam struct {
	Key     string
	Expires time.Duration
}

type CachePlugin struct {
	store cache.CacheStore
}

type QueryResult struct {
	RowsAffected int64
	Value        interface{}
}

func NewCachePlugin(store cache.CacheStore) CachePlugin {
	return CachePlugin{store: store}
}

func (CachePlugin) Name() string {
	return PluginCache
}

func (cp CachePlugin) Initialize(db *gorm.DB) error {
	_ = db.Callback().Query().Replace("gorm:query", cp.tryCache)

	_ = db.Callback().Query().After("gorm:after_query").Register("store_cache", cp.storeCache)
	return nil
}

func (cp CachePlugin) tryCache(db *gorm.DB) {
	param, ok := cp.cacheParam(db)
	if ok {
		ctx := db.Statement.Context
		key := fmt.Sprintf("%s%s", CachePrefix, param.Key)
		qr := QueryResult{
			Value: reflect.New(db.Statement.ReflectValue.Type()).Interface(),
		}
		exist, err := cp.store.Get(ctx, key, &qr)
		if err != nil {
			db.Logger.Warn(ctx, "%s get query cache failed: %v [%s]", utils.FileWithLineNum(), err, param.Key)
		}

		if exist {
			db.Logger.Info(ctx, "%s query use cache [%s]", utils.FileWithLineNum(), param.Key)
			db.Statement.RowsAffected = qr.RowsAffected
			if reflect.ValueOf(qr.Value).Kind() == reflect.Ptr {
				db.Statement.ReflectValue.Set(reflect.ValueOf(qr.Value).Elem())
			} else {
				db.Statement.ReflectValue.Set(reflect.ValueOf(qr.Value))
			}

			db.Statement.Context = context.WithValue(ctx, KeyUseCache, true)
			return
		}
	}

	callbacks.Query(db)
}

func (cp CachePlugin) storeCache(db *gorm.DB) {
	param, ok := cp.cacheParam(db)
	if !ok {
		return
	}

	ctx := db.Statement.Context

	// 结果使用缓存获得，无需保存
	if val := ctx.Value(KeyUseCache); val != nil {
		if val.(bool) {
			return
		}
	}

	key := fmt.Sprintf("%s%s", CachePrefix, param.Key)

	qr := QueryResult{
		RowsAffected: db.Statement.RowsAffected,
		Value:        db.Statement.ReflectValue.Interface(),
	}
	if err := cp.store.Set(ctx, key, qr, param.Expires); err != nil {
		db.Logger.Warn(ctx, "%s store query cache failed: %v [%s]", utils.FileWithLineNum(), err, param.Key)
	} else {
		db.Logger.Info(ctx, "%s store query cache [%s]", utils.FileWithLineNum(), param.Key)
	}
}

func (cp CachePlugin) cacheParam(db *gorm.DB) (CacheParam, bool) {
	val, ok := db.Get(CacheParamKey)
	if !ok {
		return CacheParam{}, false
	}
	param, ok := val.(CacheParam)
	if !ok {
		return CacheParam{}, false
	}
	return param, true
}
