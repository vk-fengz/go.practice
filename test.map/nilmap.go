package main

import "fmt"

func main() {
	var book map[string]string

	fmt.Println(book)
	if book == nil {
		fmt.Println("map var is nil")
	}
	if len(book) == 0 {
		fmt.Println("map var is empty")
	}
	book = make(map[string]string)
	fmt.Println(book)
	if book == nil {
		fmt.Println("map make is nil")
	} else {
		fmt.Println("map make not nil")
	}
	if len(book) == 0 {
		fmt.Println("map make is empty")
	} else {
		fmt.Println("map make not empty")
	}

	book["title"] = "love"
	fmt.Println(book)
}
