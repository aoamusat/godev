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

func TestIEliminateAdjacentDuplicates(t *testing.T) {
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

func TestEliminateAdjacentDuplicates(t *testing.T) {
	tests := []TestCase[string]{
		{"basic duplicates", []string{"a", "a", "b", "c", "c", "c"}, []string{"a", "b", "c"}},
		{"no duplicates", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
		{"all same", []string{"a", "a", "a"}, []string{"a"}},
		{"empty slice", []string{}, []string{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := eliminateAdjacentDuplicates(test.input)
			assert.Equal(t, result, test.expected)
		})
	}
}
func TestReverse(t *testing.T) {
	t.Run("reverse strings", func(t *testing.T) {
		input := []string{"a", "b", "c", "d"}
		expected := []string{"d", "c", "b", "a"}
		reverse(input)
		assert.Equal(t, expected, input)
	})

	t.Run("reverse integers", func(t *testing.T) {
		input := []int{1, 2, 3, 4, 5}
		expected := []int{5, 4, 3, 2, 1}
		reverse(input)
		assert.Equal(t, expected, input)
	})

	t.Run("single element", func(t *testing.T) {
		input := []string{"a"}
		expected := []string{"a"}
		reverse(input)
		assert.Equal(t, expected, input)
	})

	t.Run("empty slice", func(t *testing.T) {
		input := []string{}
		expected := []string{}
		reverse(input)
		assert.Equal(t, expected, input)
	})

	t.Run("two elements", func(t *testing.T) {
		input := []int{1, 2}
		expected := []int{2, 1}
		reverse(input)
		assert.Equal(t, expected, input)
	})
}
