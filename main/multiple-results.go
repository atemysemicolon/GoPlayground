package main

import "fmt"

func swap (x, y string) (string, string) {
	return y, x
}


func main() {
	a, b := swap("hello", "world") // Short variable declarations --> NOTE no type
	fmt.Println(a, b)
}