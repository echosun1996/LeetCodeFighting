/// 解法一
/// 将数字转为字符串，计算字符串的长度是否是偶数即可
struct Solution {
}
impl Solution {
    pub fn find_numbers(nums: Vec<i32>) -> i32 {
        let mut result = 0;
        for i in 0 .. nums.len() {
            if nums[i].to_string().len()%2 == 0 {
                result = result + 1;
            }
        }
        result
    }
}

/// main方法测试
fn main() {
    let nums = vec![12,5,3,4,11];
    let result = Solution::find_numbers(nums);
    println!("result{}",result);
}