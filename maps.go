package main

import (
	"fmt"
)

func main() {
	// the way we write maps in Go is to make the map first=>
	// the string is the type of key, and the int is the type of the values inside the map
	newMap := make(map[string]int)

	// to append to the map we just created we can simply write this:
	newMap["triangle"] = 2
	newMap["square"] = 3
	// the values will return: [square: 3 triangle: 2]
	
	// to delete something inside the map:
	delete(newMap, "square")
	// this will get rid of the square inside the map

	fmt.Println(newMap)
} 