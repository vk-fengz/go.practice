package main

import "fmt"

func main() {
	ms := make(map[string][]int)

	ms["aa"] = append(ms["aa"], 1)
	ms["aa"] = append(ms["aa"], 2)
	ms["bb"] = append(ms["bb"], 99)

	fmt.Println(ms)
	for value := range ms {
		fmt.Println(value, ":", ms[value])
	}

	// map[aa:[1 2] bb:[99]]
	// bb : [99]
	// aa : [1 2]
}
