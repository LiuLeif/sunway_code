// Accepted
// @lc id=231 lang=rust
// problem: power_of_two
struct Solution {}
impl Solution {
    pub fn is_power_of_two(n: i32) -> bool {
        n > 0 && (n & (n - 1) == 0)
    }
}

fn main() {
    println!("{:?}", Solution::is_power_of_two(-4));
}
