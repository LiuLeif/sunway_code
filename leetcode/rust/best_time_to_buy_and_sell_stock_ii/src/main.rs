// Accepted
// @lc id=122 lang=rust
// problem: best_time_to_buy_and_sell_stock_ii
struct Solution {}
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        prices
            .windows(2)
            .map(|v| std::cmp::max(v[1] - v[0], 0))
            .sum()
    }
}

fn main() {
    println!("{:?}", Solution::max_profit(vec![7, 1, 5, 3, 6, 4]));
}
