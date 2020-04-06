package main

import (
	"fmt"
)

func main() {
	pointers()
}

func pointers() {
	x := 5
	inc(x)
	fmt.Println(x, &x)

	fmt.Println("------------pointer-------------")

	incP(&x)
	fmt.Println(x, &x)
}

func inc(p int) {
	p++
}

func incP(p *int) {
	*p++
}