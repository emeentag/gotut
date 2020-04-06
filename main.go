package main

import "fmt"

func main() {
	conditions(5)
	conditions(15)
	conditions(10)
}

func conditions(x int) {
	if x == 10 {
		fmt.Printf("%d is equal to 10\n", x)
	} else if x > 10 {
		fmt.Printf("%d is greater than 10\n", x)
	} else {
		fmt.Printf("%d is less than 10\n", x)
	}

}