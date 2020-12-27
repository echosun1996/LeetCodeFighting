/// 解法一(看错题了，非排序数组用此法也正确)
/// 暴力破解，将每个数存进map，如果map中已存在，则证明重复，移除这个数即可
//use std::collections::HashMap;
//struct Solution {
//}
//impl Solution {
//    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
//        let mut map: HashMap<i32, i32> = HashMap::new();
//        let mut result:usize = 0;
//        for i in 0 .. nums.len() {
//            if map.contains_key(&nums[i.clone() as usize - result]) {
//                result = result + 1;
//                nums.remove(i - result);
//                continue;
//            }
//            map.insert(nums[i.clone() as usize - result], 1);
//        }
//        nums.len() as i32
//    }
//}

/// 解法二
/// 不考虑非顺序数组，只判断当前数和上一个数是否重复
//struct Solution {
//}
//impl Solution {
//    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
//        let mut result:usize = 0;
//        for i in 1 .. nums.len() {
//            if nums[i - result -1] == nums[i - result] {
//                result = result + 1;
//                nums.remove(i - result);
//                continue;
//            }
//        }
//        nums.len() as i32
//    }
//}

/// 解法三
/// 将不重复的数往前移即可
//struct Solution {
//}
//impl Solution {
//    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
//        let mut result:usize = 0;
//        for i in 1 .. nums.len() {
//            if nums[i -1] == nums[i] {
//                result = result + 1;
//                continue;
//            }
//            nums[i-result] = nums[i];
//
//        }
//        (nums.len() - result) as i32
//    }
//}

/// 解法四
/// 将不重复的数往前移即可
struct Solution {
}
impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.len() <=0{
            return 0;
        }
        let mut result:usize = 1;
        for i in 1 .. nums.len() {
            if nums[i -1] == nums[i] {
                continue;
            }
            nums[result] = nums[i];
            result = result + 1;
        }
        result as i32
    }
}

/// main方法测试
fn main() {
    let mut nums = vec![1,1,2,2,1,1,2];
    let result = Solution::remove_duplicates(&mut nums);
    println!("result{}",result);
    println!("nums{:?}",nums);
}