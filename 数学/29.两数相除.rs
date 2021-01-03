/// 解法一
/// 被除数可以减除数多少次，这个次数即为被除数除以除数的结果
/// 基于被除数可能太大以及，除数太小可以采用二分查找的思想，每一次都先确定前一半
/// 如20和3，首先3+3=6，6<20，此时结果为2，然后6+6=12，12<20，此时结果为4，然后12+12=24，24>20，此时确定前一半结果为4
/// 开始找被除数20-12=8，除数是3的结果，首先3+3=6，6<8，此时结果为2，然后6+6=12，12>8，此时确定前一半结果为2，
/// 加上上一轮的结果4，此时总结果为6
/// 开始找被除数8-6=2，除数是3的结果，因被除数本身就比3小，故不继续查找，最终结果为6
//struct Solution {
//}
//impl Solution {
//    pub fn divide(dividend: i32, divisor: i32) -> i32 {
//        let is_positive:bool = (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0);
//        let mut result:i32 = 0;
//        let mut dividend_copy:i32 = if dividend < 0 { dividend.clone() } else { -dividend.clone() };
//        let mut divisor_copy:i32 =  0;
//        loop {
//            let mut tempResult:i32 = -1;
//            dividend_copy = dividend_copy - divisor_copy;
//            divisor_copy =  if divisor < 0 { divisor.clone() } else { -divisor.clone() };
//            if dividend_copy >=0 || dividend_copy - divisor_copy > 0 {
//                break;
//            }
//            loop {
//                if (dividend_copy - divisor_copy) > divisor_copy {
//                    break;
//                }
//                tempResult = tempResult + tempResult;
//                divisor_copy = divisor_copy + divisor_copy;
//            }
//
//            result = result + tempResult;
//        }
//        if is_positive { if result <= -2147483648{ 2147483647 } else { -result } } else { result }
//
//    }
//}

/// 解法二
/// 可用竖式计算



/// main方法测试
fn main() {
    let dividend = -2147483648;
    let divisor = -2147483648;
    let result = Solution::divide(dividend, divisor);
    println!("result = {}",result);
}