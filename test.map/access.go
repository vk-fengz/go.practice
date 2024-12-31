package main

import "fmt"

func main() {

	m := make(map[string]bool, 0)
	mp := make(map[string]*int, 0)
	mps := make(map[string]string, 0)

	if m["a"] == false {
		fmt.Println(m)
		fmt.Println(m["b"])
		fmt.Println("access nil success.")
	}
	fmt.Println(mp["a"])
	if mp["b"] == nil {
		fmt.Println("access map success.")
	}

	fmt.Println(mps["a"] == "")

}
