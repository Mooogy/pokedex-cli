package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string	
	}{
		{
			input: " hello world ",
			expected: []string{"hello", "world"},
		},
		{
			input: "HELLO",
			expected: []string{"hello"},
		},
		{
			input: "",
			expected: []string{},
		},
		{
			input: "   HeLlO	wOrLd",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths do not match: %v vs %v", actual, c.expected)
			continue
		}

		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}