package ex46

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestSquashUnicodeSpace(t *testing.T) {
	tests := []struct {
		name     string
		input    []byte
		expected []byte
	}{
		{"consecutive spaces", []byte("a  b  c"), []byte("a b c")},
		{"tabs and spaces", []byte("a \t b"), []byte("a b")},
		{"no spaces", []byte("abc"), []byte("abc")},
		{"leading spaces", []byte("  abc"), []byte(" abc")},
		{"trailing spaces", []byte("abc  "), []byte("abc ")},
		{"all spaces", []byte("   "), []byte(" ")},
		{"empty slice", []byte(""), []byte("")},
		{"single character", []byte("a"), []byte("a")},
		{"newlines and spaces", []byte("a\n\nb"), []byte("a b")},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := squashUnicodeSpace(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
