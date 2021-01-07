// 59.螺旋矩阵 II - https://leetcode-cn.com/problems/spiral-matrix-ii
// By EchoSun
// 2021-01-04 20:47:24
package main

import (
	"reflect"
	"testing"
)

// 解法1 简单模拟，留意边界即可
func generateMatrix(n int) [][]int {
	// 创建一个空的二维数组
	ret := make([][]int, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
	}
	// i表示行 j表示列
	// 循环模式。0向右 1向下 2向左 3向上
	mode := 0
	i, j := 0, 0
	maxI, maxJ := n-1, n-1
	minI, minJ := -1, -1
	for count := 1; count <= n*n; count++ {
		if i == maxI && mode == 0 {
			minJ++
			mode = 1
		} else if j == maxJ && mode == 1 {
			maxI--
			mode = 2
		} else if i-1 == minI && mode == 2 {
			maxJ--
			mode = 3
		} else if j-1 == minJ && mode == 3 {
			minI++
			mode = 0
		}
		ret[j][i] = count
		if mode == 0 {
			i++
		} else if mode == 1 {
			j++
		} else if mode == 2 {
			i--
		} else if mode == 3 {
			j--
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem59(t *testing.T) {
	type args struct { // 测试样例的输入
		input int
	}
	type test struct {
		name string  // 测试样例名称
		args args    // 测试样例的参数
		want [][]int // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{3}, [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}},
		{"test2", args{1}, [][]int{{1}}},
		{"test3", args{0}, [][]int{{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateMatrix(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
