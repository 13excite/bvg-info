package api

import (
	"fmt"
	"github.com/13excite/bvg-info/pkg/conf"
	"github.com/13excite/bvg-info/pkg/store"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net"
	"net/http"
)

type cacheStore interface {
	update(store.CachedStops) error
	read(string) (store.CachedStops, error)
}

type Server struct {
	router *mux.Router
	server *http.Server
	logger *zap.SugaredLogger
	cache  cacheStore
}

func New(config *conf.Config, cache *cacheStore) *Server {
	r := mux.NewRouter()
	// put mw func here
	r.Use(RequestID)
	r.Use(loggerHTTPMiddlewareDefault(config.LoggerDisabledHttp))

	return &Server{
		logger: zap.S().With("package", "server"),
		cache:  *cache,
		router: r,
	}
}

func (s *Server) ListenAndServe(config *conf.Config) error {
	s.server = &http.Server{
		Addr:    net.JoinHostPort(config.ServerHost, config.ServerPort),
		Handler: s.router,
	}
	s.logger.Infow(s.server.Addr)

	// Listen
	listener, err := net.Listen("tcp", s.server.Addr)
	if err != nil {
		return fmt.Errorf("Could not listen on %s: %v", s.server.Addr, err)
	}

	go func() {
		if err = s.server.Serve(listener); err != nil {
			s.logger.Fatalw("API Listen error", "error", err, "address", s.server.Addr)
		}
	}()
	s.logger.Infow("API Listening", "address", s.server.Addr)

	return nil
}

// Router returns the router
func (s *Server) Router() *mux.Router {
	return s.router
}
