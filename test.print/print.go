// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Employee struct {
	Id   int64
	Name string
}

func main() {
	fmt.Println("Hello, 世界")
	a := "hello"
	a = fmt.Sprintf("%s-", a)

	fmt.Println(a)

	fmt.Println("====> Print struct")
	Employee_1 := Employee{Id: 10, Name: "Dixya Lhyaho"}
	fmt.Printf("%+v\n", Employee_1)               // with Variable name
	fmt.Printf("Employee_1:\n %+v\n", Employee_1) // Without Variable Name
	fmt.Printf("%v\n", Employee_1)                // Without Variable Name

	fmt.Printf("%d\n", Employee_1.Id)
	fmt.Printf("%s\n", Employee_1.Name)

}
