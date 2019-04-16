package main

import "fmt"

func main(){
	fmt.Println("counting")

	// Why does the following print the way it is?
	// Deferred function calls are pushed onto a stack
	// When the first function returns (i.e Println("done") ),
	// the deferred calls are popped out and executed (LIFO order)
	for i:=0;i<10;i++ {
		defer fmt.Println(i) //Woah crazy
	}

	fmt.Println("done")
}
