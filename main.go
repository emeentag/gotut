package main

import (
	"fmt"
	"errors"
	"math"
)

func main() {
	returnMultiple(16)
	returnMultiple(-16)
}

func returnMultiple(num float64) {
	result, error := sqrt(num)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}
}

func sqrt(x float64) (float64, error) {
	if x < 0 {
		return -1, errors.New("Undifined for negative numbers")
	}

	return math.Sqrt(x), nil
}