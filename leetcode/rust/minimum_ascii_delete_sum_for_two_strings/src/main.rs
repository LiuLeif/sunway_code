// Accepted
// @lc id=712 lang=rust
// problem: minimum_ascii_delete_sum_for_two_strings
struct Solution {}
impl Solution {
    pub fn minimum_delete_sum(s1: String, s2: String) -> i32 {
        // delete
        // leet
        // let
        let (m, n) = (s1.len(), s2.len());
        let mut dp = vec![vec![0; n + 1]; m + 1];

        let s1 = s1.into_bytes();
        let s2 = s2.into_bytes();

        for i in 1..m + 1 {
            for j in 1..n + 1 {
                if s1[i - 1] == s2[j - 1] {
                    dp[i][j] = dp[i - 1][j - 1] + s1[i - 1] as i32;
                } else {
                    dp[i][j] = std::cmp::max(dp[i - 1][j], dp[i][j - 1])
                }
            }
        }

        s1.into_iter().chain(s2).map(|b| b as i32).sum::<i32>() - 2 * dp[m][n]
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::minimum_delete_sum("delete".to_owned(), "leet".to_owned())
    );
}
