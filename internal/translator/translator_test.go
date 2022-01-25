package translator

import (
	"regexp"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("test compiled rules", func(t *testing.T) {
		// setup
		spec := &Specification{
			Rules: []*Rule{
				{MatchPattern: "aaa", ReplacePattern: "AAA"},
				{MatchPattern: "bbb", ReplacePattern: "BBB"},
			},
		}

		// call
		tr, ok := New(spec).(*translator)
		// assert
		if !ok {
			t.Errorf("unexpected underlying type")
		}
		for i, compRule := range tr.rules {
			// match pattern
			expected := spec.Rules[i].MatchPattern
			actual := compRule.re.String()
			if actual != expected {
				t.Errorf("unexpected match pattern for rule at index %d: %s (expected: %s))", i, actual, expected)
			}
			// replace pattern
			expected = spec.Rules[i].ReplacePattern
			actual = compRule.replacePattern
			if actual != expected {
				t.Errorf("unexpected replace pattern for rule at index %d: %s (expected: %s))", i, actual, expected)
			}
		}
	})
}

func TestTranslate(t *testing.T) {
	testCases := []struct {
		name       string
		rules      [][]string
		input      string
		expected   string
		shouldFail bool
	}{
		{"without rules", [][]string{}, "something", "something", false},
		{"matches no rules", [][]string{{`^([aeiou].*)$`, `g$1`}}, "something", "something", false},
		{"applies matched rule", [][]string{{`^([aeiou].*)$`, `g$1`}}, "apple", "gapple", false},
		{"applies first matched rule only", [][]string{{`^(a.*)$`, `a$1`}, {`^aa(.*)$`, `bb$1`}}, "a", "aa", false},
		{"with invalid characters", [][]string{{`^([aeiou].*)$`, `g$1`}}, "a'pple", "", true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			tr := &translator{
				invalidChars: []string{"'"},
			}
			for _, rule := range tc.rules {
				tr.rules = append(tr.rules, &compiledRule{
					re:             regexp.MustCompile(rule[0]),
					replacePattern: rule[1],
				})
			}
			// call
			result, err := tr.Translate(tc.input)
			// assert
			if tc.shouldFail && err == nil {
				t.Errorf("expected error: got nil")
			} else if !tc.shouldFail && err != nil {
				t.Errorf("unexpected error: %s)", err)
			}
			if result != tc.expected {
				t.Errorf("unexpected translation: %s (expected: %s))", result, tc.expected)
			}
		})
	}
}
