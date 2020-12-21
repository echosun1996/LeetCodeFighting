package main

import (
	"reflect"
	"testing"
)

// 解法1 两循环变量前后比较，这是一个空间复杂度低的方法，但时间复杂度n2
//func numIdenticalPairs(nums []int) int {
//	var ret int
//	for i := 0; i < len(nums)-1; i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[i] == nums[j] {
//				ret++
//			}
//		}
//	}
//	return ret
//}

// 解法2 由于要求是 nums[i] == nums[j]，则答案跟相同数字的个数有关，分析得到带入排列组合公式可以解决。
// 这里使用了一个大小为100的数组（由于1 <= nums.length <= 100）作为频次表，temp[i]表示第i个数出现的次数。
// 而且组合公式C^2_n = n(n-1) / 2
func numIdenticalPairs(nums []int) int {
	var temp [100]int
	var ret int
	for _, value := range nums {
		temp[value-1]++ // 这里是为了避免 nums.length = 100 的时候出现下标越界的情况。
	}
	for _, value := range temp {
		if value >= 2 {
			ret += value * (value - 1) / 2
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1512(t *testing.T) {
	type args struct {
		input []int // 测试样例的输入
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want int    // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{1, 2, 3, 1, 1, 3}}, 4},
		{"test2", args{[]int{1, 1, 1, 1}}, 6},
		{"test3", args{[]int{1, 2, 3}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIdenticalPairs(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
