package main

import "fmt"

func main(){
	sum := 1

	// No semicolons required -- look!
	for sum < 1000 {
		sum += sum
	}

	fmt.Println(sum)
}
