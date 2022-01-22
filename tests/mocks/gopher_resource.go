package mocks

// GopherResourceMock -
type GopherResourceMock struct {
	*BaseMock
}

// NewGopherResourceMock -
func NewGopherResourceMock() *GopherResourceMock {
	return &GopherResourceMock{
		BaseMock: NewBaseMock(),
	}
}

// TranslateWord -
func (tm *GopherResourceMock) TranslateWord(word string) string {
	v, ok := tm.MarkCalledAndReturn("TranslateWord", word, compareStrings).(string)
	if !ok {
		panic("unexpected return value type")
	}
	return v
}

// TranslateSentence -
func (tm *GopherResourceMock) TranslateSentence(word string) string {
	v, ok := tm.MarkCalledAndReturn("TranslateSentence", word, compareStrings).(string)
	if !ok {
		panic("unexpected return value type")
	}
	return v
}

// History -
func (tm *GopherResourceMock) History() []map[string]string {
	v, ok := tm.MarkCalledAndReturn("History", nil, compareNils).([]map[string]string)
	if !ok {
		panic("unexpected return value type")
	}
	return v
}
