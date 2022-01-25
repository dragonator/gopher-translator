package mocks

// TranslatorMock -
type TranslatorMock struct {
	*BaseMock
}

// NewTranslatorMock -
func NewTranslatorMock() *TranslatorMock {
	return &TranslatorMock{
		BaseMock: NewBaseMock(),
	}
}

// Translate -
func (tm *TranslatorMock) Translate(word string) (string, error) {
	v := tm.MarkCalledAndReturn("Translate", word, compareStrings).([]interface{})
	if v[1] == nil {
		return v[0].(string), nil
	}
	return v[0].(string), v[1].(error)
}
