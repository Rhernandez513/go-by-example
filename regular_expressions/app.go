package main

import "bytes"
import "fmt"
import "regexp"

func main () {

	// Here we test directly whether a pattern matches a string
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// for other operations we will need to compile an optimized version
	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	// here we find the match for the regexp in the provided val
	fmt.Println(r.FindString("peach punch"))
	// Here we return the first and last indices of the matched val
	// seems to be more useful to me
	fmt.Println(r.FindStringIndex("peach punch"))

	// this does p([a-z]+)ch and also p([a-z])+
	fmt.Println(r.FindStringSubmatch("peach punch"))
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// the All variant matches on all matches in the input not just the fisrt
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	// Here we limit the number of matches
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// byte arrays work and we can drop string from the function name
	fmt.Println(r.Match([]byte("peach")))

	// we can use must compile for constants

	r = regexp.MustCompile("p([a-z]+)ch")
	// printing the regex representation
	fmt.Println(r)

	// we can also replace subsets of strings with other values
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	// allows you to transform matches text with a given function
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	// convert the byte array into a string represntation of course
	fmt.Println(string(out))
}
