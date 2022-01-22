package service

import (
	"fmt"
	"strings"
	"testing"
)

func TestService(t *testing.T) {
	t.Run("error decoding spec", func(t *testing.T) {
		// setup
		b := &Bootstrap{
			Spec: strings.NewReader(`invalid spec`),
		}
		// call
		s, err := New(b)
		// assert
		if err == nil {
			t.Error("expecting decoding to fail")
		}
		if s != nil {
			t.Error("expecting service to be nil")
		}
	})

	t.Run("ok", func(t *testing.T) {
		// setup
		b := &Bootstrap{
			Port: "8080",
			Spec: strings.NewReader(`{}`),
		}
		// call
		s, err := New(b)
		// assert
		if err != nil {
			t.Errorf("unexpected failure: %v", err)
		}
		if s == nil {
			t.Error("unexpected nil value for service")
		}
		expAddr := fmt.Sprintf(":%s", b.Port)
		if s.server.Addr != expAddr {
			t.Errorf("wrong value set for server addres: %s (expected: %s)", s.server.Addr, expAddr)
		}
	})
}
