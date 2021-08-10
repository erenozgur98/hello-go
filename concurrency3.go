package main

import (
	"fmt"
)

func main() {
	// simple channel operation to understand what it's actually doing
	c := make(chan string, 2)
	// by adding 2 here, preventing the deadlock. The way we do it is buffered channel by giving the capacity.
	c <- "Hello"

	msg := <- c
	fmt.Println(msg)
}