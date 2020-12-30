// 1366.通过投票对团队排名 - https://leetcode-cn.com/problems/rank-teams-by-votes
// By EchoSun
// 2020-12-26 09:57:12
package main

import (
	"reflect"
	"testing"
)

// 解法1 先使用map存储投票情况，再将结果转存到结构体中，再使用冒泡排序对结果按照位次排序。
// 需要注意如果结果相同，需要按字符串顺序再排序。
// 本题难点在于map与结构体的使用上。
func rankTeams(votes []string) string {
	// 使用map保存投票信息
	voteMap := make(map[string][]int)
	voteLength := len(votes[0])
	for _, vote := range votes {
		for cI, char := range vote {
			if _, ok := voteMap[string(char)]; !ok {
				voteMap[string(char)] = make([]int, voteLength)
			}
			voteMap[string(char)][cI]++
		}
	}

	// 将map转为结构体数组
	type voter struct {
		name string
		vote []int
	}
	var voterList []voter
	for value := range voteMap {
		voterList = append(voterList, voter{value, voteMap[value]})
	}
	//fmt.Printf("%v", voterList)

	// 对结构体数组进行冒泡排序
	for voteI := voteLength - 1; voteI >= 0; voteI-- {
		for i := 0; i < len(voterList)-1; i++ {
			for j := 0; j < len(voterList)-1-i; j++ {
				// 从后往前，按照每个位置出现的次数从高到低排序;若结果完全并列，按照字母升序排名
				if voterList[j].vote[voteI] < voterList[j+1].vote[voteI] ||
					reflect.DeepEqual(voterList[j].vote, voterList[j+1].vote) && voterList[j].name > voterList[j+1].name { // numSlice[voteI] < numSlice[j+1] (从小到大)
					temp := voterList[j]
					voterList[j] = voterList[j+1]
					voterList[j+1] = temp
				}
			}
		}
	}

	// 从排序好的结构体中把name单独拼接到一起，组成答案
	ret := ""
	for _, value := range voterList {
		ret += value.name
	}

	return ret
}

/////////////

// 测试函数

func TestProblem1366(t *testing.T) {
	type args struct { // 测试样例的输入
		input []string
	}
	type test struct {
		name string // 测试样例名称
		args args   // 测试样例的参数
		want string // 测试样例的返回值
	}
	tests := []test{
		{"test1", args{[]string{"ABC", "ACB", "ABC", "ACB", "ACB"}}, "ACB"},
		{"test2", args{[]string{"WXYZ", "XYZW"}}, "XWYZ"},
		{"test3", args{[]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}}, "ZMNAGUEDSJYLBOPHRQICWFXTVK"},
		{"test4", args{[]string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}}, "ABC"},
		{"test5", args{[]string{"M", "M", "M", "M"}}, "M"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rankTeams(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[Wrong Answer]运行结果:%v, 正确答案: %v", got, tt.want)
			}
		})
	}
}
