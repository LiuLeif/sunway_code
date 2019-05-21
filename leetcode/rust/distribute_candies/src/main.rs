// 2019-01-03 14:40
// @lc id=575 lang=rust
// problem: distribute_candies
struct Solution {}
impl Solution {
    pub fn distribute_candies(candies: Vec<i32>) -> i32 {
        let n = candies.len() / 2;
        let mut candies = candies;
        candies.sort();
        candies.dedup();
        std::cmp::min(n, candies.len()) as i32
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::distribute_candies(vec![
            100000, 0, 100000, 0, 100000, 0, 100000, 0, 100000, 0, 100000, 0
        ])
    );
}
