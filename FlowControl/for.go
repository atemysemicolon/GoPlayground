package main


import "fmt"

func main(){
	sum := 0
	// Note 1 --> No Paranthesis around for
	// The braces for the for-body are ALWAYS required - Both unlike c, c++
	for i :=0; i<10; i++ {
		sum += i
	}
	fmt.Println(sum)
}