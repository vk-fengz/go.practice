// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
// 链接：https://leetcode-cn.com/problems/longest-substring-without-repeating-characters
// "dvdf"

package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "dvvdf"
	fmt.Println(str)
	num := lengthOfLongestSubstring(str)
	fmt.Println(num)
}

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	le, ret := 0, 0
	strtemp := ""
	for _, v := range s {
		if strings.Count(strtemp, string(v)) == 0 {
			le += 1
			strtemp += string(v)
			continue
		}
		if ret < le {
			ret = le
		}
		i := strings.Index(strtemp, string(v))
		strtemp = strtemp[i+1:] + string(v)
		le = len(strtemp)
	}
	if ret < le {
		return le
	}
	return ret
}
