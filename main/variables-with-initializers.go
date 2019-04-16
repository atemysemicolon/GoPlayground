package main

import "fmt"

// when its a list, type can be omitted
//var i,j int = 1, 2 // notice that the int can be removed
var i,j = 1, 2

func main(){
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}