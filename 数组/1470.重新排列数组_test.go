package main

import (
	"reflect"
	"testing"
)

// 解法1 对半个输入数组执行一次循环，追加两个元素
func shuffle(nums []int, n int) []int {
	var ret []int
	for i := 0; i < n; i++ {
		ret = append(ret, nums[i], nums[n+i])
		//ret = append(ret, nums[n+i])
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1470(t *testing.T) {
	type args struct {
		nums []int // 测试样例的输入
		n    int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []int  // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{2, 5, 1, 3, 4, 7}, 3}, []int{2, 3, 5, 4, 1, 7}},
		{"test2", args{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4}, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{"test3", args{[]int{1, 1, 2, 2}, 2}, []int{1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shuffle(tt.args.nums, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
