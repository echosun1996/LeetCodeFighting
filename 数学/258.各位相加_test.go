// 258.各位相加 - https://leetcode-cn.com/problems/add-digits
// By EchoSun
// 2020-12-27 08:24:26
package main

import (
	"reflect"
	"testing"
)

// 解法1 题目要求不使用循环或递归完成，因此考虑使用数学公式解决。
// 这是一个数根问题。维基页面: https://en.wikipedia.org/wiki/Digital_root 和 https://zh.wikipedia.org/wiki/數根
// 比如19的数根求法为： 19 => 1+9 = 10 => 1+0 = 1
// 数根可用来检测计算正确性（两数字的和的数根等于两数字分别的数根的和）和整除性（如果数根能被3或9整除，则原来的数也能被3或9整除）等。
//
// 数根计算公式 digital root = 1 + ((num - 1) % 9)
// 推导详情页面：https://blog.csdn.net/ling_wang/article/details/80559087
//
//下面的例子更加直观的解释了dr(n)公式的推导过程
//原数 0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30
//数根 0 1 2 3 4 5 6 7 8 9  1  2  3  4  5  6  7  8  9  1  2  3  4  5  6  7  8  9  1  2  3
//观察上述数字可以发现，除了0以外，别的数都是九个一组排布的。因而，我们可以得到如“推导详情页面”中的dr(n)公式
//然后使用三个if语句就可以完成题目要求。
//
//然而更进一步，可以观察到如下的情况。
//     原数:  0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29 30
//   -1偏移: -1 0 1 2 3 4 5 6 7 8  9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25 26 27 28 29
//   %9取余: -1 0 1 2 3 4 5 6 7 8  0  1  2  3  4  5  6  7  8  0  1  2  3  4  5  6  7  8  0  1  2
// +1得数根:  0 1 2 3 4 5 6 7 8 9  1  2  3  4  5  6  7  8  9  1  2  3  4  5  6  7  8  9  1  2  3
//可以看做是9个一组，但是结果除了0的数根是0以外，不存在0，因此，我们可以先把数左移一位，再对9取余数，然后再恢复，
//这样能够避免结果中出现0的情况。非常巧妙。
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
