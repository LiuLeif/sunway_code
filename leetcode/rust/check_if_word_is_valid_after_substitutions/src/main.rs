// Accepted
// @lc id=1003 lang=rust
// problem: check_if_word_is_valid_after_substitutions
struct Solution {}
impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack = vec![];
        for c in s.chars() {
            match c {
                'c' => {
                    if stack.pop().unwrap_or_default() != 'b' {
                        return false;
                    }
                    if stack.pop().unwrap_or_default() != 'a' {
                        return false;
                    }
                }
                _ => stack.push(c),
            }
        }
        stack.is_empty()
    }
}

fn main() {
    println!("{:?}", Solution::is_valid("ababc".to_owned()));
}
