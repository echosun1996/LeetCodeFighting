/// 解法一
/// 从前往后找到第一个等于val的值，从后往前找到第一个不等于val的值，交换它们即可
struct Solution {
}
impl Solution {
    pub fn remove_element(nums: &mut Vec<i32>, val: i32) -> i32 {
        if(nums.len()<=0){
            return 0;
        }
        let mut i:usize = 0;
        let mut j:usize = nums.len()-1;
        while i <= j {
            if(nums[i.clone()] != val){
                i = i.clone()+1;
                continue;
            }
            if(nums[j.clone()] == val){
                if(j<=0){
                    break;
                }
                j = j.clone()-1;
                continue;
            }
            nums[i.clone()] = nums[j.clone()].clone();
            nums[j.clone()] = val.clone();
            i = i.clone()+1;
            j = j.clone()-1;
        }
        return j as i32;
    }
}

/// main方法测试
fn main() {
    let mut nums = vec![3];
    let val:i32 = 2;
    let sum = Solution::remove_element(&mut nums,val);
    println!("sum{},nums{:?}",sum,nums);
}