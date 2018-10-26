package main

import "fmt"
import "os"

func main() {
	
	// defer will not be run when using os.Exit, so this fmt.Println will never 
	// be called.
	defer fmt.Println("!")

	// Note that unlike e.g. C, Go does not use an integer return value from 
	// main to indicate exit status. If you’d like to exit with a non-zero 
	// status you should use os.Exit.
	os.Exit(3)
}

// EOF
