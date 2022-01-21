package translator

import (
	"regexp"
)

// Rule -
type Rule struct {
	MatchPattern   string `json:"match_pattern"`
	ReplacePattern string `json:"replace_pattern"`
}

// Translator -
type Translator interface {
	Translate(word string) string
}

type compiledRule struct {
	re             *regexp.Regexp
	replacePattern string
}

type translator struct {
	rules []*compiledRule
}

// New -
func New(rules []*Rule) Translator {
	t := &translator{}
	for _, rule := range rules {
		t.rules = append(t.rules, &compiledRule{
			re:             regexp.MustCompile(rule.MatchPattern),
			replacePattern: rule.ReplacePattern,
		})
	}

	return t
}

// Translate -
func (t *translator) Translate(word string) string {
	for _, rule := range t.rules {
		if rule.re.MatchString(word) {
			return rule.re.ReplaceAllString(word, rule.replacePattern)
		}
	}
	return word
}
