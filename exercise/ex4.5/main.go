package ex45

// eliminateAdjacentDuplicates removes consecutive duplicate strings from the input slice.
// It takes a slice of strings and returns a new slice with adjacent duplicates eliminated.
// The original slice is not modified.
//
// Example:
//
//	input := []string{"a", "a", "b", "b", "b", "c", "a"}
//	result := eliminateAdjacentDuplicates(input)
//	// result: []string{"a", "b", "c", "a"}
func eliminateAdjacentDuplicates(strings []string) []string {
	if len(strings) == 0 {
		return strings
	}

	var result []string
	previous := strings[0]
	result = append(result, previous)

	for _, current := range strings[1:] {
		if current != previous {
			result = append(result, current)
			previous = current
		}
	}

	return result
}

// reverse reverses the elements of slice s in place.
// It swaps elements from both ends moving towards the center,
// effectively reversing the order of all elements in the slice.
// T is a type parameter that can be any type.
func reverse[T any](s []T) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func IEliminateAdjacentDuplicates[T comparable](elements []T) []T {
	if len(elements) == 0 {
		return elements
	}

	var result []T
	previous := elements[0]
	result = append(result, previous)
	for _, v := range elements[1:] {
		if v != previous {
			result = append(result, v)
			previous = v
		}
	}

	return result
}
