// Accepted
struct Solution {}
/*
 * @lc app=leetcode id=1217 lang=rust
 *
 * [1217] Play with Chips
 */

// @lc code=start
impl Solution {
    pub fn min_cost_to_move_chips(chips: Vec<i32>) -> i32 {
        let n = chips.iter().filter(|&x| x & 0x1 == 0).count();
        std::cmp::min(n, chips.len() - n) as i32
    }
}
// @lc code=end

fn main() {
    println!("{:?}", Solution::min_cost_to_move_chips(vec![1, 2, 3]));
}
