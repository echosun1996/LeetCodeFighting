package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// getFileNameList
// 传入文件夹名称，返回文件夹内所有文件名的列表
func getFileNameList(dirPath string) []string {
	var fileNameList []string
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}
	// 获取文件，并输出它们的名字
	for _, file := range files {
		fileNameList = append(fileNameList, strings.ReplaceAll(file.Name(), " ", "")) // 去掉空格
	}
	return fileNameList
}

// READMEReader
// 传入readme路径和题目类型，按照题型读取readme
func READMEReader(readmePath string, problemType string) [][]string {
	problemType = "### " + problemType // 在readme中，题型是三级标题
	var readmeDetails [][]string

	fin, err := os.OpenFile(readmePath, os.O_RDONLY, 0)
	if err != nil {
		panic(err)
	}
	defer fin.Close()

	reader := bufio.NewReader(fin)
	startRecord := false
	for {
		line, err := reader.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		line = strings.Replace(line, "\f", "", -1)

		if find := strings.Contains(line, problemType); find {
			startRecord = true
			continue
		}
		if find := strings.Contains(line, "#"); startRecord && find {
			startRecord = false
			continue
		}
		if startRecord {

			if strings.Contains(line, "Problem Number") || strings.Contains(line, ".") {
				// 处理标题行和带有题目的行
				line = strings.ReplaceAll(line, " ", "") // 去掉空格
				splitResult := strings.Split(line, "|")

				readmeDetails = append(readmeDetails, splitResult[1:len(splitResult)-1])
				continue

			} else if strings.Contains(line, "---") {
				// 处理表格标记
				continue
			}
		}
	}
	return readmeDetails
}

// getColumnNumber
// 获取对应语言在readme文件中的列号
func getColumnNumber(readmeDetails [][]string, fileSuffix []string) int {
	for ix, value := range readmeDetails[0] {
		if strings.Contains(value, fileSuffix[0]) {
			return ix
		}
	}
	return -1
}

// getProblemNameFromReadme
// 从readme中获取对应语言所完成的题目列表
func getProblemNameFromReadme(readmeDetails [][]string, fileSuffix []string) []string {
	var problemNames []string

	// 获取对应语言在readme中的列号
	columnNumber := getColumnNumber(readmeDetails, fileSuffix)

	for _, value := range readmeDetails {
		if strings.Contains(value[columnNumber], "√") {
			problemNames = append(problemNames, value[0])
		}
	}
	return problemNames
}

// stateCheck
// 返回匹配结果。
// 其中 fileSuffix 中第一个字符串时README中的标题，第二个字符串是文件的后缀名
func stateCheck(fileNameList []string, readmeDetails [][]string, fileSuffix []string) ([]string, []string, int) {
	finishSum := 0
	readmeProblemNames := getProblemNameFromReadme(readmeDetails, fileSuffix)

	// 在fileNameList中获取每一个对应文件后缀名的字符串
	readmeProblemNamesSize := len(readmeProblemNames)
	fileNameListSize := len(fileNameList)

	// 遍历所有的文件，找到readme中同名项，彼此删除
	for i := 0; i < fileNameListSize; i++ {
		// 文件后缀名符合所选的语言
		if !strings.Contains(fileNameList[i], fileSuffix[1]) {
			fileNameList = append(fileNameList[:i], fileNameList[i+1:]...)
			fileNameListSize--
			i = -1 // 从头开始寻找
			continue
		}
		fileProblemNum := strings.Split(fileNameList[i], ".")[0]
		for j := 0; j < readmeProblemNamesSize; j++ {
			readmeProblemNum := strings.Split(readmeProblemNames[j], ".")[0]
			if strings.EqualFold(fileProblemNum, readmeProblemNum) {
				fileNameList = append(fileNameList[:i], fileNameList[i+1:]...)
				readmeProblemNames = append(readmeProblemNames[:j], readmeProblemNames[j+1:]...)
				i = -1 // 从头开始寻找
				fileNameListSize--
				readmeProblemNamesSize--
				finishSum++
				break
			}
		}
	}
	return fileNameList, readmeProblemNames, finishSum
}

