// 41.缺失的第一个正数 - https://leetcode-cn.com/problems/first-missing-positive
// By EchoSun
// 2021-01-03 21:52:39
package main

import (
	"reflect"
	"testing"
)

// 解法1 原地哈希
// 即使用原始数组作为存储的空间来保存，即把1放在下标为0的位置上
func firstMissingPositive(nums []int) int {
	for _, value := range nums {
		for value > 0 && value <= len(nums) && nums[value-1] != value {
			temp := nums[value-1]
			nums[value-1] = value
			value = temp
		}
	}
	ix := 0
	for _, value := range nums {
		if value != ix+1 {
			break
		}
		ix++
	}
	return ix + 1
}

/////////////

// 测试函数

func TestProblem41(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{1, 2, 0}}, 3},
		{"test2", args{[]int{3, 4, -1, 1}}, 2},
		{"test3", args{[]int{7, 8, 9, 11, 12}}, 1},
		{"test4", args{[]int{}}, 1},
		{"test5", args{[]int{1}}, 2},
		{"test6", args{[]int{2, 1}}, 3},
		{"test7", args{[]int{0, 1, 2}}, 3},
		{"test8", args{[]int{3, 1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstMissingPositive(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
