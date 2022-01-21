package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dragonator/gopher-translator/internal/contracts"
	"github.com/dragonator/gopher-translator/internal/resources"
)

// Gopher -
type Gopher interface {
	TranslateWord(method, path string) func(w http.ResponseWriter, r *http.Request)
	TranslateSentence(method, path string) func(w http.ResponseWriter, r *http.Request)
	History(method, path string) func(w http.ResponseWriter, r *http.Request)
}

type gopher struct {
	rs resources.Gopher
}

// NewGopher -
func NewGopher(rs resources.Gopher) Gopher {
	return &gopher{
		rs: rs,
	}
}

func (gh *gopher) TranslateWord(method, path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validMethodAndPath(w, r, method, path) {
			return
		}

		req := &contracts.GopherWordRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := &contracts.GopherWordResponse{
			GopherWord: gh.rs.TranslateWord(req.EnglishWord),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
}

func (gh *gopher) TranslateSentence(method, path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validMethodAndPath(w, r, method, path) {
			return
		}

		req := &contracts.GopherSentenceRequest{}
		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		resp := &contracts.GopherSentenceResponse{
			GopherSentence: gh.rs.TranslateSentence(req.EnglishSentence),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
		return
	}
}

func (gh *gopher) History(method, path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validMethodAndPath(w, r, method, path) {
			return
		}

		h := gh.rs.History()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(h)
		return
	}
}

func validMethodAndPath(w http.ResponseWriter, r *http.Request, method, path string) bool {
	if r.URL.Path != path {
		http.NotFound(w, r)
		return false
	}
	if r.Method != method {
		http.Error(w, fmt.Sprintf("Only %s requests are allowed!", method), http.StatusMethodNotAllowed)
		return false
	}
	return true
}
