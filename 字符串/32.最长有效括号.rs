/// 解法一
/// 左括号left加一，右括号right+1，当left等于right时，记录一次最大值，此操作正反方向各循环一次
struct Solution {
}
impl Solution {
    pub fn longest_valid_parentheses(s: String) -> i32 {
        let mut max = 0;
        let mut right:i32 = 0;
        let mut left:i32 = 0;
        for i in 0 .. s.len() {
            if &s[i .. i+1] == "(" {
                left = left + 1;
            } else {
                right = right + 1;
            }

            if left == right{
                if 2*left > max{
                    max = 2*left;
                }
            } else if left < right {
                left = 0;
                right = 0;
            }
        }
        left = 0;
        right = 0;
        for i in 0 .. s.len() {
            if &s[s.len()-i-1 .. s.len()-i] == ")" {
                right = right + 1;
            } else {
                left = left + 1;
            }

            if left == right{
                if 2*left > max{
                    max = 2*left;
                }
            } else if left > right {
                left = 0;
                right = 0;
            }
        }
        max
    }
}

/// main方法测试
fn main() {
    let s = String::from("(()");
    let result = Solution::longest_valid_parentheses(s);
    println!("result = {}",result);
}