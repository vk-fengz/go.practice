package main

import "fmt"

type student struct {
    Name string
    //Age int
}

func main()  {
    m:=make(map[string]*student)
    stus:= []student{
        {Name: "Zhou"},
        {Name: "jie"},
        {Name: "lun"},
    }

    for _,stu := range stus{
        m[stu.Name] = &stu
        fmt.Printf("%p\n ", &stu)
        fmt.Println(stu)
    }
    fmt.Println(m)
    fmt.Printf("\n%p\n", &stus[0])
    fmt.Printf("%p\n", &stus[1])
    fmt.Printf("%p\n", &stus[2])
    fmt.Printf("%p\n", &stus)
}






