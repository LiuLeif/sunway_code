// Accepted
// @lc id=172 lang=rust
// problem: factorial_trailing_zeroes
struct Solution {}
impl Solution {
    pub fn trailing_zeroes(n: i32) -> i32 {
        // 5 4 3 2 1
        let mut x = 5i64;
        let mut ret = 0;
        while x <= n as i64 {
            ret += n / x as i32;
            x *= 5;
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::trailing_zeroes(1808548329));
}
