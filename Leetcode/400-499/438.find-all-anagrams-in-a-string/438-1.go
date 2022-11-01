package _38_find_all_anagrams_in_a_string

func findAnagrams(s string, p string) []int {
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(p); i++ {
		need[p[i]]++
	}
	// valid 验证满足与否
	// 记录最小覆盖子串的起始索引及 结束
	left, right := 0, 0
	valid, result := 0, []int{}
	for right < len(s) {
		// 移入的元素
		moin := s[right]
		// 扩大窗口
		right++
		// 进行窗口内数据的一系列更新
		if _, ok := need[moin]; ok {
			window[moin]++
			if window[moin] == need[moin] {
				valid++
			}
		}

		// 判断窗口是否要收缩
		for right-left >= len(p) {
			// 更新最小覆盖子串
			if valid == len(need) {
				result = append(result, left)
			}

			// moveout 是移出的字符
			moveout := s[left]
			left++
			// 对窗口数据进行更新
			if _, ok := need[moveout]; ok {
				if window[moveout] == need[moveout] {
					valid--
				}
				window[moveout]--
			}
		}

	}
	// 未找到符合条件的子串
	return result
}
