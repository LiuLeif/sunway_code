// 2019-01-03 10:41
// @lc id=557 lang=rust
// problem: reverse_words_in_a_string_iii
struct Solution {}
impl Solution {
    pub fn reverse_words(s: String) -> String {
        s.split_whitespace()
            .map(|w| w.chars().rev().collect::<String>())
            .collect::<Vec<_>>()
            .join(" ")
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::reverse_words("Let's take LeetCode contest".to_owned())
    );
}
