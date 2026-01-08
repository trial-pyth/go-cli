package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		// check slice length
		if len(actual) != len(c.expected) {
			t.Errorf(
				"cleanInput(%q) length = %d, expected %d",
				c.input,
				len(actual),
				len(c.expected),
			)
			continue
		}

		// check each word
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf(
					"cleanInput(%q)[%d] = %q, expected %q",
					c.input,
					i,
					word,
					expectedWord,
				)
			}
		}
	}
}
