package main

import "fmt"

func main() {
	arrays()
}

func arrays() {
	var arr [5]int
	arr[2] = 7
	fmt.Println(arr)

	arr2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr2)
}