// Accepted
// @lc id=121 lang=rust
// problem: best_time_to_buy_and_sell_stock
struct Solution {}
impl Solution {
    pub fn max_profit(prices: Vec<i32>) -> i32 {
        let mut max = 0;
        let mut ret = 0;
        for i in (0..prices.len()).rev() {
            max = std::cmp::max(max, prices[i]);
            ret = std::cmp::max(ret, max - prices[i])
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::max_profit(vec![7, 1, 5, 3, 6, 4]));
}
