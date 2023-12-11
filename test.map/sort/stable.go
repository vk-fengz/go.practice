// sort不保证排序的稳定性（两个相同的值，排序之后相对位置不变），排序的稳定性由sort.Stable来保证。
package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

type personSlice []person

func (s personSlice) Len() int           { return len(s) }
func (s personSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s personSlice) Less(i, j int) bool { return s[i].Age < s[j].Age }

func main() {
	a := personSlice{
		{
			Name: "AAA",
			Age:  55,
		},
		{
			Name: "BBB",
			Age:  22,
		},
		{
			Name: "CCC",
			Age:  0,
		},
		{
			Name: "DDD",
			Age:  22,
		},
		{
			Name: "EEE",
			Age:  11,
		},
	}
	// 递增, 稳定排序
	// sort.Stable(a)
	// fmt.Println(a)
	sort.Sort(sort.Reverse(a))
	fmt.Println("递减: ", a)

}

// [{CCC 0} {EEE 11} {BBB 22} {DDD 22} {AAA 55}]
// 递减:  [{AAA 55} {BBB 22} {DDD 22} {EEE 11} {CCC 0}]
