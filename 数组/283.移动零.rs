/// 解法一
/// 找到一个0则count加1，找到第一位不是0的数往前移count位，从后往前数count位全部赋0
struct Solution {
}
impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let mut count:usize = 0;
        let mut i:usize = 0;
        let mut len:usize = nums.len();
        while i < (len+count) {
            if i>=len {
                nums[i.clone()-count.clone()] = 0;
            }else if nums[i.clone()] == 0{
                count = count.clone()+1;
            } else {
                nums[i.clone() - count.clone()] = nums[i.clone()];
            }
            i = i.clone() + 1;
        }
    }
}

/// main方法测试
fn main() {
    let mut nums = vec![0,1,0,3,12];
    Solution::move_zeroes(&mut nums);
    println!("nums{:?}",nums);
}