// Accepted
// 2019-01-04 10:04
// @lc id=917 lang=rust
// problem: reverse_only_letters
struct Solution {}
impl Solution {
    pub fn reverse_only_letters(s: String) -> String {
        let mut chars: Vec<_> = s.chars().collect();
        let (mut lo, mut hi) = (0 as i32, chars.len() as i32 - 1);
        while lo < hi {
            while lo < hi && !chars[lo as usize].is_ascii_alphabetic() {
                lo += 1;
            }
            while lo < hi && !chars[hi as usize].is_ascii_alphabetic() {
                hi -= 1;
            }
            chars.swap(lo as usize, hi as usize);
            lo += 1;
            hi -= 1;
        }
        return chars.into_iter().collect();
    }
}

fn main() {
    println!("{:?}", Solution::reverse_only_letters("ab-cd".to_owned()));
    println!("{:?}", Solution::reverse_only_letters("".to_owned()));
}
