package mocks

import "github.com/dragonator/gopher-translator/internal/storage"

func compareStrings(a, b interface{}) bool {
	return a.(string) == b.(string)
}

func compareRecords(a, b interface{}) bool {
	ar := a.(*storage.Record)
	br := b.(*storage.Record)
	if ar.Input != br.Input ||
		ar.Output != br.Output {
		return false
	}
	return true
}

func compareNils(a, b interface{}) bool {
	return a == b
}
