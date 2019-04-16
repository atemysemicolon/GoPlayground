package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {

	// Assign and condition in the same loop

	// Commented an alternate implementation. This assign+condition helps us not recompute the power later on.
	// python caches the function output in the scenario that it is called again, But perhaps go does not

	//if math.Pow(x,n) < lim { //alternate
	if v := math.Pow(x,n); v < lim {
		return v
		// return math.Pow(x,n) //alternate
	}
	return lim
}

func main() {
	// Notice the comma (,) after the second call. In go, it seems like its required, if I want the command to understand the next line

	fmt.Println(
		pow(3,2, 10),
		pow(3,3, 20),
		)

	// Alternatively can be written as
	//fmt.Println(
	//	pow(3,2, 10),
	//	pow(3,3, 20))
}