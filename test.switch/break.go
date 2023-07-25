package main

import (
	"fmt"
)

func main() {
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")

	case 5:
		fmt.Println("The integer was <= 5")

	case 6:
		fmt.Println("The integer was <= 6")
		if integer == 6 {
			if integer < 10 {
				fmt.Println("in if if ")
				break //可以直接跳出case 结束swr
			}
			fmt.Println("in if oh no")
		}
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		// not have fallthrough
	case 8:
		fmt.Println("The integer was <= 8")

	default:
		fmt.Println("default case")
	}
}
