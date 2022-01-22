package handlers

import (
	"net/http"

	v1 "github.com/dragonator/gopher-translator/internal/contracts/v1"
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

		req := &v1.GopherWordRequest{}
		if err := decode(r, req); err != nil {
			jsonError(w, http.StatusBadRequest, &v1.ErrorResponse{Message: err.Error()})
			return
		}

		resp := &v1.GopherWordResponse{
			GopherWord: gh.rs.TranslateWord(req.EnglishWord),
		}
		success(w, resp)
		return
	}
}

func (gh *gopher) TranslateSentence(method, path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validMethodAndPath(w, r, method, path) {
			return
		}

		req := &v1.GopherSentenceRequest{}
		if err := decode(r, req); err != nil {
			jsonError(w, http.StatusBadRequest, &v1.ErrorResponse{Message: err.Error()})
			return
		}

		resp := &v1.GopherSentenceResponse{
			GopherSentence: gh.rs.TranslateSentence(req.EnglishSentence),
		}
		success(w, resp)
		return
	}
}

func (gh *gopher) History(method, path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validMethodAndPath(w, r, method, path) {
			return
		}

		resp := gh.rs.History()
		success(w, resp)
		return
	}
}
