package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/dragonator/gopher-translator/internal/contracts/v1"
	"github.com/dragonator/gopher-translator/internal/service/svc"
	"github.com/dragonator/gopher-translator/tests/mocks"
)

func TestNew(t *testing.T) {
	s := NewGopher(nil)
	if s == nil {
		t.Error("unexpected nil handler")
	}
}

func TestTranslateWord(t *testing.T) {

	t.Run("invalid method", func(t *testing.T) {
		// setup
		w, r := createReq("invalid", "/word", nil)
		gh := &gopher{}
		// call
		gh.TranslateWord("POST", "/word")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusMethodNotAllowed)
		assertBody(t, w.Body, "405 method not allowed\n")
	})

	t.Run("invalid path", func(t *testing.T) {
		// setup
		w, r := createReq("POST", "/invalid", nil)
		gh := &gopher{}
		// call
		gh.TranslateWord("POST", "/word")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusNotFound)
		assertBody(t, w.Body, "404 page not found\n")
	})

	t.Run("json decoding failed", func(t *testing.T) {
		// setup
		w, r := createReq("POST", "/word", `{"invalid":"a"}`)
		gh := &gopher{}
		// call
		gh.TranslateWord("POST", "/word")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusBadRequest)
		assertBody(t, w.Body, &v1.ErrorResponse{Message: `request decoding failed: json: unknown field "invalid"`})
	})

	t.Run("known translate error", func(t *testing.T) {
		// setup
		req := &v1.GopherWordRequest{EnglishWord: "ap'ple"}
		w, r := createReq("POST", "/word", req)
		rm := mocks.NewGopherResourceMock()
		gh := &gopher{rs: rm}
		rm.On("TranslateWord", req.EnglishWord).Return("", svc.ErrInvalidInput)
		// call
		gh.TranslateWord("POST", "/word")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusBadRequest)
		assertBody(t, w.Body, &v1.ErrorResponse{Message: "invalid input"})
	})

	t.Run("unknown translate error", func(t *testing.T) {
		// setup
		req := &v1.GopherWordRequest{EnglishWord: "ap'ple"}
		w, r := createReq("POST", "/word", req)
		rm := mocks.NewGopherResourceMock()
		gh := &gopher{rs: rm}
		rm.On("TranslateWord", req.EnglishWord).Return("", errors.New("test error"))
		// call
		gh.TranslateWord("POST", "/word")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusInternalServerError)
		assertBody(t, w.Body, &v1.ErrorResponse{Message: "test error"})
	})

	t.Run("ok", func(t *testing.T) {
		// setup
		req := &v1.GopherWordRequest{EnglishWord: "apple"}
		w, r := createReq("POST", "/word", req)
		rm := mocks.NewGopherResourceMock()
		gh := &gopher{rs: rm}
		rm.On("TranslateWord", req.EnglishWord).Return("gapple", nil)
		// call
		gh.TranslateWord("POST", "/word")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusOK)
		assertBody(t, w.Body, &v1.GopherWordResponse{GopherWord: "gapple"})
	})
}

func TestTranslateSentence(t *testing.T) {

	t.Run("invalid method", func(t *testing.T) {
		// setup
		w, r := createReq("invalid", "/sentence", nil)
		gh := &gopher{}
		// call
		gh.TranslateSentence("POST", "/sentence")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusMethodNotAllowed)
		assertBody(t, w.Body, "405 method not allowed\n")
	})

	t.Run("invalid path", func(t *testing.T) {
		// setup
		w, r := createReq("POST", "/invalid", nil)
		gh := &gopher{}
		// call
		gh.TranslateSentence("POST", "/sentence")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusNotFound)
		assertBody(t, w.Body, "404 page not found\n")
	})

	t.Run("json decoding failed", func(t *testing.T) {
		// setup
		w, r := createReq("POST", "/sentence", `{"invalid":"a"}`)
		gh := &gopher{}
		// call
		gh.TranslateSentence("POST", "/sentence")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusBadRequest)
		assertBody(t, w.Body, &v1.ErrorResponse{Message: `request decoding failed: json: unknown field "invalid"`})
	})

	t.Run("known translate error", func(t *testing.T) {
		// setup
		req := &v1.GopherSentenceRequest{EnglishSentence: "ap'ple"}
		w, r := createReq("POST", "/sentence", req)
		rm := mocks.NewGopherResourceMock()
		gh := &gopher{rs: rm}
		rm.On("TranslateSentence", req.EnglishSentence).Return("", svc.ErrInvalidInput)
		// call
		gh.TranslateSentence("POST", "/sentence")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusBadRequest)
		assertBody(t, w.Body, &v1.ErrorResponse{Message: "invalid input"})
	})

	t.Run("unknown translate error", func(t *testing.T) {
		// setup
		req := &v1.GopherSentenceRequest{EnglishSentence: "ap'ple"}
		w, r := createReq("POST", "/sentence", req)
		rm := mocks.NewGopherResourceMock()
		gh := &gopher{rs: rm}
		rm.On("TranslateSentence", req.EnglishSentence).Return("", errors.New("test error"))
		// call
		gh.TranslateSentence("POST", "/sentence")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusInternalServerError)
		assertBody(t, w.Body, &v1.ErrorResponse{Message: "test error"})
	})

	t.Run("ok", func(t *testing.T) {
		// setup
		req := &v1.GopherSentenceRequest{EnglishSentence: "apple"}
		w, r := createReq("POST", "/sentence", req)
		rm := mocks.NewGopherResourceMock()
		gh := &gopher{rs: rm}
		rm.On("TranslateSentence", req.EnglishSentence).Return("gapple", nil)
		// call
		gh.TranslateSentence("POST", "/sentence")(w, r)
		// assert
		assertCode(t, w.Code, http.StatusOK)
		assertBody(t, w.Body, &v1.GopherSentenceResponse{GopherSentence: "gapple"})
	})
}

