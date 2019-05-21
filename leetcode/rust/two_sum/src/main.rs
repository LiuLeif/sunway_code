// Accepted
// @lc id=1 lang=rust
// problem: two_sum
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut mapper = HashMap::new();
        nums.iter().enumerate().rev().for_each(|(i, &v)| {
            mapper.entry(v).or_insert(i);
        });
        for i in 0..nums.len() {
            let tmp = *mapper.get(&(target - nums[i])).unwrap_or(&i);
            if tmp != i {
                return vec![i as i32, tmp as i32];
            }
        }
        unreachable!();
    }
}

fn main() {
    println!("{:?}", Solution::two_sum(vec![2, 7, 11, 15], 18));
}
