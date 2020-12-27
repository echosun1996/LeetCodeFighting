// 258.各位相加 - https://leetcode-cn.com/problems/add-digits
// By EchoSun
// 2020-12-27 08:24:26
package 数学

import (
	"reflect"
	"testing"
)

// 解法1 题目要求不使用循环或递归完成，因此考虑使用数学公式解决。
// 这是一个数根问题。维基页面: https://en.wikipedia.org/wiki/Digital_root 和 https://zh.wikipedia.org/wiki/數根
// 比如19的数根求法为： 19 => 1+9 = 10 => 1+0 = 1
// 数根可用来检测计算正确性（两数字的和的数根等于两数字分别的数根的和）和整除性（如果数根能被3或9整除，则原来的数也能被3或9整除）等。
//
// 数根计算公式 digital toot = 1 + ((num - 1) % 9)
// 推导详情页面：https://blog.csdn.net/ling_wang/article/details/80559087
//
// 至于推导的公式，我认为下面的例子可能更加直观
// 原数 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30
// 数根 1 2 3 4 5 6 7 8 9  1  2  3  4  5  6  7  8  9  1  2  3  4  5  6  7  8  9  1  2  3

// 原数: 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30
// 偏移: 0 1 2 3 4 5 6 7 8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29
// 取余: 0 1 2 3 4 5 6 7 8  0  1  2  3  4  5  6  7  8  0  1  2  3  4  5  6  7  8  0  1  2
// 数根: 1 2 3 4 5 6 7 8 9  1  2  3  4  5  6  7  8  9  1  2  3  4  5  6  7  8  9  1  2  3
// 可以看做是9个一组
func addDigits(num int) int {
	return (num-1)%9 + 1
}

/////////////

// 测试函数

func TestProblem258(t *testing.T) {
	type args struct { // 测试样例的输入
		input int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{38}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
