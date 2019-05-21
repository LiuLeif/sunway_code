// Accepted
// @lc id=3 lang=rust
// problem: longest_substring_without_repeating_characters
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn length_of_longest_substring(s: String) -> i32 {
        // abcabcbb
        // abc
        let s = s.as_bytes();
        let (mut lo, mut hi) = (0, 0);
        let mut indexes = HashMap::new();
        let mut ret = 0;
        while hi < s.len() as i32 {
            let prev = indexes.entry(s[hi as usize]).or_insert(-1);
            lo = std::cmp::max(lo, *prev + 1);
            *prev = hi;
            ret = std::cmp::max(ret, hi - lo + 1);
            hi += 1;
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::length_of_longest_substring("pwwkew".to_owned())
    );
}
