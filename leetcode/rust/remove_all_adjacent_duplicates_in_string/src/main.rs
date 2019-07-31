// Accepted
// @lc id=1047 lang=rust
// problem: remove_all_adjacent_duplicates_in_string
struct Solution {}
impl Solution {
    pub fn remove_duplicates(s: String) -> String {
        let mut stack = vec![];
        for c in s.chars() {
            if stack.last().unwrap_or(&'0') == &c {
                stack.pop();
            } else {
                stack.push(c);
            }
        }
        stack.into_iter().collect::<String>()
    }
}

fn main() {
    println!("{:?}", Solution::remove_duplicates("abbaca".to_string()));
}
