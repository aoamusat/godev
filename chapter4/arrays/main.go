package arrays

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbols := [...]string{USD: "$", EUR: "€", GBP: "", RMB: "¥"}

	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	fmt.Println(RMB, symbols[RMB]) // "3 ¥"
	fmt.Printf("Length of symbols: %d", len(symbols))

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}
