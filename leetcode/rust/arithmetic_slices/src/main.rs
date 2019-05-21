// Accepted
// @lc id=413 lang=rust
// problem: arithmetic_slices
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn number_of_arithmetic_slices(a: Vec<i32>) -> i32 {
        let mut dp = vec![HashMap::<i32, i32>::new(); a.len()];
        for i in 1..a.len() {
            let delta = a[i] - a[i - 1];
            let prev = *(dp[i - 1].entry(delta).or_insert(1));
            *dp[i].entry(delta).or_insert(0) += prev + 1;
        }
        dp.iter()
            .flat_map(|m| m.values().filter(|&&v| v >= 3).map(|x| x - 2))
            .sum()
    }
}
fn main() {
    println!(
        "{:?}",
        Solution::number_of_arithmetic_slices(vec![1, 2, 3, 4, 5, 6])
    );
}
