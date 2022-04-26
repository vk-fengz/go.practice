package main

import (
    "fmt"
	"strings"
	"unicode"
)

// 主函数
func main() {

	str1 := "ABCDEFGHAB"
	str2 := "ABCDEFGHBA"

	fmt.Printf("source string: %s\n", str1)

	// 翻转字符串:
	reverstr, ok := reverseString(str1)
	if true == ok {
		fmt.Printf("REVERSEString: %s\n\n", reverstr)
	}

	// strings.Count: substr count
	fmt.Printf("strings.Count AB: %d\n", strings.Count(str1, "AB"))
	fmt.Printf("strings.Count \"\": %d\n", strings.Count(str1, ""))

	// 字符串中字符是否都不同
	fmt.Printf("isUniqueString: %t\n", isUniqueString(str1))

	// 判断排序后两个字符串是否一致
	fmt.Printf("isRegroup: %t\n",
		isRegroup(str1, str2))

}

// 翻转字符串
func reverseString(s string) (string, bool) {
	str := []rune(s)
	le := len(str)
	if le > 50000 {
		return s, false
	}
	for i := 0; i < le/2; i++ {
		str[i], str[le-1-i] = str[le-1-i], str[i]
	}
	return string(str), true
}

// 判断字符串中是否都不同  strings.Count
func isUniqueString(s string) bool {
	if strings.Count(s, "") > 3000 {
		return false
	}
	for _, v := range s {
		if v > 127 {
			return false
		}
		if strings.Count(s, string(v)) > 1 {
			return false
		}
	}
	return true
}

// 判定两个给定的字符串排序后 是否一致;
func isRegroup(s1, s2 string) bool {
	sl1 := len([]rune(s1))
	sl2 := len([]rune(s2))
	if sl1 > 5000 || sl2 > 5000 || sl1 != sl2 {
		return false
	}
	for _, v := range s1 {
		if strings.Count(s1, string(v)) == strings.Count(s2, string(v)) {
			return true
		}
	}
	return true
}

// 字符串替换问题
func replaceBlank(s string) (string, bool) {
	if len([]rune(s)) > 1000 {
		return s, false
	}
	for _, v := range s {
		if string(v) != " " && unicode.IsLetter(v) == false {
			return s, false
		}

	}
	return strings.Replace(s, " ", "%20", -1), true
}

// 机器人坐标问题