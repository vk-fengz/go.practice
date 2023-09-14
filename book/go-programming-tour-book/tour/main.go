package main

import (
	"log"
	"vk-fengz/go.practice/book/go-programming-tour-book/tour/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
