// 切片的拼接
package main

import (
	"fmt"
)

func main() {
	var a = []int{1, 2}
	a = append([]int{0}, a...)
	fmt.Println(a)
	a = append(a[:2], []int{3, 4}...)
	fmt.Println(a)

	// var
	var eslice []int
	eslice = append(eslice, 1)
	eslice = append(eslice, []int{2, 3}...)
	fmt.Println("eslice:", eslice)
	eslice = append(eslice, a...)
	fmt.Println("eslice append a:", eslice)

}
