// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/count", counterHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	// Log the request the stdout
	fmt.Fprintf(os.Stdout, "Received request from %s for %s. Current count: %d\n", r.RemoteAddr, r.URL.Path, count)
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counterHandler echoes the number of calls so far.
func counterHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	// Log the request the stdout
	fmt.Fprintf(os.Stdout, "Received request from %s for %s. Current count: %d\n", r.RemoteAddr, r.URL.Path, count)
	mu.Unlock()
}
