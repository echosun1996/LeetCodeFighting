// 1295.统计位数为偶数的数字 - https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits
// By EchoSun
// 2020-12-29 20:27:10
package main

import (
	"math"
	"reflect"
	"testing"
)

// 解法1 我发现的一个直接算位数的公式
// 位数 = 向下取整( log10(x) + 1 )  要求 x >= 1，刚好满足题目要求
// 当然，这一题也可以采取转字符串求长度、不断除2等各种方法解决。
func findNumbers(nums []int) int {
	ret := 0
	for _, value := range nums {
		if int(math.Floor(math.Log10(float64(value))+1))%2 == 0 {
			ret++
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1295(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{12, 345, 2, 6, 7896}}, 2},
		{"test2", args{[]int{555, 901, 482, 1771}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findNumbers(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
