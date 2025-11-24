package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counters := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counters)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counters)
			f.Close()
		}
	}

	// Print duplicates
	for line, n := range counters {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(file *os.File, counter map[string]int) {
	input := bufio.NewScanner(file)
	for input.Scan() {
		line := input.Text()
		counter[line]++
	}
}
