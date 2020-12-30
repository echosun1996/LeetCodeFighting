/// 解法一
/// 求出最大值，然后挨个比较相加后能不能大于等于最大值
struct Solution {
}
impl Solution {
    pub fn kids_with_candies(candies: Vec<i32>, extra_candies: i32) -> Vec<bool> {
        let mut max:i32 = 0;
        let mut result = Vec::new();
        for i in 0 .. candies.len() {
            if candies[i] > max {
                max = candies[i].clone();
            }
        }
        for i in 0 .. candies.len() {
            if candies[i] + extra_candies >= max {
                result.push(true);
            } else {
                result.push(false);
            }
        }
        result
    }
}

/// main方法测试
fn main() {
    let candies = vec![12,5,3,4,11];
    let extra_candies:i32 = 8;
    let result = Solution::kids_with_candies(candies, extra_candies);
    println!("result{:?}",result);
}