package main

import (
	"reflect"
	"testing"
)

// 解法1
type SubrectangleQueries struct {
	rectangle [][]int
	update [][]int
}


func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

// UpdateSubrectangle 用 newValue 更新以 (row1,col1) 为左上角且以 (row2,col2) 为右下角的子矩形。
func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {
	this.update = append(this.update,[]int{row1,col1,row2,col2,newValue} )
}

// GetValue 返回矩形中坐标 (row,col) 的当前值。
func (this *SubrectangleQueries) GetValue(row int, col int) int {
	retValue := this.rectangle[row][col]
	for _, value := range this.update{
		if row <= value[2] && row >= value[0] && col <= value[3] && row >= value[1]{
			retValue+= value[4]
		}
		//println(value)
	}
	return retValue
}


/////////////

// 测试函数

func Apply(f interface{}, args []interface{})([]reflect.Value){
	fun := reflect.ValueOf(f)
	in := make([]reflect.Value, len(args))
	for k,param := range args{
		in[k] = reflect.ValueOf(param)
	}
	r := fun.Call(in)
	return r

}
func TestProblem1476(t *testing.T) {



	input := [][]int{{1,2,1},{4,3,4},{3,2,1},{1,1,1}}
	obj := Constructor(input)

	retValue := obj.GetValue(0,2)
	if retValue != 1{
		t.Fatalf("2333")
	}
	obj.UpdateSubrectangle(0,0,3,2,5)
	retValue = obj.GetValue(0,2)
	if retValue != 5{
		t.Fatalf("2333")
	}
	//type args struct {
	//	accounts [][]int // 测试样例的输入
	//}
	//type test struct {
	//	name string // 测试样例名称
	//	args args   // 测试样例的参数
	//	want int    // 测试样例的返回值
	//}
	//tests := []test{
	//	{"test1", args{[][]int{{1, 2, 3}, {3, 2, 1}}}, 6},
	//	{"test2", args{[][]int{{1, 5}, {7, 3}, {3, 5}}}, 10},
	//	{"test3", args{[][]int{{2, 8, 7}, {7, 1, 3}, {1, 9, 5}}}, 17},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := maximumWealth(tt.args.accounts); !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
	//		}
	//	})
	//}
}
