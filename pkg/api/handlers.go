package api

import (
	"net/http"

	"github.com/13excite/bvg-info/pkg/store"
)

func (s *Server) GetData(w http.ResponseWriter, r *http.Request) {

	cachedStops := []store.CachedStop{}
	for stopKey := range store.NearbyDepartures() {
		stops, err := s.cache.Read(stopKey)
		if err != nil {
			//errID := RenderErrInternalWithID(w, nil)
			s.logger.Errorw("GetData handler error", "error", err)
			continue
		}
		cachedStops = append(cachedStops, stops...)
	}

	RenderJSON(w, http.StatusOK, cachedStops)
}
