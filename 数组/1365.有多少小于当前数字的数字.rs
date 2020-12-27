/// 解法一
/// 暴力破解，两层循环，里层循环数小于外层循环数，则里层加一
struct Solution {
}
impl Solution {
    pub fn smaller_numbers_than_current(nums: Vec<i32>) -> Vec<i32> {
        let len:usize = nums.len();
        let mut result = vec![0;len];
        let mut min = nums[0];
        for i in 0 .. len{
            for j in 0 .. len {
                if nums[i.clone()] > nums[j.clone()]{
                    result[i.clone()] = result[i.clone()] + 1;
                }
            }
        }
        return result;
    }
}

/// main方法测试
fn main() {
    let nums = vec![6,5,4,8];
    let result = Solution::smaller_numbers_than_current(nums);
    println!("result{:?}",result);
}