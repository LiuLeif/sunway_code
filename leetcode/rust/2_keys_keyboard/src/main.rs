// Accepted
// @lc id=650 lang=rust
// problem: 2_keys_keyboard
struct Solution {}
impl Solution {
    pub fn min_steps(n: i32) -> i32 {
        // AA AA AA A
        // 24 [2,12][3,8],[4,6]
        let n = n as usize;
        let mut dp = vec![0; n + 1];
        dp[1] = 0;
        for i in 2..=n {
            dp[i] = i;
            for j in 2..i {
                if i % j == 0 {
                    dp[i] = std::cmp::min(dp[i], dp[j] + i / j);
                    dp[i] = std::cmp::min(dp[i], dp[i / j] + j);
                }
            }
        }
        return dp[n] as i32;
    }
}

fn main() {
    println!("{:?}", Solution::min_steps(2));
}
