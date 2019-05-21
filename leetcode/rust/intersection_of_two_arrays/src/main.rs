// Accepted
// 2019-01-07 13:03
// @lc id=349 lang=rust
// problem: intersection_of_two_arrays
use std::collections::HashSet;
struct Solution {}
impl Solution {
    pub fn intersection(nums1: Vec<i32>, nums2: Vec<i32>) -> Vec<i32> {
        nums1
            .into_iter()
            .collect::<HashSet<_>>()
            .intersection(&nums2.into_iter().collect::<HashSet<_>>())
            .cloned()
            .collect::<Vec<_>>()
    }
}

fn main() {
    println!("{:?}", Solution::intersection(vec![1, 2, 2, 1], vec![2, 2]));
}
