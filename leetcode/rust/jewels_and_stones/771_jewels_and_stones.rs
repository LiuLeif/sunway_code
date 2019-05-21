// 2018-12-21 15:03
// @lc id=771 lang=rust
use std::collections::HashSet;

struct Solution {}
impl Solution {
    pub fn num_jewels_in_stones(j: String, s: String) -> i32 {
        let set = j.bytes().collect::<HashSet<_>>();
        return s.bytes().filter(|b| set.contains(b)).count() as i32;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::num_jewels_in_stones("aA".to_owned(), "aAaA".to_owned())
    );
}
