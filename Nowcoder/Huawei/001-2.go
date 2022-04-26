package main

import(
    "bufio"
    "strings"
    "os"
    "fmt"
)

func main() {
    input := bufio.NewScanner(os.Stdin)
    for input.Scan() {
        arr := strings.Split(input.Text(), " ")
        fmt.Println(len(arr[len(arr)-1]))
    }
}