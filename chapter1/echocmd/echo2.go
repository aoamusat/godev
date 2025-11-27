// gopl.io/ch1/echo2
// Echo2 prints its command-line arguments.

package main

import (
	"fmt"
	"os"
)

func echo2() {
	separator, output := "", ""

	for _, arg := range os.Args[1:] {
		output += separator + arg
		separator = " "
	}

	fmt.Println(output)
}
