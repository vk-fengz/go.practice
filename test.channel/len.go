// 测试 channel len
// 结论:
// 1.

package main

import "fmt"

func main() {
	c := make(chan int, 34)
	// c := make(chan int)
	fmt.Println("1.len:", len(c))

	for i := 0; i < 34; i++ {
		c <- 0
	}
	fmt.Println("2.len:", len(c))

	<-c
	<-c

	fmt.Println("3.len:", len(c))
}
