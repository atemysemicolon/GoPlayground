package main

import "fmt"

func main(){
	// Notes :-
	// 1. defered for Delayed execution
	// 2. The function parameters are executed immediately
	// 3. Call is done after. TODO : After next statement or ...?
	defer fmt.Println("world") //Printed after "hello"
	fmt.Println("hello")
}
