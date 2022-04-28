package main

import "fmt"

// Product
func Product(sets ...[]interface{}) [][]interface{} {
	lens := func(i int) int { return len(sets[i]) }
	product := [][]interface{}{}
	for ix := make([]int, len(sets)); ix[0] < lens(0); nextIndex(ix, lens) {
		var r []interface{}
		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		product = append(product, r)
	}
	return product
}

func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

func main() {

	sets := Product([]interface{}{"a", "b", "c"}, []interface{}{1, 2, 3})
	for _, set := range sets {
		fmt.Println(set)
	}

	// Ordered Output:
	// [a 1]
	// [b 1]
	// [c 1]
	// [a 2]
	// [b 2]
	// [c 2]
	// [a 3]
	// [b 3]
	// [c 3]
}
