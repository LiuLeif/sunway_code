// 2019-01-02 17:39
// @lc id=791 lang=rust
// problem: custom_sort_string
use std::collections::HashMap;

struct Solution {}
impl Solution {
    pub fn custom_sort_string(s: String, t: String) -> String {
        let mark: HashMap<char, usize> = s.chars().zip(0..s.len()).collect();
        let mut chars: Vec<_> = t.chars().collect();
        chars.sort_by_key(|c| match mark.get(c) {
            Some(&x) => x,
            _ => 0,
        });
        return chars.iter().collect::<String>();
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::custom_sort_string("cba".to_owned(), "abcd".to_owned())
    );
}
