package cache

import (
	"github.com/13excite/bvg-info/pkg/store"
)

type Cache interface {
	Update(store.CachedStops) error
	Read(string) (store.CachedStops, error)
}
