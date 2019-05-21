// Accepted
// @lc id=20 lang=rust
// problem: valid_parentheses
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mapping = "([{0}])"
            .chars()
            .enumerate()
            .map(|(i, c)| (c, i as i32 - 3))
            .collect::<HashMap<char, i32>>();
        let mut stack = vec![];
        for c in s.chars() {
            let x = mapping[&c];
            let prev = *stack.last().unwrap_or(&0);
            if x > 0 && prev + x == 0 {
                stack.pop();
            } else {
                stack.push(x);
            }
        }
        stack.is_empty()
    }
}

fn main() {
    println!("{:?}", Solution::is_valid("()".to_owned()));
}
