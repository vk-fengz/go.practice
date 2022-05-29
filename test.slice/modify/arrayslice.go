package main

import "fmt"

func main() {
	a := [4]int{1, 2, 3, 4}
	b := []int{100, 200, 300}

	c := a
	d := b
	c[1] = 100
	d[0] = 1

	fmt.Println("a=", a, "c=", c)
	fmt.Println("b=", b, "d=", d)
	// a= [1 2 3 4] c= [1 100 3 4]
	// b= [1 200 300] d= [1 200 300]
}
