/// 解法一
/// 求中位数：排序只排到中位数即可，长度是偶数就将中间两位数求和除以2返回,长度是奇数就将中间1位数返回
struct Solution {
}
impl Solution {

    pub fn find_median_sorted_arrays(nums1: Vec<i32>, nums2: Vec<i32>) -> f64 {
        if nums1.len()+nums2.len() == 0{
            return 0.0;
        }
        let is_even:bool = (nums1.len()+nums2.len())%2==0;
        let len =  (nums1.len()+nums2.len())/2+1;
        let mut nums = vec![0;len];
        let mut i:usize = 0;
        let mut j:usize = 0;
        while &i+&j<len {
            if i >= nums1.len(){
                nums[i+j] = nums2[j];
                j=j+1;
            } else if j >= nums2.len() {
                nums[i+j] = nums1[i];
                i=i+1;
            } else if nums1[i]<nums2[j] {
                nums[i+j] = nums1[i];
                i = i + 1;
            } else if nums1[i]>nums2[j] {
                nums[i+j] = nums2[j];
                j = j + 1;
            } else {
                nums[i+j] = nums1[i];
                i = i + 1;
                if i+j >= len {
                    break;
                }
                nums[i+j] = nums2[j];
                j = j + 1;
            }
        }
        if is_even { (nums[len-1] + nums[len-2]) as f64 / 2.0} else { nums[len-1] as f64}
    }
}

/// main方法测试
fn main() {
    let nums1=vec![0,0];
    let nums2=vec![0,0];
    let result:f64 = Solution::find_median_sorted_arrays(nums1, nums2);
    println!("result = {}",result);
}