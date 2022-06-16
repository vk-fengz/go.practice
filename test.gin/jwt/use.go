package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	// 创建一个 JWT
	mySigningKey := []byte("woxiangbianqiang!")
	c := MyClaims{
		Username: "qimiao",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,      // 生效时间 1 分钟
			ExpiresAt: time.Now().Unix() + 2*60*60, // 失效时间 2 小时
			Issuer:    "lzy",                       // 签发者
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c) // 创建 Token
	s, _ := t.SignedString(mySigningKey)              // 签发 Token
	fmt.Println(s)                                    // token 的字符串

	// 解析 JWT
	token, err := jwt.ParseWithClaims(s, &MyClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
	if err != nil {
		fmt.Println(err)
	}
	// 对 Claims 进行断言并使用
	fmt.Println(token.Claims.(*MyClaims))
}
