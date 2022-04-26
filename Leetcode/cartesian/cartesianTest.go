package main

import (
	"code.practice/Leetcode/cartesian/cartesian"
	"fmt"
)

func main() {

	a := []interface{}{1, 2, 3}
	b := []interface{}{"a", "b", "c"}
	c := []interface{}{"X", "Y"}

	result := cartesian.Iter(a, b, c)

	fmt.Println(len(result))
	// receive products through channel
	for product := range result {
		fmt.Println(product)
	}

	// Unordered Output:
	// [1 c]
	// [2 c]
	// [3 c]
	// [1 a]
	// [1 b]
	// [2 a]
	// [2 b]
	// [3 a]
	// [3 b]
}
