package nonempty

import "fmt"

func main() {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", NonEmpty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
}

// NonEmpty returns a new slice containing only the non-empty strings from the input slice.
// It filters out any empty string values and returns a slice with the remaining elements.
func NonEmpty(strings []string) []string {
	idx := 0
	for _, val := range strings {
		if val != "" {
			strings[idx] = val
			idx++
		}
	}
	return strings[:idx]
}

// reverse reverses the elements of the slice in place.
// It takes a pointer to a slice of any type T and swaps elements
// from both ends moving towards the center until the middle is reached.
// The slice is modified directly; no new slice is created.
func reverse[T any](arg *[]T) {
	for i, j := 0, len(*arg)-1; i < j; i, j = i+1, j-1 {
		(*arg)[i], (*arg)[j] = (*arg)[j], (*arg)[i]
	}
}
