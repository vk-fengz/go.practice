// Roman-to-Integer 罗马数字转换为整数;
package main

import "fmt"

var symbolValues = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) (ans int) {
	n := len(s)
	for i := range s {
		value := symbolValues[s[i]]
		if i < n-1 && value < symbolValues[s[i+1]] {
			ans -= value
		} else {
			ans += value
		}
	}
	return
}

func main() {
	str := "MCMXCIV" // 1994
    ret:=romanToInt(str)
	ret1 := romanToInt1(str)
	fmt.Println("romanToInt:", ret)
	fmt.Println("romanToInt1:", ret1)
}

// 手撕
func romanToInt1(s string) int {
	if "" == s {
		return 0
	}
	i, temp, value := 0, 0, 0
	le := len(s)
	for i, _ = range s {
		temp = symbolValues[s[i]]
		if i < le-1 && temp < symbolValues[s[i+1]] {
			value -= temp
		} else {
			value += temp
		}
	}

	return value
}


