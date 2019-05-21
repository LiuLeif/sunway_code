// Accepted
// @lc id=446 lang=rust
// problem: arithmetic_slices_ii_subsequence
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn number_of_arithmetic_slices(a: Vec<i32>) -> i32 {
        let mut dp = vec![HashMap::<i64, i32>::new(); a.len()];
        let mut ret = 0;
        for i in 1..a.len() {
            for j in 0..i {
                let delta = a[i] as i64 - a[j] as i64;
                let prev = *(dp[j].entry(delta).or_insert(0));
                *dp[i].entry(delta).or_insert(0) += prev + 1;
                ret += prev;
            }
        }
        println!("{:?}", dp);
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::number_of_arithmetic_slices(vec![2, 2, 2, 4, 6])
    );
    println!(
        "{:?}",
        Solution::number_of_arithmetic_slices(vec![1, 1, 1, 1, 1])
    );
}
