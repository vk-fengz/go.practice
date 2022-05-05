package main

import "fmt"

type index struct {
	length int
	alpha  []string
	remark []string
}

func main() {

	indexTable := make(map[string]*index)

	fmt.Println(indexTable)
	if indexTable["aa"] == nil {
		indexTable["aa"] = &index{}
	}
	// indexTable["aa"].alpha = append(indexTable["aa"].alpha, "first")
	// indexTable["aa"].remark = append(indexTable["aa"].remark, "The first one")

	indexTable["aa"] = &index{
		1,
		append(indexTable["aa"].alpha, "first"),
		append(indexTable["aa"].remark, "The first one"),
	}

	fmt.Println(indexTable)
	fmt.Println(indexTable["aa"])

	indexTable["aa"] = &index{
		1,
		append(indexTable["aa"].alpha, "sec"),
		append(indexTable["aa"].remark, "The second one"),
	}

	fmt.Println(indexTable["aa"])

}
