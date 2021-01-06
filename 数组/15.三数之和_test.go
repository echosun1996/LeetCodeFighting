// 15.三数之和 - https://leetcode-cn.com/problems/3sum
// By EchoSun
// 2021-01-05 20:00:47
package main

import (
	"reflect"
	"sort"
	"testing"
)

// 解法1 外层从头到尾循环一次，内部使用双指针,注意排除重复出现的情况。
func threeSum(nums []int) [][]int {
	var ret [][]int
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums) // 升序排序
	for ai := 0; ai < len(nums); ai++ {
		// 第一个数必须小于等于零
		if nums[ai] > 0 {
			break
		}
		// 当前数和前一个数相同，跳过
		if ai-1 >= 0 && nums[ai-1] == nums[ai] {
			continue
		}
		ci := len(nums) - 1
		for bi := ai + 1; bi < ci; {
			if nums[ai]+nums[bi]+nums[ci] == 0 {
				ret = append(ret, []int{nums[ai], nums[bi], nums[ci]})
				// 保证新的b与旧的不相同
				oldB := nums[bi]
				for bi < ci && oldB == nums[bi] {
					bi++
				}
				// 保证新的c与旧的不相同
				oldC := nums[ci]
				for ci > bi && oldC == nums[ci] {
					ci--
				}
			} else if nums[ai]+nums[bi]+nums[ci] > 0 {
				// 保证新的c与旧的不相同
				oldC := nums[ci]
				for ci > bi && oldC == nums[ci] {
					ci--
				}
			} else {
				// 保证新的b与旧的不相同
				oldB := nums[bi]
				for bi < ci && oldB == nums[bi] {
					bi++
				}
			}
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem15(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string  // 测试样例名称
		args args    // 测试样例的参数
		want [][]int // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{-1, 0, 1, 2, -1, -4}}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
		{"test2", args{[]int{0, 0, 0}}, [][]int{{0, 0, 0}}},
		{"test3", args{[]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}}, [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}},
		{"test4", args{[]int{-1, 0, 1, 0}}, [][]int{{-1, 0, 1}}},
		{"test5", args{[]int{0, 0, 0, 0}}, [][]int{{0, 0, 0}}},
		{"test6", args{[]int{-4, -2, 1, -5, -4, -4, 4, -2, 0, 4, 0, -2, 3, 1, -5, 0}}, [][]int{{-5, 1, 4}, {-4, 0, 4}, {-4, 1, 3}, {-2, -2, 4}, {-2, 1, 1}, {0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSum(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
