package mocks

import "github.com/dragonator/gopher-translator/internal/storage"

// StorageMock -
type StorageMock struct {
	*BaseMock
}

// NewStorageMock -
func NewStorageMock() *StorageMock {
	return &StorageMock{
		BaseMock: NewBaseMock(),
	}
}

// AddRecord -
func (tm *StorageMock) AddRecord(r *storage.Record) {
	_ = tm.MarkCalledAndReturn("AddRecord", r, compareRecords)
}

// History -
func (tm *StorageMock) History() []*storage.Record {
	v, ok := tm.MarkCalledAndReturn("History", nil, compareNils).([]*storage.Record)
	if !ok {
		panic("unexpected return value type")
	}
	return v
}
