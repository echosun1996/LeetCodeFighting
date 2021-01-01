// 1572.矩阵对角线元素的和 - https://leetcode-cn.com/problems/matrix-diagonal-sum
// By EchoSun
// 2020-12-30 15:20:53
package main

import (
	"reflect"
	"testing"
)

// 解法1
func diagonalSum(mat [][]int) int {
	//print(6 / 2)
	n := len(mat[0])
	println(n)
	ret := 0
	for i := 0; i <= n/2; i++ {
		if i == n-i {
			ret += mat[i][i]
		} else {
			ret += mat[i][i] + mat[i][n-i] + mat[n-i][i] + mat[n-i][n-i]
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1572(t *testing.T) {
	type args struct { // 测试样例的输入
		input [][]int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}}, 25},
		{"test2", args{[][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}}, 8},
		{"test3", args{[][]int{{5}}}, 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diagonalSum(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
