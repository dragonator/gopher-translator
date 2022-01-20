package storage

import (
	"testing"

	"github.com/dragonator/gopher-translator/internal/models"
)

func TestNew(t *testing.T) {
	s := New()
	if s == nil {
		t.Error("unexpected nil value")
	}
}

func TestAddRecord(t *testing.T) {
	// setup
	s := storage{}
	input := []*models.Record{
		{EnglishWord: "c", GopherWord: "3"},
		{EnglishWord: "a", GopherWord: "1"},
		{EnglishWord: "b", GopherWord: "2"},
	}
	// call
	for _, record := range input {
		s.AddRecord(record)
	}
	// assert
	expected := []*models.Record{
		{EnglishWord: "a", GopherWord: "1"},
		{EnglishWord: "b", GopherWord: "2"},
		{EnglishWord: "c", GopherWord: "3"},
	}
	for i, record := range s.store {
		if record.EnglishWord != expected[i].EnglishWord ||
			record.GopherWord != expected[i].GopherWord {
			t.Errorf("unexpected values for record: %v (expected: %v)", record, expected[i])
		}
	}
}

func TestHistory(t *testing.T) {
	// setup
	s := &storage{
		store: []*models.Record{{EnglishWord: "test"}},
	}
	// call
	h := s.History()
	// assert
	for i, r := range h {
		if r != s.store[i] {
			t.Errorf("unexpected value for record: %v (expected: %v)", r, s.store[i])
		}
	}
}
