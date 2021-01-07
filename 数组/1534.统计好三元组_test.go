// 1534.统计好三元组 - https://leetcode-cn.com/problems/count-good-triplets
// By EchoSun
// 2021-01-07 21:12:21
package main

import (
	"math"
	"reflect"
	"testing"
)

// 解法1 暴力遍历
func countGoodTriplets(arr []int, a int, b int, c int) int {
	ret := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			for k := j + 1; k < len(arr); k++ {
				if math.Abs(float64(arr[i]-arr[j])) <= float64(a) &&
					math.Abs(float64(arr[j]-arr[k])) <= float64(b) &&
					math.Abs(float64(arr[i]-arr[k])) <= float64(c) {
					ret++
				}

			}
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1534(t *testing.T) {
	type args struct { // 测试样例的输入
		arr []int
		a   int
		b   int
		c   int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{3, 0, 1, 1, 9, 7}, 7, 2, 3}, 4},
		{"test2", args{[]int{1, 1, 2, 2, 3}, 0, 0, 1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGoodTriplets(tt.args.arr, tt.args.a, tt.args.b, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
