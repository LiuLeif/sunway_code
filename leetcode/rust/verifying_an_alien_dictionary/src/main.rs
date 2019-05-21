// Accepted
// 2019-01-04 18:24
// @lc id=953 lang=rust
// problem: verifying_an_alien_dictionary
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn is_alien_sorted(words: Vec<String>, order: String) -> bool {
        let mapping = order
            .chars()
            .zip("abcdefghijklmnopqrstuvwxyz".chars())
            .collect::<HashMap<_, _>>();
        let converted = words
            .into_iter()
            .map(|w| {
                w.chars()
                    .map(|c| mapping.get(&c).unwrap())
                    .collect::<String>()
            })
            .collect::<Vec<_>>();
        converted.windows(2).all(|w| w[0] < w[1])
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::is_alien_sorted(
            vec!["hello".to_owned(), "leetcode".to_owned()],
            "hlabcdefgijkmnopqrstuvwxyz".to_owned()
        )
    );
}
