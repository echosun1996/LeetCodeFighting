/// 解法一
/// 从个位数开始往前加，如果大于10了就再加一次(两数相加肯定不会大于二十)，一直加到最大那个数
//struct Solution {
//}
//impl Solution {
//    pub fn add_digits(num: i32) -> i32 {
//        let mut num_copy = num.clone();
//        let mut result:i32 = num_copy.clone() % 10;
//        while num_copy >= 10 {
//            result = result + num_copy.clone()/10%10;
//            if(result>=10){
//                result = result.clone() % 10 + result.clone()/10%10;
//            }
//            num_copy = num_copy/10;
//        }
//        result
//    }
//}

/// 解法二
/// 每次相加后都减少9的倍数
struct Solution {
}
impl Solution {
    pub fn add_digits(num: i32) -> i32 {
        (num-1)%9 + 1
    }
}

/// main方法测试
fn main() {
    let num:i32 = 19;
    let result = Solution::add_digits(num);
    println!("result{}",result);
}