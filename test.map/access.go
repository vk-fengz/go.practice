package main

import "fmt"

func main() {

	m := make(map[string]bool, 0)

	if m["a"] == false {
		fmt.Println(m)
		fmt.Println("access nil success.")
	}
}
