// Accepted
// @lc id=326 lang=rust
// problem: power_of_three
struct Solution {}
impl Solution {
    pub fn is_power_of_three(n: i32) -> bool {
        n > 0 && 3i32.pow(19) % n == 0
    }
}

fn main() {
    println!("{:?}", Solution::is_power_of_three(8));
}
