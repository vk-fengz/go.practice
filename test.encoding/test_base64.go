package main

import (
    "encoding/base64"
    "fmt"
)

func main() {
    originalByte := []byte("hello world")

    // base64 编码
    encodeString := base64.StdEncoding.EncodeToString(originalByte)

    // encodeStringNew := base64.NewEncoder(originalByte, w)
    fmt.Println("标准编码结果:", encodeString)

    // base64 解码
    decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("标准解码结果:", string(decodeBytes))
    fmt.Println("--------------------------------------------------")

    // 在 url 中使用时,应该用 URLEncoding 编码
    urlString := base64.URLEncoding.EncodeToString([]byte(originalByte))
    fmt.Println("url 编码结果:", urlString)

    // 解码
    urlByte, err := base64.URLEncoding.DecodeString(urlString)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println("url 解码结果:", string(urlByte))

}
