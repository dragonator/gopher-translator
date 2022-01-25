package handlers

import (
	"fmt"
	"net/http"

	v1 "github.com/dragonator/gopher-translator/internal/contracts/v1"
	"github.com/dragonator/gopher-translator/internal/resources"
	"github.com/dragonator/gopher-translator/internal/service/svc"
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
			errorResponse(w, fmt.Errorf("%w: %v", svc.ErrDecodeRequest, err))
			return
		}

		translation, err := gh.rs.TranslateWord(req.EnglishWord)
		if err != nil {
			errorResponse(w, err)
			return
		}

		resp := &v1.GopherWordResponse{
			GopherWord: translation,
		}
		successResponse(w, resp)
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
			errorResponse(w, fmt.Errorf("%w: %v", svc.ErrDecodeRequest, err))
			return
		}

		translation, err := gh.rs.TranslateSentence(req.EnglishSentence)
		if err != nil {
			errorResponse(w, err)
			return
		}

		resp := &v1.GopherSentenceResponse{
			GopherSentence: translation,
		}
		successResponse(w, resp)
		return
	}
}

func (gh *gopher) History(method, path string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if !validMethodAndPath(w, r, method, path) {
			return
		}

		resp := gh.rs.History()
		successResponse(w, resp)
		return
	}
}
