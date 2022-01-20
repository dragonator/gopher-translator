package storage

import (
	"sort"

	"github.com/dragonator/gopher-translator/internal/models"
)

// Storage -
type Storage interface {
	AddRecord(r *models.Record)
	History() []*models.Record
}

type storage struct {
	store []*models.Record
}

// New -
func New() Storage {
	return &storage{}
}

func (s *storage) AddRecord(r *models.Record) {
	i := sort.Search(len(s.store), func(i int) bool {
		return s.store[i].EnglishWord >= r.EnglishWord
	})
	s.store = append(s.store, nil)
	copy(s.store[i+1:], s.store[i:])
	s.store[i] = r
}

func (s *storage) History() []*models.Record {
	return s.store
}
