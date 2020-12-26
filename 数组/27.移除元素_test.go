// 27.移除元素 - https://leetcode-cn.com/problems/remove-element
// By EchoSun
// 2020-12-21 17:04:55
package main

import (
	"reflect"
	"testing"
)

// 解法1 快慢指针法。
// 其中，快指针在后面寻找不为val的值，与前面的慢指针做替换。
// 最后移除元素后的数组末尾就是slow指针的位置
func removeElement(nums []int, val int) int {
	fast, slow := 0, 0
	for fast < len(nums) {
		if nums[fast] == val {
		} else {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}

/////////////

// 测试函数

func TestProblem27(t *testing.T) {
	type args struct { // 测试样例的输入
		nums []int
		val  int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{3, 2, 2, 3}, 3}, 2},
		{"test2", args{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.args.nums, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
