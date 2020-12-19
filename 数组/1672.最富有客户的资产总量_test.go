package main

import (
	"reflect"
	"testing"
)

// 解法1 输出最大的行和
func maximumWealth(accounts [][]int) int {
	maxResult := -1
	for _, valueX := range accounts {
		sumTemp := 0
		for _, valueY := range valueX {
			sumTemp += valueY
		}
		if maxResult < sumTemp {
			maxResult = sumTemp
		}
	}
	return maxResult
}

/////////////

// 测试函数

func TestProblem1672(t *testing.T) {
	type args struct {
		accounts [][]int // 测试样例的输入
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[][]int{{1, 2, 3}, {3, 2, 1}}}, 6},
		{"test2", args{[][]int{{1, 5}, {7, 3}, {3, 5}}}, 10},
		{"test3", args{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}}, 17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumWealth(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
				//Errorf 断言失败的时候，标识测试失败，打印出必要的信息，但测试继续，相当于 Logf + Fail
				//Fatalf 断言失败的时候，标识测试失败，打印出必要的信息，但中断测试，相当于 Logf + FailNow
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
