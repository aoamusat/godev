package main

import "fmt"

func main() {
	ages := []int{23, 45, 12, 67, 34}
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve"}
	fmt.Println(ages)
	fmt.Println(names)
	reverseSlice(ages)
	reverseSlice(names)
	fmt.Println(ages)
	fmt.Println(names)
}

// reverseSlice reverses the order of elements in the provided slice in place.
// It swaps elements from the ends toward the center, using O(1) additional
// memory and running in O(n) time. The function is generic over type parameter
// T (constrained to comparable), accepts nil or empty slices, and does not
// return a value.
func reverseSlice[T comparable](arg []T) {
	for i, j := 0, len(arg)-1; i < j; i, j = i+1, j-1 {
		arg[i], arg[j] = arg[j], arg[i]
	}
}
