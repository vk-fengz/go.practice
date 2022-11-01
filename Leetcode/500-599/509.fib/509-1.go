package _09_fib

// 动态规划 状态转换方程

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}

	dpi, dpi1, dpi2 := 0, 1, 0
	for i := 2; i <= n; i++ {
		dpi = dpi1 + dpi2
		dpi2 = dpi1
		dpi1 = dpi
	}
	return dpi
}
