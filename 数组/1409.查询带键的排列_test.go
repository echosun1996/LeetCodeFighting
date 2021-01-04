// 1409.查询带键的排列 - https://leetcode-cn.com/problems/queries-on-a-permutation-with-key
// By EchoSun
// 2021-01-02 21:16:02
package main

import (
	"reflect"
	"testing"
)

// 解法1 一种简单的模拟方法
// 使用数组 saveArray 存储位置。
// 数组下标 0 1 2 3 4 5 6 7 (代表对应的数)
// 存储元素 0 0 1 2 3 4 5 6 (代表位置）
func processQueries(queries []int, m int) []int {
	saveArray := make([]int, m+1)
	var ret []int
	// 对 saveArray 初始化，每一个位置代表每一个数，其上的值代表输出的结果
	for i := 1; i <= m; i++ {
		saveArray[i] = i - 1
	}

	// 选出一个待查找的元素
	for _, value := range queries {
		// 获取其所对应的位置
		ret = append(ret, saveArray[value])
		// 将剩余元素值小于对应位置值加一，即依次后移每一个数
		for i := 1; i <= m; i++ {
			if saveArray[i] < saveArray[value] {
				saveArray[i]++
			}
		}
		// 把当前这个数放在最前面
		saveArray[value] = 0
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1409(t *testing.T) {
	type args struct { // 测试样例的输入
		queries []int
		m       int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []int  // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{3, 1, 2, 1}, 5}, []int{2, 1, 2, 1}},
		{"test2", args{[]int{4, 1, 2, 2}, 4}, []int{3, 1, 2, 0}},
		{"test3", args{[]int{7, 5, 5, 8, 3}, 8}, []int{6, 5, 0, 7, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := processQueries(tt.args.queries, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
