// 2018-12-25 14:40
// @lc id=921 lang=rust
use std::collections::VecDeque;

struct Solution {}
impl Solution {
    pub fn min_add_to_make_valid(s: String) -> i32 {
        fn is_pair(a: char, b: char) -> bool {
            match (a, b) {
                ('(', ')') => true,
                _ => false,
            }
        }

        let mut stack = VecDeque::new();
        for c in s.chars() {
            if stack.is_empty() || !is_pair(*(stack.back().unwrap()), c) {
                stack.push_back(c);
            } else {
                stack.pop_back();
            }
        }
        return stack.len() as i32;
    }
}

fn main() {
    println!("{:?}", Solution::min_add_to_make_valid("()))((".to_owned()));
    println!("{:?}", Solution::min_add_to_make_valid("(()".to_owned()));
}
