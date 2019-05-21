// Accepted
// @lc id=221 lang=rust
// problem: maximal_square
struct Solution {}
impl Solution {
    pub fn maximal_square(matrix: Vec<Vec<char>>) -> i32 {
        if matrix.is_empty() {
            return 0;
        }
        let (m, n) = (matrix.len() + 1, matrix[0].len() + 1);
        let mut dp = vec![vec![0; n]; m];
        let mut ret = 0;
        for i in 1..m {
            for j in 1..n {
                if matrix[i - 1][j - 1] == '0' {
                    continue;
                }
                dp[i][j] = vec![dp[i - 1][j - 1], dp[i][j - 1], dp[i - 1][j]]
                    .into_iter()
                    .min()
                    .unwrap()
                    + 1;
                ret = std::cmp::max(dp[i][j], ret);
            }
        }
        ret * ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::maximal_square(vec![
            vec!['1', '0', '1', '0', '0'],
            vec!['1', '0', '1', '1', '1'],
            vec!['1', '1', '1', '1', '1'],
            vec!['1', '0', '0', '1', '0'],
        ])
    );
}
