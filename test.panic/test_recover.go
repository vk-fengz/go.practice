// recover 对 panic 的捕捉;
package main

import "fmt"

func test() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic")
	}()
	panic("test panic")
}
func main() {
	test()
}
