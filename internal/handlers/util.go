package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	v1 "github.com/dragonator/gopher-translator/internal/contracts/v1"
	"github.com/dragonator/gopher-translator/internal/service/svc"
)

// Header constants -
const (
	ContentTypeHeaderName = "Content-Type"
	ContentTypeJSON       = "application/json"
	XContentTypeOptions   = "X-Content-Type-Options"
	NoSniff               = "nosniff"
)

func validMethodAndPath(w http.ResponseWriter, r *http.Request, method, path string) bool {
	if r.URL.Path != path {
		w.Header().Set(ContentTypeHeaderName, ContentTypeJSON)
		http.NotFound(w, r)
		return false
	}
	if r.Method != method {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

// errorResponse -
func errorResponse(w http.ResponseWriter, err error) {
	er := &v1.ErrorResponse{Message: err.Error()}
	w.Header().Set(ContentTypeHeaderName, ContentTypeJSON)
	w.Header().Set(XContentTypeOptions, NoSniff)

	var e *svc.Error
	if errors.As(err, &e) {
		w.WriteHeader(e.StatusCode)
		json.NewEncoder(w).Encode(er)
		return
	}

	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(er)
}

func successResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set(ContentTypeHeaderName, ContentTypeJSON)
	json.NewEncoder(w).Encode(resp)
}

func decode(r *http.Request, v interface{}) error {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	return d.Decode(v)
}
