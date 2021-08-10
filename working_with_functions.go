package main

import (
	"fmt"
	"errors"
	"math"
	// here we bring errors and math libraries to be used inside the functions
)

func main() {
	result, err := sqrt(16)

	if err != nil {
		fmt.Println(err)
		// if the number we pass thru sqrt function is negative, we will return the error message
	} else {
		fmt.Println(result)
		// if the number is positive, we will return the square root of 16 which is 4
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numbers")
	}

	return math.Sqrt(x), nil
	// We have to return 2 values because that's how we declared in the function, instead of nothing in the second value we return nil which is null
}