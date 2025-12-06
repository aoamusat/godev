package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase[T comparable] struct {
	name     string
	input    []T
	expected []T
}

func TestEliminateAdjacentDuplicates(t *testing.T) {
	tests := []TestCase[string]{
		{"basic duplicates", []string{"a", "a", "b", "c", "c", "c"}, []string{"a", "b", "c"}},
		{"no duplicates", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"all same", []string{"a", "a", "a"}, []string{"a"}},
		{"empty slice", []string{}, []string{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IEliminateAdjacentDuplicates(test.input)
			assert.Equal(t, result, test.expected)
		})
	}
}