// main 本程序用于在做题列表和文件列表中相互筛选，以选出没有被标记或仅被标记的题目。
// 通过本程序，也可以统计完成题目数量。
func main() {
	dirPath := []string{"数组", "数学"}                          // dirPath既是文件夹的名字，也是题目所属分类的名字
	mdPath := "README.md"                                    // README的路径
	fileSuffix := [][]string{{"Go", ".go"}, {"Rust", ".rs"}} // 待检测的列表题目以及对应的后缀名
	SCKEY := ""                                              // 从 github action 中获取的KEY，用于向WeChat发送通知
	var messageTitle strings.Builder
	var messageDetail strings.Builder

	for idx, args := range os.Args {
		if idx == 1 {
			SCKEY = args
			println("[发现SCKEY!]")
		}
	}

	// 添加提交日志
	cmd := exec.Command("git", "config", "--global", "core.quotepath", "false") // 设置git显示中文
	err := cmd.Run()
	if err != nil {
		log.Fatalf("设置git显示中文失败: %s\n", err)
	}
	cmd = exec.Command("git", "pull", "--unshallow") //  转换存储库为完整存储库
	_ = cmd.Run()
	cmd = exec.Command("git", "log", "--stat", "--oneline", "-1") // 精简显示最后一次commit日志
	buf, _ := cmd.CombinedOutput()
	messageDetail.WriteString("# 提交日志\n\n```\n")
	messageDetail.Write(buf)
	messageDetail.WriteString("\n```\n")
	println(messageDetail.String())

	for _, dirPathValue := range dirPath {
		println("#", dirPathValue)
		messageDetail.WriteString("# " + dirPathValue + "\n\n")
		for fileSuffixIndex, fileSuffixValue := range fileSuffix {
			println("##", fileSuffixValue[0], "题型检查结果")
			messageDetail.WriteString("## " + fileSuffixValue[0] + "类型记录检查结果" + "\n\n")
			fileNameList := getFileNameList(dirPathValue)
			readmeDetails := READMEReader(mdPath, dirPathValue)
			fileNameList, readmeProblemNames, finishSum := stateCheck(fileNameList, readmeDetails, fileSuffixValue)
			if len(readmeProblemNames) == 0 {
				println("1. [OK] README中标记为完成的文件已全部添加。")
				messageDetail.WriteString("1. <font color=#008000>[OK]</font> README中标记为完成的文件已全部添加。" + "\n\n")
			} else {
				println("1. [ERROR] README中如下文件标记为完成，但未找到程序文件:")
				messageDetail.WriteString("1. <font color=red>[ERROR]</font> README中如下文件标记为完成，但未找到程序文件:" + "\n\n")
				fmt.Printf("%v\n\n", readmeProblemNames)
				for _, value := range readmeProblemNames {
					messageDetail.WriteString("* " + value + "\n")
				}
			}
			if len(fileNameList) == 0 {
				println("2. [OK] 文件夹中所有文件都被标记为完成。")
				messageDetail.WriteString("2. <font color=#008000>[OK]</font> 文件夹中所有文件都被标记为完成。" + "\n\n")
			} else {
				println("2. [ERROR] 文件夹[", dirPathValue, "]中存在如下文件尚未被标记为完成:")
				messageDetail.WriteString("2. <font color=red>[ERROR]</font> 文件夹[" + dirPathValue + "]中存在如下文件尚未被标记为完成:" + "\n\n")
				fmt.Printf("%v\n\n", fileNameList)
				for _, value := range fileNameList {
					messageDetail.WriteString("* " + value + "\n")
				}
			}
			messageDetail.WriteString("\n>[**" + dirPathValue + "**]类型的题目使用[**" + fileSuffixValue[0] + "**]语言已完成[**" + strconv.Itoa(finishSum) + "**]道。\n\n")

			// 拼接发往WeChat的字符串
			if fileSuffixIndex == 0 {
				messageTitle.WriteString("[" + dirPathValue + "]")
			}
			messageTitle.WriteString(fileSuffixValue[0] + ":")
			messageTitle.WriteString(strconv.Itoa(finishSum) + ";")
		}
	}
	if SCKEY != "" {
		sendToWechat(SCKEY, messageTitle.String(), messageDetail.String())
	}
	println(messageTitle.String())
}

// 以post请求形式发送消息到WeChat
func sendToWechat(SCKEY string, messageTitle string, messageDetail string) {
	//post请求提交application/x-www-form-urlencoded数据
	form := make(url.Values)
	form.Set("text", messageTitle)
	form.Add("desp", messageDetail)
	resp, _ := http.PostForm("https://sc.ftqq.com/"+SCKEY+".send", form)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Post request result: %s\n", string(body))
}

// 以get请求形式发送消息到WeChat
//func sendToWechat(SCKEY string, message string) {
//	resp, err := http.Get("https://sc.ftqq.com/" + SCKEY + ".send?text=" + message)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer resp.Body.Close()
//	body, _ := ioutil.ReadAll(resp.Body)
//	fmt.Println("Received from server:" + string(body))
//}
