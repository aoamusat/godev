package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !checkUrl(url) {
			fmt.Fprintf(os.Stderr, "fetch: invalid URL %s\n", url)
			os.Exit(1)
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// Print HTTP status code
		fmt.Fprintf(os.Stderr, "HTTP Status Code: %d %s\n", resp.StatusCode, resp.Status)

		// Copy response body to standard output for optimized reading
		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}

// checkUrl reports whether url begins with the HTTP or HTTPS scheme.
// It returns true if url starts with "http://" or "https://", and false otherwise.
// Note: this is a simple, case-sensitive prefix check and does not perform full
// URL parsing or validation of the host, path, query, or fragment.
func checkUrl(url string) bool {
	if strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
		return true
	}
	return false
}
