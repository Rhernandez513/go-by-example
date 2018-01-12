package main

//import "time"
import "fmt"
import (
	"math/rand"
	"time"
)

func main() {
	// get a random number n where 0 <= n < 100
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()

	// get a random float64 where 0.0 <= f 1.0
	fmt.Println(rand.Float64())

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64()*5)+5)
	fmt.Println()

	// rand gen is deterministic, so give it a seed that changes
	// use crypto/rand for secret random numbers
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)


	// call the resulting rand.Rand just like the functions on the rand package

	fmt.Print(r1.Intn(100), ",")
	fmt.Print(r1.Intn(100))
	fmt.Println()

	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Print(r2.Intn(100), ",")
	fmt.Print(r2.Intn(100))
	fmt.Println()

	// same seed = same seed of random nums
	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Print(r3.Intn(100), ",")
	fmt.Print(r3.Intn(100))
	fmt.Println()
}
