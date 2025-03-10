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
			input:    "   ",
			expected: []string{}, // Input with only spaces should return an empty slice
		},
		{
			input:    "GoLang is AWESOME",
			expected: []string{"golang", "is", "awesome"}, // Mixed casing should be converted to lowercase
		},
		{
			input:    "   trim   this   ",
			expected: []string{"trim", "this"}, // Leading and trailing spaces should be removed
		},
		{
			input:    "special!@#characters",
			expected: []string{"special!@#characters"}, // Special characters should remain intact
		},
		{
			input:    "one\ttwo\nthree",
			expected: []string{"one", "two", "three"}, // Tabs and newlines should be treated as whitespace
		},
		{
			input:    "  multiple   spaces   between words  ",
			expected: []string{"multiple", "spaces", "between", "words"}, // Multiple spaces should be treated as a single delimiter
		},
		{
			input:    "",
			expected: []string{}, // Empty input should return an empty slice
		},
		{
			input:    "123 456 789",
			expected: []string{"123", "456", "789"}, // Numbers should be treated as words
		},
		{
			input:    "  MiXeD CaSe WoRdS  ",
			expected: []string{"mixed", "case", "words"}, // Mixed case should be converted to lowercase
		},
		{
			input:    "word-with-hyphen",
			expected: []string{"word-with-hyphen"}, // Hyphenated words should remain intact
		},
		{
			input:    "punctuation, should! stay?",
			expected: []string{"punctuation,", "should!", "stay?"}, // Punctuation should remain part of the words
		},
		{
			input:    "  leading, trailing, and  in-between spaces  ",
			expected: []string{"leading,", "trailing,", "and", "in-between", "spaces"}, // Handles a mix of spaces and punctuation
		},
	}

	for _, c := range cases {
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("len error")
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("value error")
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
		}
}
}