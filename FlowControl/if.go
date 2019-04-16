package main

import (
	"fmt"
	"math"
)

func sqrt (x float64) string {

	// Notice lack of parenthesis but forced braces
	if x<0 {
		// Notice the inline concatenation. TODO: Here it is casted implicitly???
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main(){
	fmt.Println(sqrt(2), sqrt(-4))
}