package storage

import (
	"testing"
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
	input := []*Record{
		{Input: "c", Output: "3"},
		{Input: "a", Output: "1"},
		{Input: "b", Output: "2"},
	}
	// call
	for _, record := range input {
		s.AddRecord(record)
	}
	// assert
	expected := []*Record{
		{Input: "a", Output: "1"},
		{Input: "b", Output: "2"},
		{Input: "c", Output: "3"},
	}
	for i, record := range s.store {
		if record.Input != expected[i].Input ||
			record.Output != expected[i].Output {
			t.Errorf("unexpected values for record: %v (expected: %v)", record, expected[i])
		}
	}
}

func TestHistory(t *testing.T) {
	// setup
	s := &storage{
		store: []*Record{{Input: "test"}},
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
