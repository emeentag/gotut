package main

import "fmt"

func main() {
	slices()
}

func slices() {
	var ss []int
	ss = append(ss, 1, 2, 3, 4, 5)
	fmt.Println(ss)

	ss2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(ss2)
}