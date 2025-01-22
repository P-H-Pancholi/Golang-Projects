package main

import (
	"testing"
)

func TestCleanInput(T *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   hello  world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Naruto     Sasuke     guy",
			expected: []string{"naruto", "sasuke", "guy"},
		},
		{
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, c := range cases {
		actual := CleanInput(c.input)
		if len(actual) != len(c.expected) {
			T.Errorf("Expected length of the array is %d, actual is %d", len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				T.Errorf("Expected string %s, Found string %s", expectedWord, word)
			}
		}
	}

}
