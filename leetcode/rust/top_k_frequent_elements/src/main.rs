// Accepted
// @lc id=347 lang=rust
// problem: top_k_frequent_elements
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn top_k_frequent(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let mut counter = HashMap::new();
        nums.into_iter().for_each(|n| {
            *(counter.entry(n).or_insert(0)) += 1;
        });
        let mut v = counter.iter().collect::<Vec<_>>();
        v.sort_by_key(|t| -t.1);
        v.into_iter().take(k as usize).map(|t| *t.0).collect()
    }
}

fn main() {
    println!("{:?}", Solution::top_k_frequent(vec![1, 1, 1, 2, 2, 3], 2));
}
