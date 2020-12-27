// 80.删除排序数组中的重复项 II - https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii
// By EchoSun
// 2020-12-26 08:47:05
package main_test

import (
	"reflect"
	"testing"
)

// 解法1 题目要求处理一个数组，使得处理后的数组中重复元素的个数小于等于2
// 使用快慢两个指针以及一个重复标志变量解决。
// 比较快指针和快指针的前一个元素是否相同，符合重复次数小于等于2的情况才复制到慢指针上。
func removeDuplicates(nums []int) int {
	pardon := false
	fast, slow := 1, 1 // 快慢指针
	for fast < len(nums) {
		if nums[fast] == nums[fast-1] && !pardon { // 重复出现，但重复次数小于等于2
			nums[slow] = nums[fast]
			slow++
			pardon = true
		} else if nums[fast] != nums[fast-1] { // 没有重复出现
			nums[slow] = nums[fast]
			slow++
			pardon = false
		}
		fast++
	}
	return slow
}

/////////////

// 测试函数

func TestProblem80(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{1, 1, 1, 2, 2, 3}}, 5},
		{"test2", args{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}}, 7},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
