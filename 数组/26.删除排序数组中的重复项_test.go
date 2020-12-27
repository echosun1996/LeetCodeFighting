// 26.删除排序数组中的重复项 - https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
// By EchoSun
// 2020-12-27 10:39:37
package main

import (
	"reflect"
	"testing"
)

// 解法1 快慢指针，快指针遇到重复元素直接跳过即可。
func removeDuplicates(nums []int) int {
	fast, slow := 1, 1 // 快慢指针
	for fast < len(nums) {
		if nums[fast] != nums[fast-1] { // 没有重复出现
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

/////////////

// 测试函数

func TestProblem26(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{1, 1, 2}}, 2},
		{"test2", args{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
