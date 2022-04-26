package main

import (
    "fmt"
    "os"
)

func main() {
    fileName := "./test.txt"
    writeFile(fileName)
    readFile(fileName)

}

func writeFile(fileName string) {
    file, err := os.Create(fileName)

    if err != nil {
        fmt.Println(err)
        return
    }

    for i := 0; i <= 5; i++ {
        outStr := fmt.Sprintf("%s:%d\n", "hello, world", i)

        _, _ = file.WriteString(outStr)
        _, _ = file.Write([]byte("abcd\n"))

    }

    _ = file.Close()
}

func readFile(fileName string) {
    file, err := os.Open(fileName)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()  // 关闭文件;

    buf := make([]byte, 1024)

    for {
        n, _ := file.Read(buf)

        if n == 0 {
            // 0	表示到达EOF
            break
        }
		_, _ = os.Stdout.Write(buf)
    }
}
