package service

import (
	"net/http"

	"github.com/dragonator/gopher-translator/internal/handlers"
)

// NewRouter -
func NewRouter(gh handlers.Gopher) http.Handler {
	api := []struct {
		Method     string
		Path       string
		HandleFunc func(string, string) func(w http.ResponseWriter, r *http.Request)
	}{
		{"POST", "/word", gh.TranslateWord},
		{"POST", "/sentence", gh.TranslateSentence},
		{"GET", "/history", gh.History},
	}

	mux := http.NewServeMux()
	for _, endpoint := range api {
		mux.HandleFunc(endpoint.Path, endpoint.HandleFunc(endpoint.Method, endpoint.Path))
	}

	return mux
}
