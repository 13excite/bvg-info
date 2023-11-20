package cache

import (
	"errors"
	"fmt"
	"github.com/13excite/bvg-info/pkg/store"
	"github.com/bluele/gcache"
	"go.uber.org/zap"
	"time"
)

var (
	errStopNotInCache = errors.New("the stop isn't in cache")
)

const (
	cacheSize = 1_000_000
	cacheTTL  = 10 * time.Minute // default expiration
)

type gCache struct {
	stops  gcache.Cache
	logger *zap.SugaredLogger
}

func NewGCache() *gCache {
	return &gCache{
		stops:  gcache.New(cacheSize).Expiration(cacheTTL).ARC().Build(),
		logger: zap.S().With("package", "cache"),
	}
}

func (gc *gCache) Update(key string, stops []store.CachedStop) error {
	if err := gc.stops.Set(key, stops); err != nil {
		gc.logger.Error("Update cache error", "error", err)
		return err
	}
	gc.logger.Info("Cache updated for key: ", key)
	return nil
}

func (gc *gCache) Read(stopName string) ([]store.CachedStop, error) {
	val, err := gc.stops.Get(stopName)
	gc.logger.Debug("Reading from cache for key: ", stopName)
	if err != nil {
		if errors.Is(err, gcache.KeyNotFoundError) {
			gc.logger.Error("Key not found error", "error", errStopNotInCache)
			return []store.CachedStop{}, errStopNotInCache
		}
		gc.logger.Error("Read cache error", "error", err)
		return []store.CachedStop{}, fmt.Errorf("get: %w", err)
	}

	return val.([]store.CachedStop), nil
}
