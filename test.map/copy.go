package main

import "fmt"

func main() {

	var nilmap map[string]int
	fmt.Println(len(nilmap))

	studentScore := map[string]int{
		"Anna":     59,
		"Bob":      74,
		"Charlice": 96,
	}

	// copy len
	fmt.Println("copy map destination len:", len(nilmap)+len(studentScore))
	// copy a map
	studentScoreCopy := make(map[string]int, len(nilmap)+len(studentScore))
	for k, v := range studentScore {
		studentScoreCopy[k] = v
	}
	studentScore["Daniel"] = 85

	fmt.Println("The original map:")
	fmt.Println(studentScore)
	fmt.Println("--------")

	fmt.Println("The copied map")
	fmt.Println(studentScoreCopy)

	studentScoreCopy = studentScoreCopy
}
