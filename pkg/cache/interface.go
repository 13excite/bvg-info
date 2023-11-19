package cache

import (
	"github.com/13excite/bvg-info/pkg/store"
)

type Cache interface {
	Update(key string, stops []store.CachedStop) error
	Read(string) ([]store.CachedStop, error)
}
