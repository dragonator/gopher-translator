package service_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dragonator/gopher-translator/internal/handlers"
	"github.com/dragonator/gopher-translator/internal/service"
	"github.com/dragonator/gopher-translator/tests/mocks"
)

func TestRouter(t *testing.T) {
	// setup
	rm := mocks.NewGopherResourceMock()
	rm.On("History", nil).Return([]map[string]string{})
	gh := handlers.NewGopher(rm)

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/history", nil)
	// call
	router := service.NewRouter(gh)
	router.ServeHTTP(w, r)
	// assert
	if w.Code != http.StatusOK {
		t.Errorf("unexpected code: %d (expected: %d)", w.Code, http.StatusOK)
	}
}
