// Accepted
struct Solution {}
/*
 * @lc app=leetcode id=1207 lang=rust
 *
 * [1207] Unique Number of Occurrences
 */

// @lc code=start
use std::collections::HashMap;
use std::collections::HashSet;
impl Solution {
    pub fn unique_occurrences(arr: Vec<i32>) -> bool {
        let mut counter = HashMap::new();
        arr.into_iter().for_each(|x| {
            *counter.entry(x).or_insert(0) += 1;
        });

        let values = counter.values().collect::<HashSet<_>>();
        values.len() == counter.values().len()
    }
}
// @lc code=end

fn main() {
    println!("{:?}", Solution::unique_occurrences(vec![1, 2, 2, 1, 1, 3]));
}
