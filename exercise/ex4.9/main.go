// Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text
// file. Call input.Split(bufio.ScanWords) before the first call to Scan to break the input into
// words instead of lines.
package main

import (
	"bufio"
	"os"
)

func main() {
}

// wordfreq
func wordfreq(file *os.File) {
	fileContent := bufio.NewScanner(file)
	fileContent.Split(bufio.ScanWords)

	wordCount := make(map[string]int)
	for fileContent.Scan() {
		word := fileContent.Text()
		wordCount[word]++
	}

	for word, count := range wordCount {
		println(word, count)
	}
}
