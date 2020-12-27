/// 解法一
/// 暴力破解，两层循环，将index[i]位置后面的数通通后移一位，在index[i]位置插入num[i]
struct Solution {
}
impl Solution {
    pub fn create_target_array(nums: Vec<i32>, index: Vec<i32>) -> Vec<i32> {
        let mut result = vec![0;nums.len()];
        for i in 0 .. nums.len() {
            for j in index[i] as usize .. i{
                result[i - j + index[i] as usize] = result[i - j + index[i]  as usize - 1];
            }
            result[index[i] as usize] = nums[i];
        }
        result
    }
}

/// main方法测试
fn main() {
    let nums = vec![0,1,2,3,4];
    let index = vec![0,1,2,2,1];
    let result = Solution::create_target_array(nums, index);
    println!("result{:?}",result);
}