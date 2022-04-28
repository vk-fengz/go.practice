// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type M struct {
	A string
}

func main() {
	var sa = make([]string, 0)
	for i := 0; i < 10; i++ {
		sa = append(sa, fmt.Sprintf("%v", i))
	}
	print(sa)
	print(interface{}(sa))
	printM(M{"abc"})

}
func print(a interface{}) {
	fmt.Println(a)
}
func printM(a M) {
	fmt.Println(a)
}