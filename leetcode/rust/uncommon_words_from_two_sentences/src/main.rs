// 2019-01-03 11:16
// @lc id=884 lang=rust
// problem: uncommon_words_from_two_sentences
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn uncommon_from_sentences(a: String, b: String) -> Vec<String> {
        let mut counter = HashMap::new();
        a.split_whitespace()
            .chain(b.split_whitespace())
            .for_each(|w| {
                *counter.entry(w).or_insert(0) += 1;
            });

        counter
            .into_iter()
            .filter(|&e| e.1 == 1)
            .map(|e| {
                (e.0).to_owned()
            })
            .collect::<Vec<_>>()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::uncommon_from_sentences("apple apple".to_owned(), "banana".to_owned())
    );
}
