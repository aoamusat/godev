// Fetchall fetches URLs in parallel and reports their times and sizes.

package fetchall

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	ch := make(chan string)
	urls := os.Args[1:]
	totalSec := 0

	for _, url := range urls {
		go fetch(url, ch, &totalSec)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(startTime).Seconds())
	fmt.Printf("Total time taken: %ds\n", totalSec)
}

// fetch retrieves the content from the specified URL and sends a formatted result message to the provided channel.
// It measures the time taken to fetch and read the response body, and reports the elapsed time in seconds
// and the number of bytes read. If an error occurs during fetching or reading the response body,
// an error message is sent to the channel instead. The response body is always closed to prevent resource leaks.
//
// Parameters:
//   - url: the URL string to fetch
//   - ch: a send-only channel that receives the result message (timing, byte count, and URL) or error message
func fetch(url string, ch chan<- string, totalSec *int) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("while fetching %s: %v", url, err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	*totalSec += int(secs)
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
