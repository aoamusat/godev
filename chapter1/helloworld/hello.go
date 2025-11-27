// gopl.io/ch1/helloworld
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")
	printArgsLength()
	printArgs()
}

func printArgs() {
	for i := 0; i < len(os.Args); i++ {
		fmt.Printf("Arg %d: %s\n", i, os.Args[i])
	}
}

func printArgsLength() {
	fmt.Printf("Total number of arguments (including program name): %d\n", len(os.Args))
}