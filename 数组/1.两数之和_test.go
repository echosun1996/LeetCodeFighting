package main

import (
	"reflect"
	"testing"
)

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

// 测试函数
func TestProblem1(t *testing.T) {
	type args struct {
		nums   []int
		target int // 测试样例的输入
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []int  // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{2, 7, 11, 15}, 9}, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				//Errorf 断言失败的时候，标识测试失败，打印出必要的信息，但测试继续，相当于 Logf + Fail
				//Fatalf 断言失败的时候，标识测试失败，打印出必要的信息，但中断测试，相当于 Logf + FailNow
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
