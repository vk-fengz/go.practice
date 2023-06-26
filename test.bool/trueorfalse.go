package main

import "fmt"

func main() {

	a := true || false
	b := false || true
	fmt.Println(a, b)

	aint := 1
	isTrue := aint == 1
	isFalse := aint == 0
	fmt.Println("aint值:", aint, "  等于1:", isTrue, "  等于0:", isFalse)

}
