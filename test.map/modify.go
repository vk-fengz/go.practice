// 错误的修过
package main

import "fmt"

type student struct {
	Name string
	Id   int
}

func main() {
	s := make(map[string]student)
	s["chenchao"] = student{
		Name: "chenchao",
		Id:   111,
	}
	// 报错, 需要使用指针;
	// s["chenchao"].Id = 222

	s["chenchao"] = student{
		Name: "chenchao",
		Id:   222,
	}

	fmt.Println(s)

}
