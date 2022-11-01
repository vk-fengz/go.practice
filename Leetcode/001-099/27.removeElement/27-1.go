// 移除数组中的指定元素
package main

import "fmt"

func main() {
    nums := []int{3, 2, 2,3}
    val := 3
    fmt.Println(removeElement(nums, val))
    fmt.Println(nums)
}

// 手撕
func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }
    j := len(nums)
    for i, _ := range nums {
        for j > i {
            if nums[j-1] == val {
                j--
            } else {
                break
            }
        }
        if nums[i] == val && i <= j-1 {
            nums[i], nums[j-1] = nums[j-1], nums[i]
            j--
        }

    }
    return j
}
