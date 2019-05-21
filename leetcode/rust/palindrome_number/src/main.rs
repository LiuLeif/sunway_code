// Accepted
// @lc id=9 lang=rust
// problem: palindrome_number
struct Solution {}
impl Solution {
    pub fn is_palindrome(x: i32) -> bool {
        if x < 0 {
            return false;
        }
        let mut b = 0;
        let mut a = x;
        while a != 0 {
            b = b * 10 + a % 10;
            a /= 10;
        }
        b == x
    }
}

fn main() {
    println!("{:?}", Solution::is_palindrome(121));
}
