package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5}
	target := 9
	s1 := twoSum1(nums, target)

	fmt.Println(s1)
}
func twoSum1(nums []int, target int) []int {
	for i, x := range nums {
		for j := len(nums) - 1; j > i; j-- {
			if nums[j] == target-x {
				return []int{i, j}
			}
		}
	}
	return nil
}

// leetcode
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		fmt.Println(i)
		fmt.Println(hashTable)
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil
}

func twoSum3(nums []int, target int) []int {
    m := make(map[int]int)
    for index, val := range nums {
        if preIndex, ok := m[target-val]; ok {
            return []int{preIndex, index}
        } else {
            m[val] = index
        }
    }
    return []int{}
}

func twoSum2(nums []int, target int) []int {
	for i, x := range nums {
		for j := i + 1; j < len(nums); j++ {
			if x+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}
