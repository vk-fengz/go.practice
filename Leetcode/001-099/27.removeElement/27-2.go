package main

// 快慢指针的方式
func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
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
