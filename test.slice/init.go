package main

import "fmt"

func main() {
	var sli []int
	fmt.Println(sli)
	// panic: runtime error: index out of range [0] with length 0
	// fmt.Println(&sli[0])
	sli = []int{1, 2}
	fmt.Println(sli)
}
