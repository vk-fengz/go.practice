package __longest_palindromic_substring

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		// 回文串 奇数个
		s1 := palindrome(s, i, i)
		// 偶数个元素的 回文字符串
		s2 := palindrome(s, i, i+1)
		res = max(res, s1)
		res = max(res, s2)
	}
	return res
}

func palindrome(s string, l, r int) string {
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	return s[l+1 : r]
}

func max(i, j string) string {
	if len(i) > len(j) {
		return i
	}
	return j
}
