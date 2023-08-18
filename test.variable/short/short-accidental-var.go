package main

import "fmt"

func main() {
	x := 1
	fmt.Println(x)
	{ //prints 1

		fmt.Println(x) //prints 1
		// x = 3
		x := 2         // 不会影响到外部x变量的值
		fmt.Println(x) //prints 2
		//x = 5        // 不会影响到外部x变量值
	}
	fmt.Println(x) //prints 3
}

// 1
// 1
// 2
// 1
