package main

import "fmt"

func main() {
	v := 42 // The answer of everything
	fmt.Printf("v is of type %T\n", v)

	// v := false // No shorthand notation for existing var - can't redefine an existing variable like this
	// v = false // can't change datatype within a container, unlike Python!
	k := false
	fmt.Printf("k is of type %T\n", k)
}