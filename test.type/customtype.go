package main

import "fmt"

func main() {
	type foo int
	a := foo(1)

	type bar []int
	c := []int{1, 2, 3}
	// d := bar(c)
	d := append(bar{}, c...)

	fmt.Println(a)
	fmt.Println(d)
}
