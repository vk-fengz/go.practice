package main

import (
	"fmt"
)

func main() {
	src := []string{"a", "b", "c"}
	dst := make([]string, len(src))

	copy(dst, src)

	fmt.Printf("source slice: %[1]v, address: %[1]p\n", src)
	fmt.Printf("dst slice: %[1]v, address: %[1]p\n\n", dst)

	fmt.Println("-- initial")
	source := []int{1, 2, 3, 4}
	second := []int{4, 5, 6}
	second = append(second, source...)
	fmt.Printf("source slice: %[1]v, address: %[1]p\n", source)
	fmt.Printf("second slice: %[1]v, address: %[1]p\n", second)

	copy(source, second)
	fmt.Println("-- after copy --")
	fmt.Printf("source slice: %[1]v, address: %[1]p\n", source)
	fmt.Printf("second slice: %[1]v, address: %[1]p\n", second)

	fmt.Println("------------")
	source = second
	fmt.Printf("source slice: %[1]v, address: %[1]p\n", source)
	fmt.Printf("second slice: %[1]v, address: %[1]p\n", second)

}
