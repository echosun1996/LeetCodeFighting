// 1389 - 按既定顺序创建目标数组
// By EchoSun
// 2020-12-21 14:51:37
package main

import (
	"reflect"
	"testing"
)

// 解法1 切片操作。首先创建一个长度同 nums 的数组，每次添加，把 index[i] 向后移动一个，再插入
func createTargetArray(nums []int, index []int) []int {
	var ret = make([]int, len(nums)) // 最终数组的长度等于插入元素的个数，即nums的大小
	for i := 0; i < len(nums); i++ {
		position := index[i]
		copy(ret[position+1:], ret[position:]) // 使用切片将position向后移动一个位置。copy这个函数前面是目标位置，后面是源位置
		ret[position] = nums[i]
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1389(t *testing.T) {
	type args struct { // 测试样例的输入
		nums  []int
		index []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []int  // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{0, 1, 2, 3, 4}, []int{0, 1, 2, 2, 1}}, []int{0, 4, 1, 3, 2}},
		{"test2", args{[]int{1, 2, 3, 4, 0}, []int{0, 1, 2, 3, 0}}, []int{0, 1, 2, 3, 4}},
		{"test3", args{[]int{1}, []int{0}}, []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := createTargetArray(tt.args.nums, tt.args.index); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
