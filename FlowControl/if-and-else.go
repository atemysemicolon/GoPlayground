package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		// Notice how v can be accessed here as well.
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main () {
	fmt.Println(
		pow(3,2,10),

		// This prints the print in the else part AND prints the return value
		pow(3,3,20),
		)
}