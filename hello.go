package main

// choose a package from library to use inside Go application
// fmt is one of the most related in go
// no commas (,) after each package
import (
	"fmt"
)

// to declare function, simply write the func before the key word. The first function has to be main
func main() {
	fmt.Println("Hello Go!")
}