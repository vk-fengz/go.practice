// https://leetcode-cn.com/problems/search-in-rotated-sorted-array/
// 输入：nums = [4,5,6,7,0,1,2], target = 0
// 输出：4

package main

import (
    "fmt"
)

func main() {
    nums := []int{5, 1, 3}
    target := 4
    fmt.Print(search1(nums, target))
}

// leetcode官方 c++解法复现  [4,5,6,7,0,1,2]
// 0
func search(nums []int, target int) int {

	if 0 == len(nums) {
		return -1
	}
	if 1 == len(nums) && nums[0] == target {
		return 0
	}

	low, high := 0, len(nums)-1
	for low <= high {
		mid := low + (high-low)>>1
		// fmt.Println(mid)
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] {
			if nums[0] <= target && target < nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[len(nums)-1] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

// 硬撕
func search1(nums []int, target int) int {
    low, high := 0, len(nums)-1
    for low <= high {
        mid := low + (high-low)>>1
        if nums[mid] < target {
            if nums[low] <= nums[mid] { // mid target均在大数区间
                low = mid + 1
            } else if nums[mid] <= nums[high] && nums[low] <= target { // mid在小数区间, target在大数区间
                high = mid - 1
            } else if target <= nums[high] { // mid, target同在小数区间
                low = mid + 1
            } else {
                return -1
            }
        } else if target < nums[mid] {
            // mid 在大数区间, target在小数区间
            if nums[low] <= nums[mid] && target < nums[low] {
                low = mid + 1
            } else if nums[low] <= nums[mid] {
                high = mid - 1
            } else if nums[mid] <= nums[high] { // mid在小数区间,target在小数区间
                high = mid - 1
            } else {
                return -1
            }

        } else {
            return mid
        }
    }

    return -1
}