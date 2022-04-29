// map 重复赋值, 会覆盖, key-value
// 可以修改int值, string 不行;
package main

import "fmt"

type Student struct {
	Name string
	Id   string // int
}

func main() {
	s1 := make(map[string]string)
	s1["chenchao"] = "aa"
	s1["chenchao"] = "bb"

	fmt.Println(s1)
	for i := 0; i <= 10; i++ {
		fmt.Println(s1["chenchao"])
	}

}
