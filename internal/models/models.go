package models

// Rule -
type Rule struct {
	MatchPattern   string `json:"match_pattern"`
	ReplacePattern string `json:"replace_pattern"`
}

// Record -
type Record struct {
	EnglishWord string `json:"english_word"`
	GopherWord  string `json:"gopher_word"`
}
