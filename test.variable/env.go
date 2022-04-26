// main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("Path")
	fmt.Println("name is:", name)
}
