// Accepted
// @lc id=342 lang=rust
// problem: power_of_four
struct Solution {}
impl Solution {
    pub fn is_power_of_four(num: i32) -> bool {
        num > 0 && num & (num - 1) == 0 && num % 3 == 1
    }
}

fn main() {
    println!("{:?}", Solution::is_power_of_four(4));
    println!("{:?}", Solution::is_power_of_four(16));
    println!("{:?}", Solution::is_power_of_four(8));
}
