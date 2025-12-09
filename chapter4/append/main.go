package append

import "fmt"

func main() {
	// x1 := []int{1, 2, 3}
	// y1 := x1[:]
	// y1 = append(y1, 23)

	// fmt.Printf("Length of x1: %d\nCapacity of x1: %d\nx1: %v\n", len(x1), cap(x1), x1)
	// fmt.Printf("Length of y1: %d\nCapacity of y1: %d\ny1: %v\n", len(y1), cap(y1), y1)
	var x, y []int
	for i := range 10 {
		y = appendT(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

// appendT appends a single element y to a slice x and returns the resulting slice.
// If the underlying array has sufficient capacity, the slice is extended in place.
// Otherwise, a new array is allocated with capacity doubled to ensure amortized linear complexity.
// The function is generic and works with any type T.
func appendT[T any](x []T, y ...T) []T {
	var z []T
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]T, zlen, zcap)
		copy(z[len(x):], y) // a built-in function; see text
	}
	return z
}
