package main

import (
	"fmt"
	"time"
	// adding new libraries to use
)

func main() {
	c := make(chan string)
	go count("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

// we now will add our channel here, as a pipe then will use it in the main function
func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing // sending the message out with <-
 		time.Sleep(time.Millisecond * 500)
	}

	// we can close the channel like this:
	close(c)
}