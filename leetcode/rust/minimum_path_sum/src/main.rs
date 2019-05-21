// Accepted
// @lc id=64 lang=rust
// problem: minimum_path_sum
struct Solution {}
impl Solution {
    pub fn min_path_sum(grid: Vec<Vec<i32>>) -> i32 {
        if grid.is_empty() {
            return 0;
        }
        let (m, n) = (grid.len(), grid[0].len());
        let mut dp = vec![vec![std::i32::MAX; n + 1]; m + 1];
        dp[1][1] = grid[0][0];
        for i in 1..=m {
            for j in 1..=n {
                if i == 1 && j == 1 {
                    continue;
                }
                dp[i][j] = grid[i - 1][j - 1] + std::cmp::min(dp[i - 1][j], dp[i][j - 1]);
            }
        }
        dp[m][n]
    }
}

fn main() {
    println!("{:?}", Solution::min_path_sum(vec![vec![1]]));
}
