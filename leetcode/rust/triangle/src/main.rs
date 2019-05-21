// Accepted
// @lc id=120 lang=rust
// problem: triangle
struct Solution {}
impl Solution {
    pub fn minimum_total(triangle: Vec<Vec<i32>>) -> i32 {
        if triangle.is_empty() {
            return 0;
        }
        let mut dp = vec![0; triangle.len()];
        let mut dp2 = vec![0; triangle.len()];
        dp[0] = triangle[0][0];
        for row in &triangle[1..] {
            for i in 0..row.len() {
                let x = {
                    if i == 0 {
                        std::i32::MAX
                    } else {
                        dp[i - 1]
                    }
                };
                let y = {
                    if i == row.len() - 1 {
                        std::i32::MAX
                    } else {
                        dp[i]
                    }
                };
                dp2[i] = std::cmp::min(x, y) + row[i];
            }
            std::mem::swap(&mut dp, &mut dp2);
        }
        dp.into_iter().min().unwrap()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::minimum_total(vec![vec![2], vec![3, 4], vec![6, 5, 7], vec![4, 1, 8, 3]])
    );
}
