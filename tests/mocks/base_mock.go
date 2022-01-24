package mocks

import (
	"fmt"
	"testing"
)

// CallMock -
type CallMock struct {
	expectation interface{}
	called      bool
}

// Return -
func (fm *CallMock) Return(v interface{}) {
	fm.expectation = v
}

// BaseMock -
type BaseMock struct {
	funcMocks map[string]map[interface{}]*CallMock
}

// NewBaseMock -
func NewBaseMock() *BaseMock {
	return &BaseMock{
		funcMocks: map[string]map[interface{}]*CallMock{},
	}
}

// On -
func (m *BaseMock) On(funcName string, arg interface{}) *CallMock {
	cm := &CallMock{}
	_, ok := m.funcMocks[funcName]
	if !ok {
		m.funcMocks[funcName] = map[interface{}]*CallMock{}
	}
	m.funcMocks[funcName][arg] = cm
	return cm
}

// MarkCalledAndReturn -
func (m *BaseMock) MarkCalledAndReturn(
	funcName string, arg interface{}, compFunc func(interface{}, interface{}) bool,
) interface{} {

	fm, ok := m.funcMocks[funcName]
	if !ok {
		panic(fmt.Sprintf("unexpected method called: %s", funcName))
	}

	for expArg, exp := range fm {
		if compFunc(expArg, arg) {
			exp.called = true
			return exp.expectation
		}
	}
	panic(fmt.Sprintf("unexpected method call: %s(%v)", funcName, arg))
}

// AssertExpectations -
func (m *BaseMock) AssertExpectations(t *testing.T) {
	for method, fm := range m.funcMocks {
		for arg, cm := range fm {
			if !cm.called {
				t.Errorf("expected call not received: %s(\"%v\")", method, arg)
			}
		}
	}
}
