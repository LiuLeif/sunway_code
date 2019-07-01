// Accepted
// @lc id=219 lang=rust
// problem: contains_duplicate_ii
use std::collections::HashMap;

struct Solution {}
impl Solution {
    pub fn contains_nearby_duplicate(nums: Vec<i32>, k: i32) -> bool {
        let mut counter = HashMap::new();
        for i in 0..nums.len() {
            if let Some(prev) = counter.get(&nums[i]) {
                if i - prev <= k as usize {
                    return true;
                }
            }
            counter.insert(nums[i], i);
        }
        return false;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::contains_nearby_duplicate(vec![1, 2, 3, 1, 2, 3], 2)
    );
}
