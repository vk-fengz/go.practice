// https://mp.weixin.qq.com/s/OT5Jst4RNZrQ15a4TNy9OA
package main

import (
	"fmt"
)

func main() {
	var data interface{} = "great"

	if data, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", data)
	} else {
		fmt.Println("[not an int] value =>", data) //prints: [not an int] value => 0 (not "great")
	}

	if res, ok := data.(int); ok {
		fmt.Println("[is an int] value =>", res)
	} else {
		fmt.Println("[not an int] value =>", data) //prints: [not an int] value => great (as expected)
	}

	if res, ok := data.(string); ok {
		fmt.Println("[is an string] value =>", res)
	} else {
		fmt.Println("[not an string] value =>", data) //prints: [not an int] value => great (as expected)
	}
}
