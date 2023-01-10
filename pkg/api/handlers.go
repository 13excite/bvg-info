package api

import (
	"net/http"
)

func (s *Server) GetData(w http.ResponseWriter, r *http.Request) {

	// hardcode key for while
	stops, err := s.cache.read("sudost_konigsheide")
	if err != nil {
		errID := RenderErrInternalWithID(w, nil)
		s.logger.Errorw("GetData handler error", "error", err, "error_id", errID)
		return
	}

	RenderJSON(w, http.StatusOK, stops.Departes)
}
