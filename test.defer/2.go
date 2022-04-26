package main

import (
    "fmt"
    "strconv"
)

func test() (i int) {
	i = 0
	defer func() {
		fmt.Println("\ndefer1")
	}()
	defer func() {
		i += 1
		fmt.Print("defer2: ", i)
	}()
	return i
}

func main() {
	fmt.Println("return: ", strconv.Itoa(test()))
}
