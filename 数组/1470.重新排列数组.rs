/// 解法一
/// 循环数组的一半，取第i和i+n位元素
struct Solution {
}
impl Solution {
    pub fn shuffle(nums: Vec<i32>, n: i32) -> Vec<i32> {
        let mut result = vec![0;nums.len()];
        for i in 0 .. nums.len()/2 {
            result[i*2] = nums[i];
            result[i*2+1] = nums[i + n as usize];
        }
        result
    }
}

/// main方法测试
fn main() {
    let nums = vec![1,2,3,4,4,3,2,1];
    let n:i32 = 4;
    let result = Solution::shuffle(nums, n);
    println!("result{:?}",result);
}