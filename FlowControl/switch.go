package main

import (
	"fmt"
	"runtime"
)


func main(){
	fmt.Print("Go runs on ")

	// Notes on switch
	//1. The same kind of assignment + compare as in if
	//2. Using the module runtime.GOOS to get current os name
	//3. Notice also the cases after the correct one don't execute. No break required
	//4. Can have comparators as cases too, dont only have to have == relations like c, c++

	switch os := runtime.GOOS; os {
	case "darwin" :
		fmt.Println("OS X")
	case "linux" :
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}
}