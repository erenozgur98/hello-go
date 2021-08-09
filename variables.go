package main

import (
	"fmt"
)

func main() {
	// to declare a variable write var then the name of the variable, then type of the variable in this case an integer.
	var x int
	// then we can add into the x variable we just created:
	x = 5

	// or we can do this easier:
	var y int = 5

	// even simpler version of this is goes like this:
	// no need to declare the type of the variable here because Go will know what type it is
	z := 5

	// all the variables we created will print out the same exact thing.
	fmt.Println(x, y, z)
}