package translator

import (
	"regexp"
	"strings"
)

// Specification -
type Specification struct {
	Rules      []*Rule     `json:"rules"`
	Normalizer [][2]string `json:"normalizer"`
}

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
	rules    []*compiledRule
	replacer *strings.Replacer
}

// New -
func New(spec *Specification) Translator {
	var flatten []string
	for _, pair := range spec.Normalizer {
		flatten = append(flatten, pair[:]...)
	}

	t := &translator{
		replacer: strings.NewReplacer(flatten...),
	}

	for _, rule := range spec.Rules {
		t.rules = append(t.rules, &compiledRule{
			re:             regexp.MustCompile(rule.MatchPattern),
			replacePattern: rule.ReplacePattern,
		})
	}

	return t
}

// Translate -
func (t *translator) Translate(word string) string {
	normalized := t.replacer.Replace(word)
	for _, rule := range t.rules {
		if rule.re.MatchString(normalized) {
			return rule.re.ReplaceAllString(normalized, rule.replacePattern)
		}
	}
	return normalized
}
