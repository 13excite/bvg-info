package api

import (
	"github.com/13excite/bvg-info/pkg/store"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
)

type CacheStore interface {
	update(store.CachedStops) error
	read(string) (store.CachedStops, error)
}

type Server struct {
	router *mux.Router
	server *http.Server
	logger *zap.SugaredLogger
	cache  CacheStore
}

func New(cache *CacheStore) *Server {
	r := mux.NewRouter()
	// put mw func here
	r.Use(RequestID)

	return &Server{
		logger: zap.S().With("package", "server"),
		cache:  *cache,
	}
}
