package main

import "fmt"

// Looks similar to a struct in C
type person struct {
	name string
	age int
}

func main() {
	fmt.Println(person{"Bob", 20})

	fmt.Println(person{"Alice", 30})

	fmt.Println(person{name: "Fred"})

	fmt.Println(&person{name:"Ann", age: 40})

	s := person { name: "Sean", age: 50 }
	fmt.Println(s.name)

	sp := &s
	// Auto dereference of pointers w/ the same dot syntax
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
