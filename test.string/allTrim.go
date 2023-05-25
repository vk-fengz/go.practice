package main

import (
	"fmt"
	"strings"
)

func main() {
	inx := "59fb8210-65ba-4acb-9327-1d45c793ad1f-ct6-rs-0"
	uid := "59fb8210-65ba-4acb-9327-1d45c793ad1f"

	// 很不靠谱
	a2 := strings.TrimLeft(inx, uid)
	a1 := strings.TrimLeft(inx, uid+"-")

	fmt.Println(len("-"))
	a3 := inx[len(uid)+1:]
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}
