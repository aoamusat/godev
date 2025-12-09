package dup2

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counters := make(map[string]int)
	fileNames := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counters, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counters, fileNames)
			f.Close()
		}
	}

	// Print duplicates
	for line, n := range counters {
		if n > 1 {
			fmt.Printf("%d\t%s\t%v\n", n, line, fileNames[line])
		}
	}
}

// countLines scans the provided file from its current offset line by line and
// increments the entry in counter for each line encountered.
//
// The counter map is mutated in place; each distinct line (including any
// leading/trailing whitespace) is used as a key and its count is incremented.
// This function uses bufio.Scanner to read lines and does not return errors:
// any scanning errors are ignored. It is not safe for concurrent use — callers
// must synchronize access to the counter if it may be accessed from multiple
// goroutines.
func countLines(file *os.File, counter map[string]int, fileNames map[string][]string) {
	input := bufio.NewScanner(file)
	fileName := getFileName(file)
	for input.Scan() {
		line := input.Text()
		counter[line]++
		if !contains(fileNames[line], fileName) {
			fileNames[line] = append(fileNames[line], fileName)
		}
	}
}

// getFileName returns the base name of the provided file by calling file.Stat()
// and returning the FileInfo.Name(). If Stat() fails, it returns the literal
// "unknown". Passing a nil *os.File will cause a runtime panic.
func getFileName(file *os.File) string {
	info, err := file.Stat()
	if err != nil {
		return "unknown"
	}
	return info.Name()
}

// contains checks if a string slice contains a specific string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
