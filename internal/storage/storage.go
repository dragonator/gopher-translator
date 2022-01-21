package storage

import (
	"sort"
)

// Record -
type Record struct {
	Input  string
	Output string
}

// Storage -
type Storage interface {
	AddRecord(r *Record)
	History() []*Record
}

type storage struct {
	store []*Record
}

// New -
func New() Storage {
	return &storage{}
}

func (s *storage) AddRecord(r *Record) {
	i := sort.Search(len(s.store), func(i int) bool {
		return s.store[i].Input >= r.Input
	})
	s.store = append(s.store, nil)
	copy(s.store[i+1:], s.store[i:])
	s.store[i] = r
}

func (s *storage) History() []*Record {
	return s.store
}
