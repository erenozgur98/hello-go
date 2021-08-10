package main

import (
	"fmt"
	"time"
)

// final construct go has select (switch in JS) statement:
func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()
	go func() {
		for {
			c2 <- "Every 2 seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	// using a select statement here to prevent the half a second one to wait for 2 seconds
	for {
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <- c2:
			fmt.Println(msg2)
		}
	}
}