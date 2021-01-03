/// 解法一
/// 利用韦恩图，积水面积等与s1加s2-柱子面积-矩形面积
/// 题解
/// https://leetcode-cn.com/problems/trapping-rain-water/solution/wei-en-tu-jie-fa-zui-jian-dan-yi-dong-10xing-jie-j/
struct Solution {
}
impl Solution {
    pub fn trap(height: Vec<i32>) -> i32 {
        let mut max_left:i32 = 0;
        let mut max_right:i32 = 0;
        let mut s1:i32 = 0;
        let mut s2:i32 = 0;
        let mut sum_height:i32 = 0;
        for i in 0 .. height.len() {
            if height[i] > max_left{
                max_left = height[i]
            }
            if height[height.len()-i-1] > max_right{
                max_right = height[height.len()-i-1]
            }

            s1 = s1 + max_left;
            s2 = s2 + max_right;
            sum_height = sum_height + height[i];
        }

        s1 + s2 - sum_height - height.len() as i32 * if max_left >= max_right { max_left } else { max_right }
    }
}

/// main方法测试
fn main() {
    let height = vec![0,1,0,2,1,0,1,3,2,1,2,1];
    let result = Solution::trap(height);
    println!("result = {}",result);
}