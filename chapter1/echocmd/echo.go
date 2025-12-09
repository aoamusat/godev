// Echo1 prints its command-line arguments.
// gopl.io/ch1/echo1
package echocmd

import (
	"os"
	"strings"
	"time"
)

func main() {
	var startTime = time.Now()
	echo1()
	elapsedTime := time.Since(startTime)
	println("Echo1 execution time:", elapsedTime)
}

func echo1() string {
	var output, separator string

	for i := 1; i < len(os.Args); i++ {
		output += separator + os.Args[i]
		separator = " "
	}

	return output
}

func echo3() string {
	return strings.Join(os.Args[1:], " ")
}
