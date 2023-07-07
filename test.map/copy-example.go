package main

import "fmt"

func main() {
	strIntMap := map[string]int64{
		"Anna": 78,
		"Bob":  82,
	}

	intIntMap := map[int64]int64{
		0: 14,
		1: 23,
	}

	copyMap1 := copyMap(strIntMap)
	copyMap2 := copyMap(intIntMap)

	strIntMap["Test"] = 52
	fmt.Println("The orginal map:")
	fmt.Println(strIntMap)
	fmt.Println(intIntMap)

	fmt.Println("---------------")
	fmt.Println("The copy map:")
	fmt.Println(copyMap1)
	fmt.Println(copyMap2)
}

func copyMap[K, V comparable](m map[K]V) map[K]V {
	result := make(map[K]V)
	for k, v := range m {
		result[k] = v
	}
	return result
}
