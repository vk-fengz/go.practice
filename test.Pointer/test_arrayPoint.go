// 切片底层数组 元素的存放;
package main

import (
    "fmt"
    "unsafe"
)

func main() {
    var s1 = []int{1, 2, 3, 4, 5, 6, 7}
    var s2 = s1[2:5]
    fmt.Println(s1)
    fmt.Println(s2)
    fmt.Println(cap(s1), cap(s2))

   // unsafe.Pointer(&s1) 取到s1的首地址的指针, 其实也就是s1 array位置的指针
   // unsafe.Pointer(*(*unsafe.Pointer)(unsafe.Pointer(&s1))) 将s1 array存放的值取出来, 也就是底层数组的 首元素的地址
   // unsafe.Pointer(uintptr(s1SliceArrayPointer) + unsafe.Sizeof(&s1[0])*2) 向后偏移2位, 也就是取到3的位置

    s1SliceArrayPointer := unsafe.Pointer(*(*unsafe.Pointer)(unsafe.Pointer(&s1)))
    fmt.Println("s1[0]:", *(*int)(s1SliceArrayPointer))
    fmt.Println("s1 address:", s1SliceArrayPointer)

    s2SliceArrayPointer := unsafe.Pointer(*(*unsafe.Pointer)(unsafe.Pointer(&s2)))
    fmt.Println("s2[0]:", *(*int)(s2SliceArrayPointer))
    fmt.Println("s2 address:", s2SliceArrayPointer)

    offset2Pointer := unsafe.Pointer(uintptr(s1SliceArrayPointer) + unsafe.Sizeof(&s1[0])*2)
    fmt.Println("offset s1[2]:", *(*int)(offset2Pointer))
    fmt.Println("offset s1 address:", offset2Pointer)
}
