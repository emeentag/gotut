package main

import "fmt"

func main() {
	loops()
}

func loops() {
	for i := 0; i < 5; i++ {
		fmt.Printf("idx: %d\n", i)
	}
	fmt.Println("------------loop like while-------------")
	j := 0
	for j < 5 {
		fmt.Printf("jdx: %d\n", j)
		j++
	}
	fmt.Println("------------loop in arr-------------")
	arr := []int{1, 2, 3, 4, 5}
	for idx, val := range arr {
		fmt.Println("index", idx, "value", val)
	}
}