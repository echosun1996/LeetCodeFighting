/// 解法一
/// 遍历一遍，将num[i]个num[i+1]串起来即可
struct Solution {
}
impl Solution {
    pub fn decompress_rl_elist(nums: Vec<i32>) -> Vec<i32> {
        let mut result =  vec![];
        let mut i:usize = 0;
        while i < nums.len() - 1 {
            let mut j:i32 = 0;
            while j < nums[i.clone()] {
                result.push(nums[i.clone() + 1]);
                j = j.clone() + 1;
            }
            i = i.clone()+2;
        }
        return result;
    }
}

/// main方法测试
fn main() {
    let nums = vec![1,2,3,4];
    let result = Solution::decompress_rl_elist(nums);
    println!("result{:?}",result);
}