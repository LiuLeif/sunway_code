// Accepted
// @lc id=1021 lang=rust
// problem: remove_outermost_parentheses
struct Solution {}
impl Solution {
    pub fn remove_outer_parentheses(s: String) -> String {
        let mut stack = vec![];
        let mut start = 0;
        let s = s.as_bytes();
        let mut ret = "".to_owned();
        for i in 0..s.len() {
            if s[i] == b'(' {
                if stack.is_empty() {
                    start = i;
                }
                stack.push(s[i]);
            } else {
                stack.pop();
                if stack.is_empty() {
                    ret.push_str(
                        String::from_utf8(s[start + 1..i].to_vec())
                            .unwrap()
                            .as_str(),
                    );
                }
            }
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::remove_outer_parentheses("(()())(())(()(()))".to_owned())
    );
}
