package main

import (
	"fmt"

	"a4m.dev/godev/chapter4/treesort"
)

func main() {
	values := []int{5, 3, 8, 1, 4, 7, 6, 2}
	treesort.Sort(values)
	fmt.Printf("Sorted values:")
	for _, v := range values {
		fmt.Printf(" %d", v)
	}
}
