/// 解法一
/// 暴力破解直接算
struct Solution {
}
impl Solution {
    pub fn maximum_wealth(accounts: Vec<Vec<i32>>) -> i32 {
        let mut result = 0;
        for i in 0 .. accounts.len() {
            let mut sum = 0;
            for j in 0 .. accounts[i].len() {
                sum = sum + accounts[i][j];
            }
            if sum > result {
                result = sum;
            }
        }
        result
    }
}

/// main方法测试
fn main() {
    let accounts = vec![ vec![1,2,3], vec![3,2,1]];
    let result = Solution::maximum_wealth(accounts);
    println!("result{}",result);
}