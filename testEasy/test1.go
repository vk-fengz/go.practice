package main

import "fmt"




func main() {
	x := []int{1, 2, 3}
	i := 2
	switch i {
	case x[1]:
		println("a")
		fallthrough
	case 1, 3:
		println("b")
	default:
		println("c")
	}
	// 数组打印
	fmt.Println(x)

	//==append
	var s = make([]int, 0, 5)
	fmt.Printf("%p\n", &s)
	s2 := append(s,1)
	fmt.Printf("%p\n", &s2)

	fmt.Println(s, s2)

	//Map
	println("==Map")
	m := map[int] struct {
		name string
		age	 int
	}{
		1: {"user1", 10},
		2: {"user2", 20},
	}

	println(m[1].name)

	//迭代删除键值   新增可能出错;
	for i := 0; i < 5; i++ {
		m := map[int]string{
			0: "a", 1: "a", 2: "a", 3: "a", 4: "a",
			5: "a", 6: "a", 7: "a", 8: "a", 9: "a",
		}
		for k := range m {
			m[k+k] = "x"
			delete(m, k)
		}
		fmt.Println(m)
	}













}
