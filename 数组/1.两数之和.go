package main

import "fmt"

// 解法1
// 时间复杂度：两层 for 循环，O（n²）
// 空间复杂度：O（1）
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}


func main() {
	nums := []int{2, 7, 11, 15}
	output := twoSum(nums, 9)
	fmt.Printf("%v", output)
}
