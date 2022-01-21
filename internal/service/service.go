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
	CfgParamRulesFilepath = "RulesFilepath"
	CfgParamPortNumber    = "PortNumber"
)

// Service -
type Service struct {
	server *http.Server
}

// New -
func New(config map[string]string) (*Service, error) {
	requiredKeys := []string{
		CfgParamRulesFilepath,
		CfgParamPortNumber,
	}
	for _, key := range requiredKeys {
		_, ok := config[key]
		if !ok {
			return nil, fmt.Errorf("missing required configuration: %s", key)
		}
	}

	data, err := os.ReadFile(config[CfgParamRulesFilepath])
	if err != nil {
		return nil, err
	}

	var rules []*translator.Rule
	if err := json.Unmarshal(data, &rules); err != nil {
		return nil, err
	}

	tr := translator.New(rules)
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
