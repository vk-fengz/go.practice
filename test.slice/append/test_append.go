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

	// make
	// var names []string
	names := make([]string, 3)
	fmt.Println("names:", names, len(names))
	names=append(names, "aaa")
	names = append(names, "bbb")
	fmt.Println("names:", names)
	names = append(names, "ccc")
	fmt.Println("names append ccc:", names)

	names=append(names, "ddd")
	fmt.Println("names append ddd:", names)


	

		var s []string
		fmt.Println("uninit:", s, s == nil, len(s) == 0)
	
		s = make([]string, 3)
		fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))
	
		s[0] = "a"
		s[1] = "b"
		s[2] = "c"
		fmt.Println("set:", s)
		fmt.Println("get:", s[2])
	
		fmt.Println("len:", len(s))
	
		s = append(s, "d")
		s = append(s, "e", "f")
		fmt.Println("apd:", s)
	
		c := make([]string, len(s))
		copy(c, s)
		fmt.Println("cpy:", c)
	
		l := s[2:5]
		fmt.Println("sl1:", l)
	
		l = s[:5]
		fmt.Println("sl2:", l)
	
		l = s[2:]
		fmt.Println("sl3:", l)
	
		t := []string{"g", "h", "i"}
		fmt.Println("dcl:", t)
	
		t2 := []string{"g", "h", "i"}
		if slices.Equal(t, t2) {
			fmt.Println("t == t2")
		}
	
		twoD := make([][]int, 3)
		for i := 0; i < 3; i++ {
			innerLen := i + 1
			twoD[i] = make([]int, innerLen)
			for j := 0; j < innerLen; j++ {
				twoD[i][j] = i + j
			}
		}
		fmt.Println("2d: ", twoD)
	

}
