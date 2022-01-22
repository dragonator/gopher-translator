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
func (tm *TranslatorMock) Translate(word string) string {
	v, ok := tm.MarkCalledAndReturn("Translate", word, compareStrings).(string)
	if !ok {
		panic("unexpected return value type")
	}
	return v
}
