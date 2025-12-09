package nonempty

import (
	"testing"
)

func TestReverseStrings(t *testing.T) {
	testcases := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "reverse string slice",
			input:    []string{"a", "b", "c"},
			expected: []string{"c", "b", "a"},
		},
		{
			name:     "reverse single element",
			input:    []string{"x"},
			expected: []string{"x"},
		},
		{
			name:     "reverse empty slice",
			input:    []string{},
			expected: []string{},
		},
		{
			name:     "reverse two elements",
			input:    []string{"one", "two"},
			expected: []string{"two", "one"},
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			reverse(&testcase.input)
			for i, v := range testcase.input {
				if v != testcase.expected[i] {
					t.Errorf("got %v, want %v", testcase.input, testcase.expected)
				}
			}
		})
	}
}

func TestReverseInts(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}
	reverse(&input)
	for i, v := range input {
		if v != expected[i] {
			t.Errorf("got %v, want %v", input, expected)
		}
	}
}
