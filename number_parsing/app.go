package main

import "strconv"
import "fmt"

func main() {
	// this 64 shows how many bits of precision to parse
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	//for parsein 0 means infer the base from the string
	// 64 means that the result should fit in 64 bits
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// parseint will recognize hex formatted numbers
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// unsigned
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// convenience func for basic base 10 parsing
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	_, e := strconv.Atoi("wat")
	fmt.Println(e)
}

