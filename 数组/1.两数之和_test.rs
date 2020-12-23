//pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
//    let mut i: i32 = 0;
//    let mut j: i32 = 0;
//    let mut flag: bool = false;
//    let len: i32 = nums.len() as i32;
//    while i < (len - 1) {
//        j = i + 1;
//        while j < len {
//            if nums[i.clone() as usize] + nums[j as usize] == target {
//                flag = true;
//                break;
//            }
//            j = j.clone() + 1;
//        }
//        if flag {
//            break;
//        }
//
//        i = i.clone() + 1;
//    }
//    vec![i, j]
//}
///
/// 解法2
pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32>{
    let mut i:i32 = 0;
    let len:i32=nums.len() as i32;
    let mut map: HashMap<i32,i32> = HashMap::new();
    while i < len {
        let temp = target.clone() - nums[i as usize];
        if map.contains_key(&temp) {
            return vec![i, *map.get(&temp).unwrap()]
        }
        map.insert(nums[i.clone() as usize], i.clone());
        i = i.clone() + 1;
    }
    vec![i,i]
}