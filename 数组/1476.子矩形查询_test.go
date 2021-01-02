// 1476.子矩形查询 - https://leetcode-cn.com/problems/subrectangle-queries
// By EchoSun

package main

import (
	"testing"
)

// 解法1 用一个新矩阵update去覆盖旧的矩阵rectangle
type SubrectangleQueries struct {
	rectangle [][]int
	update    [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

// UpdateSubrectangle 用 newValue 更新以 (row1,col1) 为左上角且以 (row2,col2) 为右下角的子矩形。
func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	this.update = append(this.update, []int{row1, col1, row2, col2, newValue})
}

// GetValue 返回矩形中坐标 (row,col) 的当前值。
func (this *SubrectangleQueries) GetValue(row int, col int) int {
	retValue := this.rectangle[row][col]
	for _, value := range this.update {
		if row <= value[2] && row >= value[0] && col <= value[3] && col >= value[1] {
			retValue = value[4]
		}
	}
	return retValue
}

/////////////

// 测试函数

func TestProblem1476(t *testing.T) {
	// 测试1
	input := [][]int{{1, 2, 1}, {4, 3, 4}, {3, 2, 1}, {1, 1, 1}}
	obj := Constructor(input)

	retValue := obj.GetValue(0, 2)
	if retValue != 1 {
		t.Fatalf("Error!")
	}
	obj.UpdateSubrectangle(0, 0, 3, 2, 5)
	retValue = obj.GetValue(0, 2)
	if retValue != 5 {
		t.Fatalf("Error!")
	}

	// 测试2
	input = [][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}}
	obj = Constructor(input)
	retValue = obj.GetValue(0, 0)
	if retValue != 1 {
		t.Fatalf("Error!")
	}
	obj.UpdateSubrectangle(0, 0, 2, 2, 100)
	retValue = obj.GetValue(2, 2)
	if retValue != 100 {
		t.Fatalf("Error!")
	}
	obj.UpdateSubrectangle(1, 1, 2, 2, 20)
	retValue = obj.GetValue(2, 2)
	if retValue != 20 {
		t.Fatalf("Error!")
	}

	// 测试3
	input = [][]int{{2, 8}, {8, 8}, {7, 4}}
	obj = Constructor(input)

	retValue = obj.GetValue(1, 0)
	if retValue != 8 {
		t.Fatalf("Error!")
	}
	obj.UpdateSubrectangle(1, 1, 1, 1, 4)
	retValue = obj.GetValue(1, 0)
	if retValue != 8 {
		t.Fatalf("Error!")
	}
	retValue = obj.GetValue(0, 0)
	if retValue != 2 {
		t.Fatalf("Error!")
	}
	obj.UpdateSubrectangle(2, 1, 2, 1, 9)
	retValue = obj.GetValue(1, 1)
	if retValue != 4 {
		t.Fatalf("Error!")
	}
	retValue = obj.GetValue(1, 0)
	if retValue != 8 {
		t.Fatalf("Error!")
	}
}
