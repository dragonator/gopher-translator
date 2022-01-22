package resources

import (
	"strings"
	"testing"

	"github.com/dragonator/gopher-translator/internal/storage"
	"github.com/dragonator/gopher-translator/tests/mocks"
)

func TestNewGopher(t *testing.T) {
	s := NewGopher(nil, nil)
	if s == nil {
		t.Error("unexpected nil value")
	}
}

func TestTranslateWord(t *testing.T) {
	// setup
	input := "apple"
	expected := "gapple"
	tm := mocks.NewTranslatorMock()
	tm.On("Translate", input).Return(expected)
	sm := mocks.NewStorageMock()
	sm.On("AddRecord", &storage.Record{Input: input, Output: expected})
	gopher := &gopher{
		translator: tm,
		store:      sm,
	}
	// call
	res := gopher.TranslateWord(input)
	// assert
	tm.AssertExpectations(t)
	sm.AssertExpectations(t)
	if res != expected {
		t.Errorf("unexpected result: %s (expected: %s)", res, expected)
	}
}

func TestTranslateSentence(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
	}{
		{"boring sentence", "boring sentence.", "|boring| |sentence|."},
		{"sentence with commas", "this one, has comma!", "|this| |one,| |has| |comma|!"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			tm := mocks.NewTranslatorMock()
			words := strings.Split(tc.input[:len(tc.input)-1], " ")
			for _, word := range words {
				tm.On("Translate", word).Return("|" + word + "|")
			}

			sm := mocks.NewStorageMock()
			sm.On("AddRecord", &storage.Record{Input: tc.input, Output: tc.expected})

			gopher := &gopher{
				translator: tm,
				store:      sm,
			}
			// call
			res := gopher.TranslateSentence(tc.input)
			// assert
			if res != tc.expected {
				t.Errorf("unexpected translation: %s (expected: %s)", res, tc.expected)
			}
		})
	}
}

func TestHistory(t *testing.T) {
	// setup
	sm := mocks.NewStorageMock()
	sm.On("History", nil).Return([]*storage.Record{
		{Input: "a", Output: "1"},
		{Input: "b", Output: "2"},
	})
	gopher := &gopher{
		store: sm,
	}
	// call
	res := gopher.History()
	// assert
	expected := []map[string]string{
		{"a": "1"},
		{"b": "2"},
	}
	if len(expected) != len(res) {
		t.Errorf("unexpected records count: %d (expected: %d)", len(res), len(expected))
		t.FailNow()
	}
	for i, er := range expected {
		for ek, ev := range er {
			av, ok := res[i][ek]
			if !ok {
				t.Errorf("missing expected key for record (expected: %s)", ek)
				continue
			}
			if av != ev {
				t.Errorf("unexpected value for record: %s (expected: %s)", av, ev)
			}
		}
	}
}
