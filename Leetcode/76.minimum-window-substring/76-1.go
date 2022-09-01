package _6_minimum_window_substring

// 滑动窗口代码框架

// m1. 使用map[char]int 来记录
func minWindow(s string, t string) string {
	// 获得键值表 need: 目标值;
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	// valid 验证满足与否
	left, right, valid := 0, 0, 0
	// 记录最小覆盖子串的起始索引及 结束
	finalLeft, finalRight, minw:= -1, -1, len(s)
	for right < len(s) {
		// 移入的元素
		moin := s[right]
		// 扩大窗口
		right++
		// 进行窗口内数据的一系列更新
		if need[moin] != 0 {
			window[moin]++
			if window[moin] == need[moin] {
				valid++
			}
		}

		// 判断窗口是否要收缩
		for valid == len(need){
			// 更新最小覆盖子串
			if right -left+1 <
		}

	}
}

// m2. 字符串 'a' 转换成数字 来记录  -- leetcode gobook
// https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0076.Minimum-Window-Substring/
// 方法2很巧妙;
