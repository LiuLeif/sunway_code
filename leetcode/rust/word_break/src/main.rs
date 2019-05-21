// Accepted
// @lc id=139 lang=rust
// problem: word_break
use std::collections::HashSet;
struct Solution {}
impl Solution {
    pub fn word_break(s: String, word_dict: Vec<String>) -> bool {
        let counter = word_dict.into_iter().collect::<HashSet<String>>();
        let mut dp = vec![false; s.len() + 1];
        dp[0] = true;
        for i in 0..s.len() + 1 {
            for j in 0..i + 1 {
                if dp[j] && counter.contains(&s[j..i]) {
                    dp[i] = true;
                    break;
                }
            }
        }
        dp[s.len()]
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::word_break(
            "leetcode".to_owned(),
            vec!["leet".to_owned(), "code".to_owned()]
        )
    );
}
