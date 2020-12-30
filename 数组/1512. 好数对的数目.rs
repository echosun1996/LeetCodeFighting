/// 解法一
/// 暴力破解，双重循环符合条件就加1
struct Solution {
}
impl Solution {
    pub fn num_identical_pairs(nums: Vec<i32>) -> i32 {
        let mut result = 0;
        for i in 0 .. nums.len()-1 {
            for j in i+1 .. nums.len() {
                if nums[i] == nums[j] {
                    result = result + 1
                }
            }
        }
        result
    }
}

/// main方法测试
fn main() {
    let nums = vec![0,1,0,3,12];
    let result = Solution::num_identical_pairs(nums);
    println!("result{}",result);
}