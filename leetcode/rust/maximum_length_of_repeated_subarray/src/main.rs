// Accepted
// @lc id=718 lang=rust
// problem: maximum_length_of_repeated_subarray
struct Solution {}
impl Solution {
    pub fn find_length(a: Vec<i32>, b: Vec<i32>) -> i32 {
        let n = a.len() + 1;
        let mut dp = vec![vec![0; n]; n];
        let mut ret = 0;
        for i in 1..n {
            for j in 1..n {
                if a[i - 1] == b[j - 1] {
                    dp[i][j] = dp[i - 1][j - 1] + 1
                } else {
                    dp[i][j] = 0;
                }
                ret = std::cmp::max(ret, dp[i][j]);
            }
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_length(vec![1, 2, 3, 2, 1], vec![3, 2, 1, 4, 7])
    );
}
