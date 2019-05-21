// Accepted
// 2019-01-07 01:32
// @lc id=292 lang=rust
// problem: nim_game
struct Solution {}
impl Solution {
    pub fn can_win_nim(n: i32) -> bool {
        n % 4 != 0
    }
}

fn main() {
    println!("{:?}", Solution::can_win_nim(4));
}
