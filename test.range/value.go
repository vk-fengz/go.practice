package main

import (
	"fmt"
)

type user struct {
	name string
	age  uint64
}

func main() {
	origin := []user{
		{"asong", 23},
		{"song", 19},
		{"asong2020", 18},
	}
	fmt.Println("origin slice:", origin)

	second := make([]user, 0, len(origin))
	for _, v := range origin {
		second = append(second, v)
	}
	fmt.Println("range append:", second)

	fmt.Println("==range print:")
	for _, v := range second {
		fmt.Println(v)
	}

	fmt.Println("==range print one value: index :")
	for k := range second {
		fmt.Println(k)
	}

}
