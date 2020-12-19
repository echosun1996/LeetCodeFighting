package main

import "fmt"

// bracketsMatch 括号匹配
// 职匹配中括号
func bracketsMatch(input string) bool {
	temp := 0
	for _, value := range input {
		if value == '[' {
			temp += 1
		} else if value == ']' {
			temp -= 1
		}
	}
	if temp != 0 {
		//println("@ Brackets Temp Value: [",temp,"]")
		fmt.Printf("@ Brackets Temp Value: [%v]\n", temp)
		println("@ ERROR: brackets NOT match!")
		return false
	}
	return true
}

// 中括号格式样例转
// Go 语言的字符串是不可变的。
// 修改字符串时，可以将字符串转换为 []byte 进行修改。
// []byte 和 string 可以通过强制类型转换互转。
func inputFormatChanger(input string) string {
	output := []byte(input)
	if !bracketsMatch(input) {
		return ""

	}
	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			output[i] = '{'
		} else if input[i] == ']' {
			output[i] = '}'
		}

	}
	return string(output)
}
func main() {

	input := "[[[[1,1,1],[2,2,2],[3,3,3]]],[0,0],[0,0,2,2,100],[0,0],[2,2],[1,1,2,2,20],[2,2]]"
	output := inputFormatChanger(input)
	println(output)
}
