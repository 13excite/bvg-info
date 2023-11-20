package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/xid"
)

// ErrResponse is a struct for JSON error response
type ErrResponse struct {
	Status  string `json:"status,omitempty"`
	Error   string `json:"error,omitempty"`
	ErrorID string `json:"error_id,omitempty"`
}

// RenderJSON renders JSON response with given status code
func RenderJSON(w http.ResponseWriter, code int, v interface{}) {
	w.Header().Set("Content-Type", "application/json")
	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(b, `{"render_error":"%s"}`, errString(err))
	} else {
		w.WriteHeader(code)
	}
	_, _ = w.Write(b.Bytes())
}

func RenderErrInternal(w http.ResponseWriter, err error) {
	RenderJSON(w, http.StatusInternalServerError, ErrResponse{Status: "internal error", Error: errString(err)})
}

func RenderErrInternalWithID(w http.ResponseWriter, err error) string {
	errID := xid.New().String()
	RenderJSON(w, http.StatusInternalServerError, ErrResponse{Status: "internal error", Error: errString(err), ErrorID: errID})
	return errID
}

func errString(err error) string {
	if err == nil {
		return ""
	}
	return err.Error()
}
