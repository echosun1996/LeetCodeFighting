// 1572.矩阵对角线元素的和 - https://leetcode-cn.com/problems/matrix-diagonal-sum
// By EchoSun
// 2020-12-30 15:20:53
package main

import (
	"reflect"
	"testing"
)

// 解法1 从外层开始，逐渐向内层循环，外层都是四个数相加，而最内层需要注意区分矩阵大小为奇数还是偶数，偶数依旧是四个数相加，奇数只加中心一个数。
func diagonalSum(mat [][]int) int {
	n := len(mat[0])
	ret := 0
	// 大小为3的矩阵，循环到下标为1；大小为4的矩阵，循环到下标也为1
	for i := 0; i <= (n-1)/2; i++ {
		// 如果矩阵大小是奇数（比如3），且循环到最后（即矩阵正中间位置），则该元素只添加一次。
		if i == n/2 {
			ret += mat[i][i]
		} else {
			ret += mat[i][i] + mat[i][n-1-i] + mat[n-1-i][i] + mat[n-1-i][n-1-i]
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
