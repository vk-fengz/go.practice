// https://mp.weixin.qq.com/s/ea7LfF2jOoHOSozX-qUZLA

package main

import (
	"fmt"
	"runtime"
)

type TestStruct struct {
	name string
}

//go:noinline
func newTestStruct() *TestStruct {
	v := &TestStruct{"n1"}
	runtime.SetFinalizer(v, func(p *TestStruct) {
		fmt.Println("gc Finalizer")
	})
	return v
}

func main() {
	t := newTestStruct()
	fmt.Println("== start ===")
	_ = t
	fmt.Println("== ... ===")
	runtime.GC() // 调用gc
	fmt.Println("== end ===")
}
