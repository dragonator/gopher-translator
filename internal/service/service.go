package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/dragonator/gopher-translator/internal/handlers"
	"github.com/dragonator/gopher-translator/internal/resources"
	"github.com/dragonator/gopher-translator/internal/storage"
	"github.com/dragonator/gopher-translator/internal/translator"
)

// Constants -
const (
	CfgParamSpecFile   = "SpecFile"
	CfgParamPortNumber = "PortNumber"
)

// Bootstrap -
type Bootstrap struct {
	Port string
	Spec io.Reader
}

// Service -
type Service struct {
	server *http.Server
}

// New -
func New(b *Bootstrap) (*Service, error) {
	spec := &translator.Specification{}
	d := json.NewDecoder(b.Spec)
	d.DisallowUnknownFields()
	if err := d.Decode(spec); err != nil {
		return nil, err
	}

	tr := translator.New(spec)
	store := storage.New()
	gr := resources.NewGopher(tr, store)
	gh := handlers.NewGopher(gr)
	router := NewRouter(gh)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", b.Port),
		Handler: router,
	}

	return &Service{server: srv}, nil
}

// Start -
func (s *Service) Start() {
	go func() {
		if err := s.server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe() failed: %v", err)
		}
	}()
	log.Printf("Listening on port %s ...", s.server.Addr)
}

// Stop -
func (s *Service) Stop() {
	if err := s.server.Shutdown(context.Background()); err != nil {
		log.Fatalf("Shutdown() failed: %v", err)
	} else {
		log.Print("HTTP server shut down gracefully.")
	}
}
