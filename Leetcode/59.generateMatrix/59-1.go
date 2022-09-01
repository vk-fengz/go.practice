package _9_generateMatrix

func generateMatrix(n int) [][]int {

	// 构建二维动态矩阵 n*n
	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}
	// [m][n]int  高度: m*n
	upper_bound, lower_bound, left_bound, right_bound := 0, n-1, 0, n-1
	num := 1

	// len(res) == m*n 遍历整个数组
	for num <= n*n {
		if upper_bound <= lower_bound {
			// 顶部 自左向右
			for j := left_bound; j <= right_bound; j++ {
				matrix[upper_bound][j] = num
				num++
			}
			// 上边界向下移动
			upper_bound++
		}

		// 右侧 自上而下
		if left_bound <= right_bound {
			for i := upper_bound; i <= lower_bound; i++ {
				matrix[i][right_bound] = num
				num++
			}
			// right_bound go to left_bound
			right_bound--
		}
		// 底部 自右向左
		if upper_bound <= lower_bound {
			for j := right_bound; j >= left_bound; j-- {
				matrix[lower_bound][j] = num
				num++
			}
			lower_bound--
		}

		// left  from lower_bound to upper_bound
		if left_bound <= right_bound {
			for i := lower_bound; i >= upper_bound; i-- {
				matrix[i][left_bound] = num
				num++
			}
			left_bound++
		}
	}

	return matrix
}
