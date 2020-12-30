/// 解法一
/// 暴力破解，向后累加即可
struct Solution {
}
impl Solution {
    pub fn running_sum(nums: Vec<i32>) -> Vec<i32> {
        let mut result = nums.clone();
        for i in 1 .. nums.len() {
            result[i] = result[i] + result[i - 1];
        }
        result
    }
}

/// main方法测试
fn main() {
    let nums = vec![1,2,3,4,4,3,2,1];
    let result = Solution::running_sum(nums);
    println!("result{:?}",result);
}