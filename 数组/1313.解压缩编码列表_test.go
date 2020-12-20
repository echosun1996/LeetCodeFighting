package main

import (
	"reflect"
	"testing"
)

// 解法1 简单模拟,题意就是答案中需要出现 nums[2*i] 个 nums[2*i+1]] 的值，遍历一次，结果串起来
func decompressRLElist(nums []int) []int {
	var ret []int
	for i := 0; 2*i < len(nums); i += 1 {
		for j := nums[2*(i)]; j > 0; j-- {
			ret = append(ret, nums[2*(i)+1])
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1313(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []int  // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{1, 2, 3, 4}}, []int{2, 4, 4, 4}},
		{"test2", args{[]int{1, 1, 2, 3}}, []int{1, 3, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decompressRLElist(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
