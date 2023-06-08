package main

import (
	"fmt"
	"sync"
)

var m sync.Map

func main() {

	//写
	m.Store("dablelv", "27")
	m.Store("cat", "28")

	//读
	v, ok := m.Load("dablelv")
	fmt.Printf("Load: v, ok = %v, %v\n", v, ok)

	//删除
	m.Delete("dablelv")

	//读或写
	v, ok = m.LoadOrStore("dablelv", "18")
	fmt.Printf("LoadOrStore: v, ok = %v, %v\n", v, ok)

	//遍历
	//操作函数
	f := func(key, value interface{}) bool {
		fmt.Printf("Range: k, v = %v, %v\n", key, value)
		return true
	}
	m.Range(f)
}
