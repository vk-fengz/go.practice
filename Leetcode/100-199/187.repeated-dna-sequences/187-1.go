package _87_repeated_dna_sequences

import (
	"math"
)

// 效率不高, 见官方解法
func findRepeatedDnaSequences(s string) []string {
	// 将字符串 转化成 四进制的数字数组
	numstr := []int{}
	res := []string{}
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			numstr = append(numstr, 0)
			break
		case 'G':
			numstr = append(numstr, 1)
		case 'C':
			numstr = append(numstr, 2)
		case 'T':
			numstr = append(numstr, 3)
		}
	}

	// 记录重复出现的hash值
	seen := make(map[float64]struct{})
	// 记录重复出现的字符串结果
	result := make(map[string]struct{})

	// 数字位数L  进制R  // 存储 R^(L - 1) 的结果
	L, R := 10.0, 4.0
	rl := math.Pow(R, L-1)
	// 维护滑动窗口的hash值
	windowHash := 0.0

	// 滑动窗口代码框架  时间O(N)
	left, right := 0, 0
	for right < len(numstr) {
		windowHash = windowHash*R + float64(numstr[right])
		right++

		// 当子串长度 达到要求
		if float64(right-left) == L {
			// 根据hash值 判断子串是否重复出现
			if _, ok := seen[windowHash]; ok {
				result[s[left:right]] = struct{}{}
			} else {
				seen[windowHash] = struct{}{}
			}

			windowHash = windowHash - rl*float64(numstr[left])
			left++
		}
	}

	for str, _ := range result {
		res = append(res, str)
	}

	return res
}
