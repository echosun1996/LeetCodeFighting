// 1365 - 有多少小于当前数字的数字
// By EchoSun
// 2020-12-21 15:22:22
package main

import (
	"reflect"
	"testing"
)

// 解法1 频次数组解决
func smallerNumbersThanCurrent(nums []int) []int {
	var ret []int
	var temp [101]int
	for _, value := range nums {
		temp[value]++
	}
	for _, value := range nums {
		sumTemp := 0
		for i := 0; i <= value-1; i++ {
			sumTemp += temp[i]
		}
		ret = append(ret, sumTemp)
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1365(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []int  // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{8, 1, 2, 2, 3}}, []int{4, 0, 1, 1, 3}},
		{"test2", args{[]int{6, 5, 4, 8}}, []int{2, 1, 0, 3}},
		{"test3", args{[]int{7, 7, 7, 7}}, []int{0, 0, 0, 0}},
		{"test4", args{[]int{5, 0, 10, 0, 10, 6}}, []int{2, 0, 4, 0, 4, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smallerNumbersThanCurrent(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
