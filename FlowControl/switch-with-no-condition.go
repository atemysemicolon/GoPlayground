package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// 1. No condition switch
	// 2. Notice how we have relational operators
	// 3. Cases are executed top-down => simplifying complex conditions in other languages
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning!")
	case t.Hour() < 17:
		fmt.Println("Good Afternoon!")
	default:
		fmt.Println("Good Evening.")
	}
}
