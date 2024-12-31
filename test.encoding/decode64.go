package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
)

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func main() {
	// 模拟Base64编码后的JSON数据
	base64Data := []byte("eyJuYW1lIjoiSm9obiIsImFnZSI6MzAsImVtYWlsIjoiam9obi5leGFtcGxlLmNvbSJ9")

	// 先进行Base64解码
	decoded, err := base64.StdEncoding.DecodeString(string(base64Data))
	if err != nil {
		fmt.Println("Base64解码错误:", err)
		return
	}

	var user User
	// 再进行JSON解码
	err = json.Unmarshal(decoded, &user)
	if err != nil {
		fmt.Println("JSON解码错误:", err)
		return
	}

	fmt.Printf("姓名: %s, 年龄: %d, 邮箱: %s\n", user.Name, user.Age, user.Email)
}
