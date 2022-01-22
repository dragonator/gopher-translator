package handlers

import (
	"encoding/json"
	"net/http"

	v1 "github.com/dragonator/gopher-translator/internal/contracts/v1"
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

// jsonError -
func jsonError(w http.ResponseWriter, code int, er *v1.ErrorResponse) {
	w.Header().Set(ContentTypeHeaderName, ContentTypeJSON)
	w.Header().Set(XContentTypeOptions, NoSniff)
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(er)
}

func success(w http.ResponseWriter, resp interface{}) {
	w.Header().Set(ContentTypeHeaderName, ContentTypeJSON)
	json.NewEncoder(w).Encode(resp)
}

func decode(r *http.Request, v interface{}) error {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()
	return d.Decode(v)
}
