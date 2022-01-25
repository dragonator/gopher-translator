package resources

import (
	"errors"
	"strings"
	"testing"

	"github.com/dragonator/gopher-translator/internal/service/svc"
	"github.com/dragonator/gopher-translator/internal/storage"
	"github.com/dragonator/gopher-translator/tests/mocks"
)

func TestNewGopher(t *testing.T) {
	s := NewGopher(nil, nil)
	if s == nil {
		t.Error("unexpected nil resource")
	}
}

func TestTranslateWord(t *testing.T) {
	testCases := []struct {
		name       string
		input      string
		expected   string
		shouldFail bool
	}{
		{"invalid word", "ap'ple", "", true},
		{"ok", "apple", "|apple|", false},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tm := mocks.NewTranslatorMock()
			sm := mocks.NewStorageMock()

			var translateErr error
			if tc.shouldFail {
				translateErr = errors.New("test error")
			} else {
				sm.On("AddRecord", &storage.Record{Input: tc.input, Output: tc.expected})
			}
			tm.On("Translate", tc.input).Return("|"+tc.input+"|", translateErr)

			gopher := &gopher{
				translator: tm,
				store:      sm,
			}
			// call
			res, err := gopher.TranslateWord(tc.input)
			// assert
			assertError(t, err, tc.shouldFail)
			tm.AssertExpectations(t)
			sm.AssertExpectations(t)
			if res != tc.expected {
				t.Errorf("unexpected result: %s (expected: %s)", res, tc.expected)
			}
		})
	}
}

func TestTranslateSentence(t *testing.T) {
	testCases := []struct {
		name         string
		input        string
		expected     string
		indexErr     int
		translateErr error
		shouldFail   bool
	}{
		{"boring sentence", "boring sentence.", "|boring| |sentence|.", 0, nil, false},
		{"sentence with commas", "this one, has comma!", "|this| |one,| |has| |comma|!", 0, nil, false},
		{"invalid character", "this contains inva'lid character!", "|this| |contains| |character|!", 2, svc.ErrInvalidInput, false},
		{"other translate error", "this contains inva'lid character!", "", 2, errors.New("test error"), true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			tm := mocks.NewTranslatorMock()
			words := strings.Split(tc.input[:len(tc.input)-1], " ")
			for i, word := range words {
				var err error
				if i == tc.indexErr {
					err = tc.translateErr
				}
				tm.On("Translate", word).Return("|"+word+"|", err)
			}

			sm := mocks.NewStorageMock()
			sm.On("AddRecord", &storage.Record{Input: tc.input, Output: tc.expected})

			gopher := &gopher{
				translator: tm,
				store:      sm,
			}
			// call
			res, err := gopher.TranslateSentence(tc.input)
			// assert
			assertError(t, err, tc.shouldFail)
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

func assertError(t *testing.T, err error, exists bool) {
	if exists && err == nil {
		t.Errorf("expected error: got nil")
	} else if !exists && err != nil {
		t.Errorf("unexpected error: %s)", err)
	}
}
