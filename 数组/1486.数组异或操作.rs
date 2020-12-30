/// 解法一
/// 直接挨个亦或下去就可以了
struct Solution {
}
impl Solution {
    pub fn xor_operation(n: i32, start: i32) -> i32 {
        let mut result = start;
        for i in 1 .. n {
            result = result ^ (i * 2 + start);
        }
        result
    }
}

/// main方法测试
fn main() {
    let n: i32 = 5;
    let start: i32 = 0;
    let result = Solution::xor_operation(n, start);
    println!("result{}",result);
}