package main

import (
	"reflect"
	"testing"
)

// 解法1 简单的创建一个空数组，然后追加求和
func runningSum(nums []int) []int {
	var output []int
	temp := 0
	for _, value := range nums {
		temp += value
		output = append(output, temp)
	}
	return output
}

/////////////

// 测试函数

func TestProblem1480(t *testing.T) {
	type args struct {
		input []int // 测试样例的输入
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
