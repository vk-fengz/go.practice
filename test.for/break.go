package main

import "fmt"

func main() {
	dic := map[string]int{
		"a": 1,
		"b": 1,
		"c": 0,
	}

	var res string
	for _, num := range dic {
		if num != 1 {
			res = "red"
			break
		}
		res = "green"
	}
	fmt.Println(res)
}
