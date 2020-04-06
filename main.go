package main

import "fmt"

func main() {
	var sum int = sum(2, 3)
	fmt.Println(sum)

	sum = sumWithoutTypeDef(2, 4)
	fmt.Println(sum)
}

func sumWithoutTypeDef(x int, y int) int {
	sum := x + y

	return sum
}

func sum(x int, y int) int {

	var sum int = x + y

	return sum
}