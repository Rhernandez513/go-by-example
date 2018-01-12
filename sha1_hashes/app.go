package main

import "crypto/sha1"
import "fmt"

func main() {
	s := "sha1 this string"

	h := sha1.New()

	h.Write([]byte(s))

	// this gets the finalized hash result as a byte slice,
	// the argument to sum can be used to append to an existing one
	// but it usually is not needed
	byte_slice := h.Sum(nil)

	fmt.Println(s)
	// converts a hash result into a hex strings
	fmt.Printf("%x\n", byte_slice)

	// MD5 works the same way
	// import crypto/md5 and use md5.New()
}
