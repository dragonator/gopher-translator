package svc_test

import (
	"testing"

	"github.com/dragonator/gopher-translator/internal/service/svc"
)

func TestError(t *testing.T) {
	input := "some message"
	e := svc.Error{
		Message: input,
	}
	m := e.Error()
	if m != input {
		t.Errorf("unexpected error message: %s (expected: %s)", m, input)
	}
}
