/// 解法一
/// 将数组排序，循环一遍确定第一个数，第二和第三个数采用夹逼指针方法确定
///
/// 夹逼指针方法
/// 从第一个数开始向后循环至倒数第三个数，这是确定的第一位数
/// 第二个数从第一位数坐标后面那个数向后找，第三个数从最后一位坐标向前找
/// 所加结果大于目标值则第三位数向前移，所加结果小于目标值则第二位数向后移
/// 当第二位数到达第三位数所在坐标，再继续查找会找到重复的，故直接跳出循环
/// 因不能查找重复的数组，但输入数组中有重复的值，故第一二三位数的下一个数和前一个一样时都直接略过，不进行本轮查找
/// 注：因为排序后左小右大，向前移即减小结果值，向后移即增大结果值
/// 注：下一个数和前一个要么已经找到了(再找就重复)，要么就是本来就达不到条件，故不用担心因此遗漏
struct Solution {
}
impl Solution {
    pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut result = Vec::new();
        if nums.len() < 3 {
            return result;
        }

        let mut nums_copy = nums.clone();
        nums_copy.sort();
        for i in 0 .. nums_copy.len() - 2 {
            if i>0 && nums_copy[i] == nums_copy[i-1]{
                continue;
            }
            let mut j = i+1;
            let mut k = nums_copy.len() - 1;
            while j < k {
                if j > (i + 1) && nums_copy[j] == nums_copy[j-1]{
                    j = j + 1;
                    continue;
                }
                if k < (nums_copy.len() - 1) && nums_copy[k] == nums_copy[k+1]{
                    k = k - 1;
                    continue;
                }
                if (nums_copy[i] + nums_copy[j] + nums_copy[k]) == 0{
                    result.push(vec![nums_copy[i],nums_copy[j],nums_copy[k]]);
                    j = j + 1;
                    k = k - 1;
                } else if (nums_copy[i] + nums_copy[j] + nums_copy[k]) > 0 {
                    k = k - 1;
                } else {
                    j = j + 1;
                }
            }
        }
        result
    }
}

/// main方法测试
fn main() {
    let nums = vec![-2,0,0,2,2];
    let result = Solution::three_sum(nums);
    println!("result = {:?}",result);
}