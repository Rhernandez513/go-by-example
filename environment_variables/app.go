package main

import "os"
import "strings"
import "fmt"

func main() {
	
	// To set a key/value pair, use os.Setenv()
	os.Setenv("FOO", "1")
	// To get a value for a key, use us.Getenv()
	// This will return an empty string if the key isn't present in the env
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	// Use os.Environ() to list all key/value pairs in the env
	// This returns a slice of string in the form KEY=value
	// You can string.Split() them to get the key and value.
	// Here we print all the keys
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}

// EOF
