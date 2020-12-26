/// 解法一
/// 暴力破解，直接两层for循环数组，得到相加数等于目标值的加数的下标
//struct Solution {
//}
//impl Solution {
//    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
//        let mut i: i32 = 0;
//        let mut j: i32 = 0;
//        let len: i32 = nums.len() as i32;
//        while i < (len - 1) {
//            j = i + 1;
//            while j < len {
//                if nums[i.clone() as usize] + nums[j as usize] == target {
//                    return vec![i, j];
//                }
//                j = j.clone() + 1;
//            }
//            i = i.clone() + 1;
//        }
//        vec![i, j]
//    }
//}

/// 解法2
/// 一层循环，用map存储其中一个加数和它的下标(循环数组时，map中没有找到与当前加数相加符合结果值的另一个加数，则将当前加数存进map)，
/// 循环一次数组得到相加数为目标值的map中加数与数组中加数的下标
use std::collections::HashMap;
struct Solution {
}
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut i: i32 = 0;
        let len: i32 = nums.len() as i32;
        let mut map: HashMap<i32, i32> = HashMap::new();
        while i < len {
            let temp = target.clone() - nums[i as usize];
            if map.contains_key(&temp) {
                return vec![i, *map.get(&temp).unwrap()]
            }
            map.insert(nums[i.clone() as usize], i.clone());
            i = i.clone() + 1;
        }
        vec![i, i]
    }
}


/// main方法测试
fn main() {
    let nums = vec![3,2,4];
    let target:i32 = 6;
    let result = Solution::two_sum(nums,target);
    println!("result{},{}",result[0],result[1]);
}
