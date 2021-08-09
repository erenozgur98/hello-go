package main

import (
	"fmt"
)

func main() {
	x := 5
	
	// simple if statement, just like in JavaScript but no () around the condition
	if x > 6 {
		fmt.Println(x, "is bigger than 6")
	} else {
		fmt.Println(x, "is smaller than 6")
	}
}