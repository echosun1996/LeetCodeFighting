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
///
/// 证明
/// 设n为输入数，m为n第一次各位相加后的数
/// n=10^n*an+...10*a1+a0
/// m=an+...a1+a0
/// n-m=(10^n-1)*an+...9*a1
/// 10^n-1是9的倍数，故每次相加后都减少9的倍数
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