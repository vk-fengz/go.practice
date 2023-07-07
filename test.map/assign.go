package main

import "fmt"

func main() {

	//scale := make(map[string]int, 0)
	//scaleDst := make(map[string]int, 0)

	scale := make(map[string]string, 0)
	scaleDst := make(map[string]string, 0)

	fmt.Println("The original map:")
	fmt.Println("map scale: ", scale)
	fmt.Println("map scaleDst: ", scaleDst)

	scaleDst["aa"] = scale["aa"]

	fmt.Println("--------")
	fmt.Println("The final map")
	fmt.Println("map scale: ", scale)
	fmt.Println("map scaleDst: ", scaleDst)

}
