// 面试题0804.幂集 - https://leetcode-cn.com/problems/power-set-lcci
// By EchoSun
// 2020-12-30 14:46:56
package main

import (
	"math"
	"reflect"
	"testing"
)

// 解法1 只要输出数组的所有组合即可，没必要注重顺序，因此本题的答案可以是无序的，只要每一种情况都出现即可。
// 本题采用的解法是用二进制数字表示每一个位在结果中是否出现，这个数字从0到 2^length -1，循环一次，即可得到答案。
func subsets(nums []int) [][]int {
	ret := [][]int{{}} // 默认数组中有一个空集
	length := len(nums)
	max := int(math.Pow(2, float64(length))) - 1
	for sequence := 1; sequence <= max; sequence++ { // 排序序列，使用二进制对应该位是否出现
		var temp []int
		// mask 用于从高到低依次取每一位，i 用于标记在数组上对应哪一位
		i := length - 1
		for mask := 1; mask <= int(math.Pow(2, float64(length-1))); mask = mask << 1 {
			if sequence&mask != 0 {
				temp = append(temp, nums[i])
			}
			i--
		}
		ret = append(ret, temp)
	}
	return ret
}

/////////////

// 测试函数

func TestProblem0804(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string  // 测试样例名称
		args args    // 测试样例的参数
		want [][]int // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{1, 2, 3}}, [][]int{{}, {3}, {2}, {3, 2}, {1}, {3, 1}, {2, 1}, {3, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subsets(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
