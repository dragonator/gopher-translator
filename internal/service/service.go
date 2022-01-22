package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

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

// Service -
type Service struct {
	server *http.Server
}

// New -
func New(config map[string]string) (*Service, error) {
	requiredKeys := []string{
		CfgParamSpecFile,
		CfgParamPortNumber,
	}
	for _, key := range requiredKeys {
		_, ok := config[key]
		if !ok {
			return nil, fmt.Errorf("missing required configuration: %s", key)
		}
	}

	data, err := os.ReadFile(config[CfgParamSpecFile])
	if err != nil {
		return nil, err
	}

	spec := &translator.Specification{}
	if err := json.Unmarshal(data, spec); err != nil {
		return nil, err
	}

	tr := translator.New(spec)
	store := storage.New()
	gr := resources.NewGopher(tr, store)
	gh := handlers.NewGopher(gr)
	router := NewRouter(gh)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", config[CfgParamPortNumber]),
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
