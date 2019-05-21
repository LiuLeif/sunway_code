// Accepted
// 2019-01-07 12:57
// @lc id=258 lang=rust
// problem: add_digits
struct Solution {}
impl Solution {
    pub fn add_digits(num: i32) -> i32 {
        // 38 11 2
        // 9
        return (num - 1) % 9 + 1;
    }
}

fn main() {
    println!("{:?}", Solution::add_digits(111));
}
