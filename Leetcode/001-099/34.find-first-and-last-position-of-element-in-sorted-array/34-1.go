package _4_find_first_and_last_position_of_element_in_sorted_array

// 二分查找  左右两侧

// 二分查找  左右两侧
func searchRange(nums []int, target int) []int {
	return []int{left_bound(nums, target), right_bound(nums, target)}
}

func left_bound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			right = mid - 1
		}
	}
	// 判断 target 是否存在于 nums 中
	// 此时 target 比所有数都大，返回 -1
	if left == len(nums) {
		return -1
	}
	if nums[left] == target {
		return left
	} else {
		return -1
	}
}

func right_bound(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		} else if nums[mid] == target {
			left = mid + 1
		}
	}

	// 检查left-1 是否越界
	if left-1 < 0 {
		return -1
	}
	if nums[left-1] == target {
		return left - 1
	} else {
		return -1
	}
}
