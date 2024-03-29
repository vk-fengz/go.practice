package _83_move_zeroes

// 与27 类似

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	bound := removeElement(nums, 0)
	for bound < len(nums) {
		nums[bound] = 0
		bound++
	}
}

func removeElement(nums []int, val int) int {
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != val && slow <= fast {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		} else {
			fast++
		}
	}
	return slow
}
