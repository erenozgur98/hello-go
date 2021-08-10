package main

import (
	"fmt"
	"time"
	"sync"
	// adding new libraries to use
)

func main() {
	var wg sync.WaitGroup // this is our counter
	wg.Add(1)

	go func() {
		count("sheep")
		wg.Done()
	}()

	wg.Wait()
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}