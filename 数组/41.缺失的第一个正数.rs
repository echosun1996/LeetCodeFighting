/// 解法一
/// 第一次循环将正整数存进map，第二次循环从最小的正整数开始遍历，map没有该数即返回
//use std::collections::HashMap;
//struct Solution {
//}
//impl Solution {
//    pub fn first_missing_positive(nums: Vec<i32>) -> i32 {
//        let mut map: HashMap<i32, i32> = HashMap::new();
//        for i in 0 .. nums.len() {
//            if nums[i] > 0{
//                map.insert(nums[i],1);
//            }
//        }
//
//        for i in 1 .. nums.len()+2 {
//            if !map.contains_key(&(i.clone() as i32)) {
//                return i as i32;
//            }
//        }
//        return 1;
//    }
//}

/// 解法二
/// 第一次循环将正整数存进map，第二次循环从最小的正整数开始遍历，map没有该数即返回(用loop进行循环操作)
use std::collections::HashMap;
struct Solution {
}
impl Solution {
    pub fn first_missing_positive(nums: Vec<i32>) -> i32 {
        let mut map: HashMap<i32, i32> = HashMap::new();
        for i in 0 .. nums.len() {
            if nums[i] > 0{
                map.insert(nums[i],1);
            }
        }

        let mut i = 1;
        loop {
            if !map.contains_key(&(i.clone() as i32)) {
                return i;
            }
            i = i + 1;
        }
    }
}

/// main方法测试
fn main() {
    let nums=vec![];
    let result = Solution::first_missing_positive(nums);
    println!("result = {}",result);
}