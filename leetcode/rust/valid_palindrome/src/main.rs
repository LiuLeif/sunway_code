// Accepted
// @lc id=125 lang=rust
// problem: valid_palindrome
struct Solution {}
impl Solution {
    pub fn is_palindrome(s: String) -> bool {
        if s.is_empty() {
            return true;
        }
        let s = s.as_bytes();
        let (mut lo, mut hi) = (0, s.len() - 1);
        while lo < hi {
            if !(s[lo] as char).is_alphanumeric() {
                lo += 1;
                continue;
            }
            if !(s[hi] as char).is_alphanumeric() {
                hi -= 1;
                continue;
            }
            if (s[lo] as char).to_ascii_lowercase() != (s[hi] as char).to_ascii_lowercase() {
                return false;
            }
            lo += 1;
            hi -= 1;
        }
        true
    }
}

fn main() {
    println!("{:?}", Solution::is_palindrome("0p".to_owned()));
}
