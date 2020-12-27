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
/// 证明每次相加后都减少9的倍数
/// 解:
/// 设n为输入数，m为n第一次各位相加后的数
/// n=10^n*an+...10*a1+a0
/// m=an+...a1+a0
/// n-m=(10^n-1)*an+...9*a1
/// 因为10^n-1是9的倍数，故每次相加后都减少9的倍数，则n-m=9的倍数
///
/// 证明10^n-1是9的倍数
/// 解:
/// 由10^n-1=9*10^(n-1)+10^(n-1)-1，得出
/// 证明10^(n-1)-1是9的倍数即可证明10^n-1是9的倍数，
/// 由此递归直至n=2，10^(n-1)-1=10^(2-1)-1=9是9的倍数
/// 故10^n-1是9的倍数
///
/// 证明各位相加结果等于(n-1)%9+1
/// 解:
/// 设最终结果为m，且因为n-m=9的倍数
/// 则推导出m!=9时，m=n%9
/// 但m=9时，n%9=0，此时可破坏n%9=0的情况，
/// 可先用n-1并%9之后重新+1即可(不能>1的原因，n%9最大值是8，+2之后就会出现大于9的情况)
/// n%9!=0的时候，n%9=(n-1)%9+1
/// n%9=0的时候，(n-1)%9+1=9
/// 故最终结果使用(n-1)%9+1
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