// 950.按递增顺序显示卡牌 - https://leetcode-cn.com/problems/reveal-cards-in-increasing-order
// By EchoSun
// 2021-01-10 10:37:28
package main

import (
	"reflect"
	"sort"
	"testing"
)

// 解法1 本题主要难在对题意的理解上。首先，输入的顺序不重要，输出的顺序很重要。
// 需要保证按照输出的顺序翻牌得到的整数是顺序递增的。
// 需要对题目要求进行反向处理，即可得到答案。
// 具体方法：
// 1. 对输入数组降序排序。
// 2. “返回数组”最后一个元素放首部，再在首部放一个“输入数组”的最大值，移除“输入数组”的最大值。
// 3. 直到“输入数组”为空
//leetcode submit region begin(Prohibit modification and deletion)
func deckRevealedIncreasing(deck []int) []int {
	var ret []int
	// 输入数组降序排序
	sort.Sort(sort.Reverse(sort.IntSlice(deck)))
	for len(deck) > 0 {
		if len(ret) > 0 {
			retLast := ret[len(ret)-1]
			ret = append([]int{retLast}, ret[:len(ret)-1]...)
		}
		ret = append([]int{deck[0]}, ret...)
		deck = deck[1:]
	}
	return ret
}

//leetcode submit region end(Prohibit modification and deletion)
/////////////

// 测试函数

func TestProblem950(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []int  // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{17, 13, 11, 2, 3, 5, 7}}, []int{2, 13, 3, 11, 5, 17, 7}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deckRevealedIncreasing(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
