package main

import "fmt"

func main(){
	i, j := 42, 2701
	p := &i
	fmt.Println(*p)

	*p = 21 // this changes i
	fmt.Println(i)

	p = &j
	*p = *p / 37 // This changes j

	// NOTE :: GO HAS NO POINTER ARITHMETIC !!!!
	fmt.Println(j)
}
