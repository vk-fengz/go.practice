package main

import "fmt"

func main() {
	var sli []int
	alpha := make(map[string]*[]int)

	sli = append(sli, 1)
	alpha["a"] = &sli

	fmt.Println(alpha)
	sli = append(sli, []int{2, 3}...)
	// alpha["a"] = &sli
	fmt.Println(*alpha["a"])

}
