package main

// problem Statement : Given a value x, find z such that z² is closest to it
// Find the nearest square and go ahead

import (
	"fmt"
	"math"
	"math/rand"
)

func Sqrt(x float64) float64 {

	//z := 1.0
	z := float64(rand.Int())
	fmt.Println("Initialized value of z : ", z)

	// Seems like the initial value of 1 brings us to minima much much faster than the previous

	for i :=0; i<100; i++ {
		// Lets update z
		diff := (z*z - x ) / (2*z)
		if math.Abs(diff) < 0.0001 {
			fmt.Println ("Breaking because update value is too small, z = ", z)
			break
		}
		z -= diff
		fmt.Println (i, ". The current value of z is : ", z)
	}
	return z
}

func main() {
	fmt.Println("Returned value: ", Sqrt(2))
}