// 1588.所有奇数长度子数组的和 - https://leetcode-cn.com/problems/sum-of-all-odd-length-subarrays
// By EchoSun
// 2021-01-02 20:42:26
package main

import (
	"reflect"
	"testing"
)

// 解法1 求奇数长度的连续子数列的和，模拟法解决
// 注意，“子数组”定义为原数组中的一个“连续”子序列。
func sumOddLengthSubarrays(arr []int) int {
	ret := 0
	// i 表示基本位置，依次加一
	for i := 0; i < len(arr); i++ {
		// adder 表示每次横跨的子序列下标扩增的长度，从0到最大，每次加2
		for adder := 0; adder+i < len(arr); adder += 2 {
			// temp 表示子序列坐标，从 i 到 i+adder ,每次加1
			for temp := i; temp <= i+adder; temp++ {
				ret += arr[temp]
			}
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1588(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{1, 4, 2, 5, 3}}, 58},
		{"test2", args{[]int{1, 2}}, 3},
		{"test3", args{[]int{10, 11, 12}}, 66},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOddLengthSubarrays(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
