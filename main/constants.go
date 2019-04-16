package main

import "fmt"

// Just the const keyword bro
const Pi = 3.14

// can also specify the type if you want, but can't redefine a constant ;)
// const Pi float64 = 3.14

func main()  {
	const World = "España"
	fmt.Println("Hello", World)
	fmt.Println("Heppy", Pi, "Day")

	// const Truth := true // fails. No shorthand notation for constants
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
