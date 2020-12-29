# Leetcode刷题计划
## 练习流程
### Golang
1. 在题型列表中，按照难度递增顺序刷题
1. 复制题目到笔记本中
1. 再将题目保存到README的题目列表中
1. 创建测试程序，以题目名+`_test.go`命名
1. 完成题目后fmt程序后提交
1. 结果保存到笔记本中(以leetcode页面的格式)
1. 更新README题目列表的“难度”和“状态”信息

## 完成进度数组

### 数组

| Problem Number & Title | Difficulty | Go State | Rust State |
| -------------- | ----- | -------- | ------------ |
| 1.两数之和 | Easy | √ | √ |
| 1476.子矩形查询 | Easy | √ |  |
| 1480.一维数组的动态和 | Easy | √ | |
| 1672.最富有客户的资产总量 | Easy | √ | |
| 1470.重新排列数组 | Easy  | √ | |
| 1431.拥有最多糖果的孩子 | Easy | √ | √ |
| 1512.好数对的数目 | Easy | √ | |
| 1486.数组异或操作 |Easy  |√ | |
| 1395.统计作战单位数 |Easy  | √ | √ |
| 1313.解压缩编码列表 | Easy | √ | √ |
| 1389.按既定顺序创建目标数组 | Easy | √ | √ |
| 1365.有多少小于当前数字的数字 | Easy | √ | √ |
| 283.移动零 | Easy | √ | √ |
| 27.移除元素 | Easy | √ | √ |
| 26.删除排序数组中的重复项 | Easy | √ | √ |
| 80.删除排序数组中的重复项 II | Easy | √ | √ |
| 1266.访问所有点的最小时间 |  | | |
| 1295.统计位数为偶数的数字 | Easy | √ | √ |
|  |  | | |
|  |  | | |
|  |  | | |
|  |  | | |
|  |  | | |
|  |  | | |
|  |  | | |

### 数学
| Problem Number & Title | Difficulty | Go State | Rust State |
| -------------- | ----- | -------- | ------------ |
| 258.各位相加 | Medium | √ | √ |


## 疑难问题
### 1. git拉取失败

运行`git pull`时报如下错误

```shell
git pull
# remote: Enumerating objects: 21, done.
# remote: Counting objects: 100% (21/21), done.
# remote: Compressing objects: 100% (9/9), done.
# remote: Total 15 (delta 8), reused 13 (delta 6), pack-reused 0
# error: insufficient permission for adding an object to repository database .git/objects
# fatal: failed to write object
# fatal: unpack-objects failed
```

这是出现了权限问题，需要执行：

```shell
cd .git/objects

ll | grep root  # 可以看到许多权限错误的文件

sudo chown -Rv <用户名> *

ll | grep root  # 权限错误的文件消失
```