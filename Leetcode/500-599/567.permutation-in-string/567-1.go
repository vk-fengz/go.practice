package _67_permutation_in_string

func checkInclusion(s1 string, s2 string) bool {
	// 获得键值表 need: 目标值;
	need, window := make(map[byte]int), make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		need[s1[i]]++
	}
	// valid 验证满足与否

	// 记录最小覆盖子串的起始索引及 结束
	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		// 移入的元素
		moin := s2[right]
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
		for right-left >= len(s1) {
			// 更新最小覆盖子串
			if valid == len(need) {
				return true
			}

			// moveout 是移出的字符
			moveout := s2[left]
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
	return false
}
