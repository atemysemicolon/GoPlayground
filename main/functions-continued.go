package main

import "fmt"

// add gets redeclared - notice the warning
// Notice the shorthand parameter decleration
func add(x, y int) int {
	return x+y
}

func main() {
	fmt.Println(add(42, 13 ))
}