package main

import "fmt"

func main() {
	maps()
}

func maps() {
	vertices := make(map[string]int)
	vertices["one"] = 1
	vertices["two"] = 2
	vertices["three"] = 3

	fmt.Println(vertices)
	fmt.Println(vertices["two"])

	delete(vertices, "two")
	fmt.Println("deleted", vertices)
}