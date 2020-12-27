// 80.删除排序数组中的重复项 II - https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array-ii
// By EchoSun
// 2020-12-26 08:47:05
package main

import (
	"reflect"
	"testing"
)

// 解法1
//leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(nums []int) int {
	fast, slow := 2, 1
	for fast < len(nums) {
		if fast != slow && nums[fast] == nums[slow] {
			fast++
			slow++
		}
	}
}

//leetcode submit region end(Prohibit modification and deletion)

/////////////

// 测试函数

func TestProblem80(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []int  // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{1, 2, 3, 4}}, []int{1, 3, 6, 10}},
		{"test2", args{[]int{1, 1, 1, 1, 1}}, []int{1, 2, 3, 4, 5}},
		{"test3", args{[]int{3, 1, 2, 10, 1}}, []int{3, 4, 6, 16, 17}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSum(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
