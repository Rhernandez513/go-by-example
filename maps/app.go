package main

import "fmt"

// Maps are go's "Associative data type"
// like HASHMAPS or DICTIONARIES
func main() {
	// map_result = make(map[key-type] val-type)
	m := make(map[string]int)

	// Set key/value pairs using
	// name[key] = value

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]

	fmt.Println("v1: ", v1)

	fmt.Println("len: ", len(m))

	// the builtin delete() removes a key from the map
	delete(m, "k2")
	fmt.Println("map: ", m)

	// determine with the second return value if the key was missing
	// or if the value is 0 or ""
	_, prs := m["k2"]
	fmt.Println("prs: ", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)
}
