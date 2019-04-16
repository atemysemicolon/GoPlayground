package main

import "fmt"

// OUTSIDE A FUNCTION, SHORTHAND(:=) CANNOT BE USED!
// everything in this space should start from a keyword (var, func, ...)

func main() {
	var i, j int = 1,2
	k := 3 // shorthand. Inside function, so all is good
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
