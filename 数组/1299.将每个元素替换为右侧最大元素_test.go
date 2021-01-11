// 1299.将每个元素替换为右侧最大元素 - https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side
// By EchoSun
// 2021-01-10 10:18:50
package main

import (
	"reflect"
	"testing"
)

// 解法1 先找最大，找到最大以后用最大值来追加到新数组中。
// 只要外层循环没有到最大值对应的下标，就不用再找最大。
// 最后再追加个-1即可
//leetcode submit region begin(Prohibit modification and deletion)
func replaceElements(arr []int) []int {
	max := -1
	maxI := -1
	var ret []int
	for i := 0; i < len(arr)-1; i++ {
		// 找最大
		if i >= maxI {
			max = -1
			for j := i + 1; j < len(arr); j++ {
				if max < arr[j] {
					max = arr[j]
					maxI = j
				}
			}
		}
		// 用 max 替换当前元素
		ret = append(ret, max)
	}
	// 最后一个元素，用 -1 替换。
	return append(ret, -1)
}

//leetcode submit region end(Prohibit modification and deletion)

/////////////

// 测试函数

func TestProblem1299(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []int  // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{17, 18, 5, 4, 6, 1}}, []int{18, 6, 6, 6, 1, -1}},
		{"test2", args{[]int{400}}, []int{-1}},
		{"test3", args{[]int{6, 6, 6, 6}}, []int{6, 6, 6, -1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceElements(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
