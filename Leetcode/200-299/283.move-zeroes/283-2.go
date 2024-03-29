package _83_move_zeroes

func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}
	slow, fast := 0, 0
	for fast < len(nums) {
		if nums[fast] != 0 && slow <= fast {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		} else {
			fast++
		}
	}
}
