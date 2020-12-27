/// 解法一
/// 暴力破解
//struct Solution {
//}
//impl Solution {
//    pub fn num_teams(rating: Vec<i32>) -> i32 {
//        let mut result:i32 = 0;
//        for i in 0 .. rating.len(){
//            for j in i+1 .. rating.len(){
//                for k in j+1 .. rating.len(){
//                    if (rating[i.clone()] < rating[j.clone()] && rating[j.clone()] < rating[k.clone()]) ||
//                        (rating[i.clone()] > rating[j.clone()] && rating[j.clone()] > rating[k.clone()]) {
//                        result = result + 1;
//                    }
//                }
//            }
//        }
//        result
//    }
//}

/// 解法二
/// 以中间数为主，这个中间数可以组成组合的数量=左边<中间数的数量*右边>中间数的数量+左边>中间数的数量*右边<中间数的数量
struct Solution {
}
impl Solution {
    pub fn num_teams(rating: Vec<i32>) -> i32 {
        let mut result:i32 = 0;
        for i in 1 .. rating.len() - 1{
            let mut left_larger = 0;
            let mut right_larger = 0;
            let mut left_small = 0;
            let mut right_small = 0;

            for j in i .. rating.len(){
                if rating[j] > rating[i]{
                    right_larger = right_larger + 1;
                }
                if rating[j] < rating[i]{
                    right_small = right_small + 1;
                }
            }
            for k in 0 .. i{
                if rating[k] > rating[i]{
                    left_larger = left_larger + 1;
                }
                if rating[k] < rating[i]{
                    left_small = left_small + 1;
                }
            }
            result = result + left_larger * right_small + left_small * right_larger;
        }
        result
    }
}

/// main方法测试
fn main() {
    let rating = vec![2,1,3];
    let result = Solution::num_teams(rating);
    println!("result{}",result);
}