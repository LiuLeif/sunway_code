// Accepted
// @lc id=383 lang=rust
// problem: ransom_note
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn can_construct(ransom_note: String, magazine: String) -> bool {
        let mut a_counter = HashMap::new();
        let mut b_counter = HashMap::new();
        ransom_note.chars().for_each(|c| {
            *a_counter.entry(c).or_insert(0) += 1;
        });
        magazine.chars().for_each(|c| {
            *b_counter.entry(c).or_insert(0) += 1;
        });
        a_counter
            .iter()
            .all(|(k, v)| b_counter.get(k).unwrap_or(&0) >= v)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::can_construct("aa".to_owned(), "aab".to_owned())
    );
}
