// 使用struct{} 做value

package main

func main() {
	mm := make(map[string]map[string]struct{})

	mm["a"] = map["aa"]struct{}{}

}
