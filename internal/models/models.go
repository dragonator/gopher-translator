package models

// Rule -
type Rule struct {
	MatchPattern   string `json:"match_pattern"`
	ReplacePattern string `json:"replace_pattern"`
}
