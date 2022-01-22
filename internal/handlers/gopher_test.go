package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/dragonator/gopher-translator/internal/contracts/v1"
	"github.com/dragonator/gopher-translator/tests/mocks"
)

func TestNew(t *testing.T) {
	s := NewGopher(nil)
	if s == nil {
		t.Error("unexpected nil handler")
	}
}

func TestTranslateWord(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		path         string
		body         interface{}
		expectedCode int
		expectedBody interface{}
	}{
		{"invalid method", "invalid", "/word", nil, http.StatusMethodNotAllowed, "405 method not allowed\n"},
		{"invalid path", "POST", "/invalid", nil, http.StatusNotFound, "404 page not found\n"},
		{"json decoding failed", "POST", "/word", `{"invalid":"a"}`, http.StatusBadRequest, &v1.ErrorResponse{Message: `json: unknown field "invalid"`}},
		{"ok", "POST", "/word", &v1.GopherWordRequest{EnglishWord: "apple"}, http.StatusOK, &v1.GopherWordResponse{GopherWord: "gapple"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			w, r := createReq(tc.method, tc.path, tc.body)
			rm := mocks.NewGopherResourceMock()
			gh := &gopher{rs: rm}
			if tc.expectedCode == http.StatusOK {
				req := tc.body.(*v1.GopherWordRequest)
				exp := tc.expectedBody.(*v1.GopherWordResponse)
				rm.On("TranslateWord", req.EnglishWord).Return(exp.GopherWord)
			}
			// call
			gh.TranslateWord("POST", "/word")(w, r)
			// assert
			rm.AssertExpectations(t)
			assertCode(t, w.Code, tc.expectedCode)
			assertBody(t, w.Body, tc.expectedBody)
		})
	}
}

func TestTranslateSentence(t *testing.T) {
	testCases := []struct {
		name         string
		method       string
		path         string
		body         interface{}
		expectedCode int
		expectedBody interface{}
	}{
		{"invalid method", "invalid", "/sentence", nil, http.StatusMethodNotAllowed, "405 method not allowed\n"},
		{"invalid path", "POST", "/invalid", nil, http.StatusNotFound, "404 page not found\n"},
		{"json decoding failed", "POST", "/sentence", `{"invalid":"a"}`, http.StatusBadRequest, &v1.ErrorResponse{Message: `json: unknown field "invalid"`}},
		{"ok", "POST", "/sentence", &v1.GopherSentenceRequest{EnglishSentence: "apple."}, http.StatusOK, &v1.GopherSentenceResponse{GopherSentence: "gapple."}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			w, r := createReq(tc.method, tc.path, tc.body)
			rm := mocks.NewGopherResourceMock()
			gh := &gopher{rs: rm}
			if tc.expectedCode == http.StatusOK {
				req := tc.body.(*v1.GopherSentenceRequest)
				exp := tc.expectedBody.(*v1.GopherSentenceResponse)
				rm.On("TranslateSentence", req.EnglishSentence).Return(exp.GopherSentence)
			}
			// call
			gh.TranslateSentence("POST", "/sentence")(w, r)
			// assert
			rm.AssertExpectations(t)
			assertCode(t, w.Code, tc.expectedCode)
			assertBody(t, w.Body, tc.expectedBody)
		})
	}
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
