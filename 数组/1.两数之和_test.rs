pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
    let mut i: i32 = 0;
    let mut j: i32 = 0;
    let mut flag: bool = false;
    let len: i32 = nums.len() as i32;
    while i < (len - 1) {
        j = i + 1;
        while j < len {
            if nums[i.clone() as usize] + nums[j as usize] == target {
                flag = true;
                break;
            }
            j = j.clone() + 1;
        }
        if flag {
            break;
        }

        i = i.clone() + 1;
    }
    vec![i, j]
}