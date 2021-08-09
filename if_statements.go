package main

import (
	"fmt"
)

func main() {
	x := 5
	if x > 6 {
		fmt.Println(x, "is bigger than 6")
	} else {
		fmt.Println(x, "is smaller than 6")
	}
}