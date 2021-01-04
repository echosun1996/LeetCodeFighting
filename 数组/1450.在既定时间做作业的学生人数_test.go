// 1450.在既定时间做作业的学生人数 - https://leetcode-cn.com/problems/number-of-students-doing-homework-at-a-given-time
// By EchoSun
// 2021-01-04 20:24:37
package main

import (
	"reflect"
	"testing"
)

// 解法1 这一题过于简单了，只要判断 queryTime 在开始和结束之间就把结果加一。
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	ret := 0
	for i := 0; i < len(startTime); i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			ret++
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1450(t *testing.T) {
	type args struct { // 测试样例的输入
		startTime []int
		endTime   []int
		queryTime int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{1, 2, 3}, []int{3, 2, 7}, 4}, 1},
		{"test2", args{[]int{4}, []int{4}, 4}, 1},
		{"test3", args{[]int{4}, []int{4}, 5}, 0},
		{"test4", args{[]int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7}, 0},
		{"test5", args{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := busyStudent(tt.args.startTime, tt.args.endTime, tt.args.queryTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
