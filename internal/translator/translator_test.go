package translator

import (
	"encoding/json"
	"regexp"
	"strings"
	"testing"
)

func TestSpecification(t *testing.T) {

	testCases := []struct {
		name        string
		input       string
		expectedErr bool
		expectedOut [][2]string
	}{
		{"with pairs", `{"normalizer":[["a", "A"],["p", "P"]]}`, false, [][2]string{{"a", "A"}, {"p", "P"}}},
		{"non-pairs", `{"normalizer":[["a", "A", "X],["p", "P"]]}`, true, [][2]string{}},
		{"with single value", `{"normalizer":[["a"]]}`, false, [][2]string{{"a", ""}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			spec := &Specification{}
			err := json.Unmarshal([]byte(tc.input), spec)
			if tc.expectedErr {
				if err == nil {
					t.Errorf("unexpected unmarshal success: %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("unexpected unmarshal failure: %v", err)
				}
			}
			for i, pair := range spec.Normalizer {
				for j, e := range pair {
					if e != tc.expectedOut[i][j] {
						t.Errorf("unexpected pair: %s (expected: %s)", pair, tc.expectedOut[i])
						break
					}
				}
			}
		})
	}
}

func TestNew(t *testing.T) {
	t.Run("test created replacer", func(t *testing.T) {
		testCases := []struct {
			name     string
			input    string
			expected string
			spec     [][2]string
			flatten  []string
		}{
			{"with single pair", "appa", "aPPa", [][2]string{{"p", "P"}}, []string{"p", "P"}},
			{"with multiple pairs", "appa", "APPA", [][2]string{{"p", "P"}, {"a", "A"}}, []string{"p", "P", "a", "A"}},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				// setup
				spec := &Specification{
					Normalizer: tc.spec,
				}
				// call
				tr := New(spec).(*translator)
				// assert
				res := tr.replacer.Replace(tc.input)
				if res != tc.expected {
					t.Errorf("unexpected replacement result: %s (expected: %s)", res, tc.expected)
				}
			})
		}
	})

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
		name     string
		input    string
		expected string
		rules    [][]string
	}{
		{"without rules", "something", "something", [][]string{}},
		{"matches no rules", "something", "something", [][]string{{`^([aeiou].*)$`, `g$1`}}},
		{"applies matched rule", "apple", "gapple", [][]string{{`^([aeiou].*)$`, `g$1`}}},
		{"applies first matched rule only", "a", "aa", [][]string{{`^(a.*)$`, `a$1`}, {`^aa(.*)$`, `bb$1`}}},
		{"normalizes word before translation", "a’pp'le", "gapple", [][]string{{`^([aeiou].*)$`, `g$1`}}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// setup
			tr := &translator{
				replacer: strings.NewReplacer("'", "", "’", ""),
			}
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