func TestHistory(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		path         string
		expectedCode int
		expectedBody interface{}
	}{
		{"invalid method", "invalid", "/history", http.StatusMethodNotAllowed, "405 method not allowed\n"},
		{"invalid path", "GET", "/invalid", http.StatusNotFound, "404 page not found\n"},
		{"ok", "GET", "/history", http.StatusOK, &v1.HistoryResponse{{"apple": "gapple."}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			w, r := createReq(tc.method, tc.path, nil)
			rm := mocks.NewGopherResourceMock()
			gh := &gopher{rs: rm}
			if tc.expectedCode == http.StatusOK {
				exp := tc.expectedBody.(*v1.HistoryResponse)
				rm.On("History", nil).Return(([]map[string]string)(*exp))
			}
			// call
			gh.History("GET", "/history")(w, r)
			// assert
			rm.AssertExpectations(t)
			assertCode(t, w.Code, tc.expectedCode)
			assertBody(t, w.Body, tc.expectedBody)
		})
	}
}

func createReq(method, path string, body interface{}) (*httptest.ResponseRecorder, *http.Request) {
	w := httptest.NewRecorder()
	var b []byte
	if body != nil {
		switch v := body.(type) {
		case string:
			b = []byte(v)
		default:
			b, _ = json.Marshal(body)
		}
	}
	return w, httptest.NewRequest(method, path, bytes.NewBuffer(b))
}

func assertCode(t *testing.T, code, expectedCode int) {
	if code != expectedCode {
		t.Errorf("unexpected status code: %d - %s (expected: %d - %s)",
			code, http.StatusText(code),
			expectedCode, http.StatusText(expectedCode))
	}
}

func assertBody(t *testing.T, body *bytes.Buffer, expectedBody interface{}) {
	switch expected := expectedBody.(type) {
	case string:
		bodyStr := body.String()
		if expected != bodyStr {
			t.Errorf("unexpected error message: %s (expected: %s)", bodyStr, expected)
		}

	case *v1.ErrorResponse:
		er := &v1.ErrorResponse{}
		err := json.Unmarshal(body.Bytes(), er)
		if err != nil {
			t.Errorf("unexpected failure unmarshaling response: %v", err)
		}

		if expected.Message != er.Message {
			t.Errorf("unexpected error message: %s (expected: %s)", er.Message, expected.Message)
		}

	case *v1.GopherWordResponse:
		resp := &v1.GopherWordResponse{}
		err := json.Unmarshal(body.Bytes(), resp)
		if err != nil {
			t.Errorf("unexpected failure unmarshaling response: %v", err)
		}

		if resp.GopherWord != expected.GopherWord {
			t.Errorf("unexpected gopher translation: %s (expected: %s)", resp.GopherWord, expected.GopherWord)
		}

	case *v1.GopherSentenceResponse:
		resp := &v1.GopherSentenceResponse{}
		err := json.Unmarshal(body.Bytes(), resp)
		if err != nil {
			t.Errorf("unexpected failure unmarshaling response: %v", err)
		}

		if resp.GopherSentence != expected.GopherSentence {
			t.Errorf("unexpected gopher translation: %s (expected: %s)", resp.GopherSentence, expected.GopherSentence)
		}

	case *v1.HistoryResponse:
		resp := &v1.HistoryResponse{}
		err := json.Unmarshal(body.Bytes(), resp)
		if err != nil {
			t.Errorf("unexpected failure unmarshaling response: %v", err)
		}

		if len(*resp) != len(*expected) {
			t.Errorf("unexpected history records count: %d (expected: %d)", len(*resp), len(*expected))
		}

		expMapSlice := ([]map[string]string)(*expected)
		for i, record := range *resp {
			for in, out := range record {
				eOut, ok := expMapSlice[i][in]
				if !ok {
					t.Errorf("unexpected record input: %s", out)
				}
				if out != eOut {
					t.Errorf("unexpected record output: %s (expeced: %s)", out, eOut)
				}
			}
		}

	default:
		t.Error("unexpected response format")
	}
}
