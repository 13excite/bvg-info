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
	errUserNotInCache = errors.New("the stop isn't in cache")
)

const (
	cacheSize = 1_000_000
	cacheTTL  = 10 * time.Minute // default expiration
)

type gCache struct {
	stops  gcache.Cache
	logger *zap.SugaredLogger
}

type stops struct {
	Name     string
	Departes []store.StopDepartures
}

func NewGCache() *gCache {
	return &gCache{
		stops:  gcache.New(cacheSize).Expiration(cacheTTL).ARC().Build(),
		logger: zap.S().With("package", "cache"),
	}
}

func (gc *gCache) update(s stops, expireIn time.Duration) error {
	if err := gc.stops.SetWithExpire(s.Name, s, expireIn); err != nil {
		gc.logger.Error("Update cache error", "error", err)
		return err
	}
	gc.logger.Info("Cache updated for key: ", s.Name)
	return nil
}

func (gc *gCache) read(stopName string) (stops, error) {
	val, err := gc.stops.Get(stopName)
	gc.logger.Info("Reading from cache for key: ", stopName)
	if err != nil {
		if errors.Is(err, gcache.KeyNotFoundError) {
			gc.logger.Error("Read cache error", "error", errUserNotInCache)
			return stops{}, errUserNotInCache
		}
		gc.logger.Error("Read cache error", "error", err)
		return stops{}, fmt.Errorf("get: %w", err)
	}

	return val.(stops), nil
}
