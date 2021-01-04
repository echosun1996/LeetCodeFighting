// 32.最长有效括号 - https://leetcode-cn.com/problems/longest-valid-parentheses
// By EchoSun
// 2021-01-03 21:06:27
package 字符串

import (
	"reflect"
	"testing"
)

// 解法1 最好的方法是采用动态规划的方法
// 思路相当简单，即用两个变量存左右括号的个数，若左=右，则比较是否是最大。
// 注意！采用本方法需要向左、向右各循环一次，否则像(()这样的永远都不会出现左=右
// 本题最简单的做法就是用栈来模拟括号匹配，然后如果匹配，则将对应数组的位置赋1.最后看这个数组最长的连续1
func longestValidParentheses(s string) int {
	ret := 0
	left := 0
	right := 0

	// 向右循环
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if left*2 > ret {
				ret = left * 2
			}
		} else if left < right {
			left, right = 0, 0
		}
	}
	// 循环变量初始化
	left, right = 0, 0
	// 向左循环
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if left*2 > ret {
				ret = left * 2
			}
		} else if left > right {
			left, right = 0, 0
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem32(t *testing.T) {
	type args struct { // 测试样例的输入
		input string
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{"(()"}, 2},
		{"test2", args{")()())"}, 4},
		{"test3", args{"()(()"}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
