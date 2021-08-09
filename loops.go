package main

import (
	"fmt"
)

func main() {
	// example for loop:
	arr := []string{"a", "b", "c"}

	for index, value := range arr {
		fmt.Println("index:", index, "value:", value)
		// this will return: index:0 value: a, then will loop through other items and print the same thing until it's done.
	}
}

func loopDefinition() {
	// there is only for loops in Golang but it can be turned into while loops
	// to write a traditional for loop:
	// again, no () around the conditional
	for i := 0; i < 5; i++ {}
	
	// changing this to a while loop:
	i := 0
	
	for i < 5 {}

}