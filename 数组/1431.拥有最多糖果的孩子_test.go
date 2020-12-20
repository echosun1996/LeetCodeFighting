package main

import (
	"reflect"
	"testing"
)

// 解法1 直接求一个数组最大值 max ,若 candies[i]+extraCandies >= max 则为 true
func kidsWithCandies(candies []int, extraCandies int) []bool {
	var ret []bool
	var max int
	for _, value := range candies {
		if value > max {
			max = value
		}
	}
	for _, value := range candies {
		if value+extraCandies >= max {
			ret = append(ret, true)
		} else {
			ret = append(ret, false)
		}
	}
	return ret
}

/////////////

// 测试函数

func TestProblem1431(t *testing.T) {
	type args struct {
		candies      []int // 测试样例的输入
		extraCandies int
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want []bool // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]int{2, 3, 5, 1, 3}, 3}, []bool{true, true, true, false, true}},
		{"test2", args{[]int{4, 2, 1, 1, 2}, 1}, []bool{true, false, false, false, false}},
		{"test3", args{[]int{12, 1, 12}, 10}, []bool{true, false, true}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kidsWithCandies(tt.args.candies, tt.args.extraCandies); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
