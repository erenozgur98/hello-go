package main

import (
	"fmt"
)

func main() {
	// if an array declared like this, the length will be 5 and can not add another item inside the array.
	var a [5]int
	// if no value is declared, the value will always be 0 or an empty string, in this case it'll be => [0 0 0 0 0]
	// to give it a value;
	a[2] = 7
	fmt.Println(a)
	// this will return the array => [0 0 7 0 0]

	// the neater version is:
	// we have to add them inside an object after the type
	y := [5]int{5, 4, 3, 2, 1}
	fmt.Println(y)

	// if we want to add into the array after we create it, this is how we do it:
	// but we should leave the array empty => [] so we can append to it.
	z := []int{5, 4, 3, 2, 1}

	z = append(z, 0)
	fmt.Println(z)
}