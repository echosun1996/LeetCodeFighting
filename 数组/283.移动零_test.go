// 283 - 移动零
// By EchoSun
// 2020-12-21 16:49:23
package main

import (
	"reflect"
	"testing"
)

// 解法1
func moveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			nums[left] = nums[right]
			left++
		}
	}
	for ; left < len(nums); left++ {
		nums[left] = 0
	}
}

/////////////

// 测试函数

func TestProblem283(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []int  // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{0, 1, 0, 3, 12}}, []int{1, 3, 12, 0, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.input)
			if got := tt.args.input; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
