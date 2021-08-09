package main

import (
	"fmt"
)

func main() {
	result := sum(2, 3)
	fmt.Println(result)
	// here will return the sum of 2+3 which is 5
}

// here we declare our second function, which takes two integers (we have to declare what type are they) then retuns another integer (we also have to declare the type) if you want to return a value
func sum(x int, y int) int {
	return x+y
}