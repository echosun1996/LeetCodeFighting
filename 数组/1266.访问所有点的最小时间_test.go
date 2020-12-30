// 1266.访问所有点的最小时间 - https://leetcode-cn.com/problems/minimum-time-visiting-all-points
// By EchoSun
// 2020-12-26 09:57:22
package main

import (
	"math"
	"reflect"
	"testing"
)

// 解法1 每次走一个点，步数刚好等于 X或Y 差值的最大值
func minTimeToVisitAllPoints(points [][]int) int {
	old := []int{points[0][0], points[0][1]}
	ret := 0
	for i := 1; i < len(points); i++ {
		max := int(math.Max(math.Abs(float64(points[i][0]-old[0])), math.Abs(float64(points[i][1]-old[1]))))
		ret += max
		old = points[i]
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1266(t *testing.T) {
	type args struct { // 测试样例的输入
		input [][]int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[][]int{{1, 1}, {3, 4}, {-1, 0}}}, 7},
		{"test2", args{[][]int{{3, 2}, {-2, 2}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minTimeToVisitAllPoints(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
