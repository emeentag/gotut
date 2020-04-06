package main

import (
	"fmt"
)

type person struct {
	name string
	weight int
}

func main() {
	runStruct()
}

func runStruct() {
	p := person{name: "serdar", weight: 75}
	fmt.Println(p)
	fmt.Println(p.weight)
}