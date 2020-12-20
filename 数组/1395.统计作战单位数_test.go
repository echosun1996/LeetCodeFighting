package main

import (
	"reflect"
	"testing"
)

// 解法1 暴力模拟，时间复杂度O(n3)
//func numTeams(rating []int) int {
//	var ret int
//	for i := 0; i < len(rating)-2; i++ {
//		for j := i + 1; j < len(rating)-1; j++ {
//			for k := j + 1; k < len(rating); k++ {
//				if rating[i] > rating[j] && rating[j] > rating[k] || rating[i] < rating[j] && rating[j] < rating[k] {
//					ret++
//				}
//			}
//		}
//	}
//	return ret
//}

// 解法2 使用三元组的思想解决本题。关键点在“中间元素”，时间复杂度为O(n2)
// 即最后的结果 = 递减三元组数 + 递增三元组数
// 其中，
// 递减三元组数 = （"左侧 > 中间元素" 的个数）* （"右侧 < 中间元素" 的个数 ）
// 递增三元组数 = （"左侧 < 中间元素" 的个数）* （"右侧 > 中间元素" 的个数 ）
func numTeams(rating []int) int {
	var ret int
	for i := 1; i < len(rating)-1; i++ { // 中间元素的下标
		leftLarger, rightLarger, leftSmaller, rightSmaller := 0, 0, 0, 0
		for j := 0; j < len(rating); j++ {
			if i == j {
				continue
			}
			if j < i {
				if rating[j] > rating[i] {
					leftLarger++
				} else if rating[j] < rating[i] {
					leftSmaller++
				}
			} else {
				if rating[j] > rating[i] {
					rightLarger++
				} else if rating[j] < rating[i] {
					rightSmaller++
				}
			}
		}
		ret += leftLarger*rightSmaller + leftSmaller*rightLarger
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1395(t *testing.T) {
	type args struct { // 测试样例的输入
		input []int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{2, 5, 3, 4, 1}}, 3},
		{"test2", args{[]int{2, 1, 3}}, 0},
		{"test3", args{[]int{1, 2, 3, 4}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numTeams(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
