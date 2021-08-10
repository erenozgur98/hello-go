package main

import (
	"fmt"
)

// here we declare our struct (object)
type person struct {
	name string
	age int
}

func main() {
	p := person{name: "Eren", age: 23}

	fmt.Println(p)
	// here we return an object; {Eren 23}
}