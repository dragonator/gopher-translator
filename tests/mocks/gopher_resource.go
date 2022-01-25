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
func (tm *GopherResourceMock) TranslateWord(word string) (string, error) {
	v := tm.MarkCalledAndReturn("TranslateWord", word, compareStrings).([]interface{})
	if v[1] == nil {
		return v[0].(string), nil
	}
	return v[0].(string), v[1].(error)
}

// TranslateSentence -
func (tm *GopherResourceMock) TranslateSentence(word string) (string, error) {
	v := tm.MarkCalledAndReturn("TranslateSentence", word, compareStrings).([]interface{})
	if v[1] == nil {
		return v[0].(string), nil
	}
	return v[0].(string), v[1].(error)
}

// History -
func (tm *GopherResourceMock) History() []map[string]string {
	return tm.MarkCalledAndReturn("History", nil, compareNils).([]interface{})[0].([]map[string]string)
}
