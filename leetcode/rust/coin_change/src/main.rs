// Accepted
struct Solution {}
/*
 * @lc app=leetcode id=322 lang=rust
 *
 * [322] Coin Change
 */

// @lc code=start
impl Solution {
    pub fn coin_change(coins: Vec<i32>, amount: i32) -> i32 {
        let m = coins.len();
        let amount = amount as usize;
        let mut dp = vec![std::i32::MAX; amount + 1];
        dp[0] = 0;
        for i in 1..amount + 1 {
            for j in 0..m {
                let coin = coins[j] as usize;
                if i >= coin && dp[i - coin] != std::i32::MAX {
                    dp[i] = std::cmp::min(dp[i], dp[i - coin] + 1);
                }
            }
        }
        if dp[amount] == std::i32::MAX {
            -1
        } else {
            dp[amount]
        }
    }
}
// @lc code=end

fn main() {
    println!("{:?}", Solution::coin_change(vec![2, 5], 11));
}
