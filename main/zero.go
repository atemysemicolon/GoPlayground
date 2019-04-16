package main

import "fmt"

func main() {
	// Notice the default values
	// Not garbage like c, c++ but zero values
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}