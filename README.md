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

| Problem Number & Title | Description | Difficulty | Go State | Rust State |
| -------------- | ----- | ----- | -------- | ------------ |
| 1.两数之和 | 给总和求两数下标 | Easy | √ | √ |
| 4.寻找两个正序数组的中位数 | | Easy | | √ |
| 15.三数之和 | 夹逼指针 | Easy | | √ |
| 41. 缺失的第一个正数 | | Easy | | √ |
| 1476.子矩形查询 | 查询块更新后的矩阵元素值 | Easy | √ |  |
| 1480.一维数组的动态和 | 累加求和 | Easy | √ | √ |
| 1672.最富有客户的资产总量 |  | Easy | √ | √ |
| 1470.重新排列数组 |  | Easy  | √ | √ |
| 1431.拥有最多糖果的孩子 |  | Easy | √ | √ |
| 1512.好数对的数目 |  | Easy | √ | √ |
| 1486.数组异或操作 |  |Easy  |√ | √ |
| 1395.统计作战单位数 |  |Easy  | √ | √ |
| 1313.解压缩编码列表 |  | Easy | √ | √ |
| 1389.按既定顺序创建目标数组 |  | Easy | √ | √ |
| 1365.有多少小于当前数字的数字 |  | Easy | √ | √ |
| 283.移动零 |  | Easy | √ | √ |
| 27.移除元素 | 从数组中移除给定值的元素 | Easy | √ | √ |
| 26.删除排序数组中的重复项 | 「如题」 | Easy | √ | √ |
| 80.删除排序数组中的重复项 II | 限制重复元素个数<=2 | Easy | √ | √ |
| 1266.访问所有点的最小时间 |  | Easy | √ | |
| 1295.统计位数为偶数的数字 |  | Easy | √ | √ |
| 1366.通过投票对团队排名 |  | Medium | √ | |
| 面试题0804.幂集 | 数组子集的所有组合 | Easy | √ | |
| 1572.矩阵对角线元素的和 | 「如题」 | Easy | √ | |
| 1588.所有奇数长度子数组的和 | 「如题」 | Easy | √ | |
| 1409. 查询带键的排列 |  | | | |
|  |  | | | |
|  |  | | | |
|  |  | | | |
###字符串
| Problem Number & Title | Description | Difficulty | Go State | Rust State |
| -------------- | ----- | ----- | -------- | ------------ |
| 32.最长有效括号| | Medium | | √ |
### 数学
| Problem Number & Title | Description| Difficulty | Go State | Rust State |
| -------------- | ----- | ----- | -------- | ------------ |
| 258.各位相加 | 求数根公式 | Medium | √ | √ |
| 42.接雨水| 韦恩图 | Medium |  | √ |


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