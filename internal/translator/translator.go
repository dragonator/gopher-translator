package translator

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/dragonator/gopher-translator/internal/service/svc"
)

// Specification -
type Specification struct {
	Rules        []*Rule  `json:"rules"`
	InvalidChars []string `json:"invalid_characters"`
}

// Rule -
type Rule struct {
	MatchPattern   string `json:"match_pattern"`
	ReplacePattern string `json:"replace_pattern"`
}

// Translator -
type Translator interface {
	Translate(word string) (string, error)
}

type compiledRule struct {
	re             *regexp.Regexp
	replacePattern string
}

type translator struct {
	rules        []*compiledRule
	invalidChars []string
}

// New -
func New(spec *Specification) Translator {
	t := &translator{
		invalidChars: spec.InvalidChars,
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
func (t *translator) Translate(word string) (string, error) {
	for _, c := range t.invalidChars {
		if strings.Contains(word, c) {
			return "", fmt.Errorf("%w: word contains invalid character: %s", svc.ErrInvalidInput, c)
		}
	}

	for _, rule := range t.rules {
		if rule.re.MatchString(word) {
			return rule.re.ReplaceAllString(word, rule.replacePattern), nil
		}
	}
	return word, nil
}
