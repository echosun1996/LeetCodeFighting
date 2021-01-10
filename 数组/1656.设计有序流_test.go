// 1656.设计有序流 - https://leetcode-cn.com/problems/design-an-ordered-stream
// By EchoSun
// 2021-01-07 21:26:43
package main_test

import (
	"reflect"
	"testing"
)

// 解法1
type OrderedStream struct {
	stringList []string
	ptr        int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{stringList: make([]string, n+1), ptr: 1}
}

func (this *OrderedStream) Insert(id int, value string) []string {
	this.stringList[id-1] = value
	var ret []string
	for this.stringList[this.ptr-1] != "" {
		ret = append(ret, this.stringList[this.ptr-1])
		this.ptr++
	}
	return ret
}

/**
 * Your OrderedStream object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Insert(id,value);
 */

/////////////

// 测试函数

func TestProblem1656(t *testing.T) {
	type args struct { // 测试样例的输入
		id    int
		value string
	}
	type test struct {
		name string   // 测试样例名称
		args args     // 测试样例的参数
		want []string // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{3, "ccccc"}, []string{}},
		{"test2", args{1, "aaaaa"}, []string{"aaaaa"}},
		{"test3", args{2, "bbbbb"}, []string{"bbbbb", "ccccc"}},
		{"test4", args{5, "eeeee"}, []string{}},
		{"test5", args{4, "ddddd"}, []string{"ddddd", "eeeee"}},
	}
	//t.Run(tt.name, func(t *testing.T) {
	obj := Constructor(5)
	for _, tt := range tests {
		if got := obj.Insert(tt.args.id, tt.args.value); !reflect.DeepEqual(got, tt.want) {
			if len(got) == 0 && len(tt.want) == 0 {
				continue
			}
			t.Fatalf("[Wrong Answer]%v-运行结果:%v, 正确答案: %v", tt.name, got, tt.want)
		}
	}
}
