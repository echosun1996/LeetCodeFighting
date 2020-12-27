/// 解法一(非排序数组用此法也正确)
/// 暴力破解，将每个数存进map，存为1(map中不存在这个数字)
/// 如果map中已存在这个数字
///  - 这个数字存值>1，则证明已有2个重复，移除这个数即可
///  - 这个数字存值<=1，将这个数存为2
//use std::collections::HashMap;
//struct Solution {
//}
//impl Solution {
//    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
//        let mut map: HashMap<i32, i32> = HashMap::new();
//        let mut result:usize = 0;
//        for i in 0 .. nums.len() {
//            if map.contains_key(&nums[i - result]) {
//                if map.get(&nums[i - result]).unwrap() > &1 {
//                    nums.remove(i - result);
//                    result = result + 1;
//                } else {
//                    map.insert(nums[i - result], 2);
//                }
//                continue;
//            }
//            map.insert(nums[i - result], 1);
//        }
//        nums.len() as i32
//    }
//}

/// 解法二
/// 将不重复的数往前移即可，因为是两个两个的重复，
/// 往前移后会导致下一组数字比较出错，故将往前移数字延迟一步
struct Solution {
}
impl Solution {
    pub fn remove_duplicates(nums: &mut Vec<i32>) -> i32 {
        if nums.len() <=1{
            return nums.len() as i32;
        }
        let mut result:usize = 2;
        for i in 2 .. nums.len(){
            if nums[i -2] == nums[i] {
                continue;
            }
            if result > 2 {
                nums[result - 1] = nums[i - 1];
            }
            result = result + 1;
        }
        nums[result - 1] = nums[nums.len() - 1];
        result.clone() as i32
    }
}


/// main方法测试
fn main() {
    let mut nums = vec![1,1,2,2,2,2,2,1];
    let result = Solution::remove_duplicates(&mut nums);
    println!("result{}",result);
    println!("nums{:?}",nums);
}