package main

import (
	"reflect"
	"testing"
)

// 解法1 简答的异或运算
func xorOperation(n int, start int) int {
	var ret int
	for i := 0; i < n; i++ {
		nums := start + 2*i
		ret = ret ^ nums
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1486(t *testing.T) {
	type args struct { // 测试样例的输入
		n     int
		start int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{5, 0}, 8},
		{"test2", args{4, 3}, 8},
		{"test3", args{1, 7}, 7},
		{"test4", args{10, 5}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := xorOperation(tt.args.n, tt.args.start); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
