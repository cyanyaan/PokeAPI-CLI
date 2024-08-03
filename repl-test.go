package main

import (
	"fmt"
	"testing"
)

func TestingCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		actual := cleanInput(cs.input)
		if len(actual) != len(cs.expected) {
			t.Errorf("The lengths are not equal: %v vs %v", len(actual), len(cs.expected))
			continue
		}
		for i := range actual {
			actualWord := actual[i]
			expectedWord := cs.expected[i]
			fmt.Printf("Actual Word: %v, Expected word: %v", actualWord, expectedWord)
			if actualWord != expectedWord {
				t.Errorf("%v does not equal %v", actualWord, expectedWord)
			}
		}
	}

}
