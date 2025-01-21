package main

import (
	"testing"
)

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
			input:    "Good testing is 	a bit tiring thing",
			expected: []string{"good", "testing", "is", "a", "bit", "tiring", "thing"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Arrays are of different size, expected: %v, actual: %v", len(c.expected), len(actual))
			t.Fail()
			return
		}
		for i := 0; i < len(actual); i++ {
			if actual[i] != c.expected[i] {
				t.Errorf("Word at index %v is not right, expected: %s, actual: %s", i, c.expected[i], actual[i])
				t.Fail()
			}
		}
	}

}

