package main

import (
	"fmt"
	"math"
)


func main() {
	var x,y int = 3,4

	// compile fails if you don't cast. All casting is explicit!!!!
	var f float64 = math.Sqrt(float64(x*x + y*y))

	var z uint = uint(f)
	fmt.Println(x, y, z)
}