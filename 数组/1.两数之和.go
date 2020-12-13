package main

import "fmt"

// 解法1 两层 for 循环求和判断
// 时间复杂度：O(n2)
// 空间复杂度：O(1)
//func twoSum(nums []int, target int) []int {
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i]+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}

// 解法2 使用map存储了另一个数和它的下标
// 只需要循环一次就可以完成
// 时间复杂度：O(n)
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		_, ok := m[another]
		if ok {
			// m[another] 是排在前面的数
			return []int{m[another], i}
		}
		// 如果另一个数不在map中，则把当前这个数及下标存到map里
		m[nums[i]] = i
	}
	return nil
}

/////////////
// main 主函数
func main() {
	nums := []int{2, 7, 11, 15}
	output := twoSum(nums, 9)
	fmt.Printf("%v", output)
}
