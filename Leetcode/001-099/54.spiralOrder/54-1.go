package _4_spiralOrder

func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	// [m][n]int  高度: m*n
	upper_bound, lower_bound, left_bound, right_bound := 0, m-1, 0, n-1
	res := []int{}

	// len(res) == m*n 遍历整个数组
	for len(res) < m*n {
		if upper_bound <= lower_bound {
			// 顶部 自左向右
			for j := left_bound; j <= right_bound; j++ {
				res = append(res, matrix[upper_bound][j])
			}
			// 上边界向下移动
			upper_bound++
		}

		// 右侧 自上而下
		if left_bound <= right_bound {
			for i := upper_bound; i <= lower_bound; i++ {
				res = append(res, matrix[i][right_bound])
			}
			// right_bound go to left_bound
			right_bound--
		}
		// 底部 自右向左
		if upper_bound <= lower_bound {
			for j := right_bound; j >= left_bound; j-- {
				res = append(res, matrix[lower_bound][j])
			}
			lower_bound--
		}

		// left  from lower_bound to upper_bound
		if left_bound <= right_bound {
			for i := lower_bound; i >= upper_bound; i-- {
				res = append(res, matrix[i][left_bound])
			}
			left_bound++
		}
	}

	return res
}
