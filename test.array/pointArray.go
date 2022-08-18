package main

import "fmt"

func main() {
	arr := [8]int{}
	for i := 0; i < 8; i++ {
		arr[i] = i
	}

	fmt.Println(arr)
	exchangeByAddress(&arr)
	fmt.Println(arr)
}

func exchangeByAddress(arr *[8]int) {
	for k, v := range *arr {
		arr[k] = v * 2
	}
}

// [0 1 2 3 4 5 6 7]
// [0 2 4 6 8 10 12 14]
