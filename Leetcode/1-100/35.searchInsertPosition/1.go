// 给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。
// 如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
// 输入: nums = [1,3,5,6], target = 5
// 输出: 2

package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

// leetcode
func searchInsertlc(nums []int, target int) int {
    n := len(nums)
    left, right := 0, n-1
    ans := n
    for left <= right {
        mid := (right-left)>>1 + left
        fmt.Println("midlc: ",mid)
        if target <= nums[mid] {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return ans
}

func main() {
    // var input string
    // Scanf(&input) // string

    nums := []int{1, 3, 5,6}
    target := 0
    fmt.Println(searchInsertlc(nums, target))
    fmt.Println(searchInsert(nums, target))
}

// 手撕
func searchInsert(nums []int, target int) int {
    n := len(nums)
    left, right := 0, n-1
    ans := 0
    for left <= right {
        mid := left + (right-left)>>1
        if target <= nums[mid] {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return ans
}

// 读取输入的数组
// 输入: *string
// 输出: 输入变量 *string
func Scanf(input *string) {
    reader := bufio.NewReader(os.Stdin)
    data, _, _ := reader.ReadLine()
    *input = string(data)
    fmt.Printf("Input: %s", *input)
}

// 字符串 转换为 数字
func StrConv(str *string) int {
    stringtoint, _ := strconv.Atoi(*str)
    fmt.Println("转换后数字: ", stringtoint)
    return stringtoint
}

// 字符串 转 整型切片
func stringtointslice(str *string) []int {
    fields := strings.Fields(*str)
    sliceint := make([]int, 0)
    for index, _ := range fields {
        temp, _ := strconv.Atoi(fields[index]) // ParseInt 更强大
        sliceint = append(sliceint, temp)
    }
    return sliceint
}
