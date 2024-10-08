package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "Hello World",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "just a new a bit larger test",
			expected: []string{
				"just",
				"a",
				"new",
				"a",
				"bit",
				"larger",
				"test",
			},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "SiNgLe",
			expected: []string{"single"},
		},
	}

	for _, c := range cases {
		r := cleanInput(c.input)
		if len(r) != len(c.expected) {
			t.Errorf("Not the expected length. \nExpected: %v\nActual: %v",
				len(c.expected),
				len(r),
			)
			continue
		}
		for i := range r {
			actualWord := r[i]
			expectedWord := c.expected[i]
			if actualWord != expectedWord {
				t.Errorf("\nExpected: %v\nReceived: %v", expectedWord, actualWord)
			}
		}
	}
}
