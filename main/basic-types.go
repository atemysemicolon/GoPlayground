package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe  	bool 		= false
	MaxInt 	uint64 		= 1<<64 -1
	z		complex128	= cmplx.Sqrt(-5+12i)
)

func main() {
	// Notice the way to display the type is so easy - %T
	// But this is only with Printf (not the same as Println)
	// Printf allows us to specify formatting strings in our message.
	// Println prints the string as is  - also printing %v and %T
	fmt.Printf("Type %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type %T Value: %v\n", z, z)
}