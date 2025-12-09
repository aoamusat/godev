package dup1

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counter := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counter[input.Text()]++
	}

	for line, n := range counter {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
