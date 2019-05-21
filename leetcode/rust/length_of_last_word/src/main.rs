// Accepted
// @lc id=58 lang=rust
// problem: length_of_last_word
struct Solution {}
impl Solution {
    pub fn length_of_last_word(s: String) -> i32 {
        s.trim().split(' ').last().map_or(0, |x| x.len() as i32)
    }
}

fn main() {
    println!("{:?}", Solution::length_of_last_word("a".to_owned()));
}
