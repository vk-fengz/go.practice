package main

import (
	"fmt"
	"reflect" // The reflect library is used to obtain the datatype of an object.
	"strconv"
)

func main() {

	// 整数转换为string
	i := 10
	s1 := strconv.FormatInt(int64(i), 10)
	s2 := strconv.Itoa(i)
	s3 := fmt.Sprint(i)
	fmt.Printf("s1, s2, s3: %v, %v, %v\n", s1, s2, s3)

	// print 2
	num1 := 10
	fmt.Println(num1, reflect.TypeOf(num1))
	str := strconv.Itoa(num1)
	fmt.Println(str, reflect.TypeOf(str))
}
