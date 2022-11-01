// 有效括号
// ()[]{}  true
package main

import (
    "fmt"
)

// leetcode
func isValid1(s string) bool {
    n := len(s)
    if n%2 == 1 {
        return false
    }
    pairs := map[byte]byte{
        ')': '(',
        ']': '[',
        '}': '{',
    }
    stack := []byte{}
    for i := 0; i < n; i++ {
        if pairs[s[i]] > 0 {
            if len(stack) == 0 || stack[len(stack)-1] != pairs[s[i]] {
                return false
            }
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, s[i])
        }
    }
    return len(stack) == 0
}

func main() {
    str := "()[]{"

    fmt.Println(isValid(str))

}

// 手撕
func isValid(s string) bool {
    if 0 == len(s) {
        return true
    }

    stack := []byte{}
    for _, v := range s {
        if ('(' == v) || ('[' == v) || ('{' == v) {
            stack = append(stack, byte(v))
        } else {
            if (')' == v) && len(stack) > 0 && (stack[len(stack)-1] == '(') ||
                    (']' == v) && len(stack) > 0 && (stack[len(stack)-1] == '[') ||
                    ('}' == v) && len(stack) > 0 && (stack[len(stack)-1] == '{') {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }

    return len(stack) == 0
}
