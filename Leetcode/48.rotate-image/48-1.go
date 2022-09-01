package _8_rotate_image

func rotate(matrix [][]int) {
	lenth := len(matrix)
	// 沿着对角线 翻转
	for i := 0; i < lenth; i++ {
		for j := i; j < lenth; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
		// 反转二维矩阵的每一行
		matrix[i] = reverse(matrix[i])
	}

}

func reverse(matrix []int) []int {
	lp, rp := 0, len(matrix)-1
	for lp < rp {
		matrix[lp], matrix[rp] = matrix[rp], matrix[lp]
		lp++
		rp--
	}
	return matrix
}
