package ex46

import "unicode"

// squashUnicodeSpace takes a byte slice and returns a new slice with consecutive
// unicode whitespace characters collapsed into single spaces. Non-space characters
// are preserved as-is. The function modifies the input slice in-place and returns
// a resliced view of it containing only the processed bytes.
func squashUnicodeSpace(b []byte) []byte {
	n := 0
	spaceFound := false

	for _, byteValue := range b {
		if unicode.IsSpace(rune(byteValue)) {
			if !spaceFound {
				b[n] = ' '
				n++
				spaceFound = true
			}
		} else {
			b[n] = byteValue
			n++
			spaceFound = false
		}
	}

	return b[:n]
}
