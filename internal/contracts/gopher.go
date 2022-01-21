package contracts

// GopherWordRequest -
type GopherWordRequest struct {
	EnglishWord string `json:"english_word"`
}

// GopherWordResponse -
type GopherWordResponse struct {
	GopherWord string `json:"gopher_word"`
}

// GopherSentenceRequest -
type GopherSentenceRequest struct {
	EnglishSentence string `json:"english_sentence"`
}

// GopherSentenceResponse -
type GopherSentenceResponse struct {
	GopherSentence string `json:"gopher_sentence"`
}
