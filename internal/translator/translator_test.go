package translator

import (
	"regexp"
	"testing"
)

func TestNew(t *testing.T) {
	// setup
	rules := []*Rule{
		{MatchPattern: "aaa", ReplacePattern: "AAA"},
		{MatchPattern: "bbb", ReplacePattern: "BBB"},
	}
	// call
	tr, ok := New(rules).(*translator)
	// assert
	if !ok {
		t.Errorf("unexpected underlying type")
	}
	for i, compRule := range tr.rules {
		// match pattern
		expected := rules[i].MatchPattern
		actual := compRule.re.String()
		if actual != expected {
			t.Errorf("unexpected match pattern for rule at index %d: %s (expected: %s))", i, actual, expected)
		}
		// replace pattern
		expected = rules[i].ReplacePattern
		actual = compRule.replacePattern
		if actual != expected {
			t.Errorf("unexpected replace pattern for rule at index %d: %s (expected: %s))", i, actual, expected)
		}
	}
}

func TestTranslate(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected string
		rules    [][]string
	}{
		{"without rules", "something", "something", [][]string{}},
		{"matches no rules", "something", "something", [][]string{{`^([aeiou].*)$`, `g$1`}}},
		{"applies matched rule", "apple", "gapple", [][]string{{`^([aeiou].*)$`, `g$1`}}},
		{"applies first matched rule only", "a", "aa", [][]string{{`^(a.*)$`, `a$1`}, {`^aa(.*)$`, `bb$1`}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			tr := &translator{}
			for _, rule := range tc.rules {
				tr.rules = append(tr.rules, &compiledRule{
					re:             regexp.MustCompile(rule[0]),
					replacePattern: rule[1],
				})
			}
			// call
			result := tr.Translate(tc.input)
			// assert
			if result != tc.expected {
				t.Errorf("unexpected translation: %s (expected: %s))", result, tc.expected)
			}
		})
	}
}
