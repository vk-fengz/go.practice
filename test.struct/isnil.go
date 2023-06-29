package main

import (
	"fmt"
)

type Person struct {
	name  string
	age   int
	email int
}

func main() {
	var p1 Person

	p2 := Person{
		name: "John",
		age:  45,
	}

	fmt.Println(p1 == Person{}) // true
	fmt.Println(p2 == Person{}) // false
}
