// Accepted
// @lc id=873 lang=rust
// problem: length_of_longest_fibonacci_subsequence
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn len_longest_fib_subseq(a: Vec<i32>) -> i32 {
        // 1 2 3 4 5 6 7 8
        // 1 2 3 5 8
        let mut dp = vec![HashMap::new(); a.len()];
        let mut ret = 0;
        for i in 1..a.len() {
            for j in 0..i {
                let update = *dp[j].entry(a[i] - a[j]).or_insert(1) + 1;
                let curr = dp[i].entry(a[j]).or_insert(update);
                *curr = std::cmp::max(*curr, update);
                ret = std::cmp::max(ret, *curr);
            }
        }
        if ret >= 3 {
            return ret;
        }
        0
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::len_longest_fib_subseq(vec![1, 2, 3, 4, 5, 6, 7, 8])
    );
    println!("{:?}", Solution::len_longest_fib_subseq(vec![1, 3, 5]));
}
