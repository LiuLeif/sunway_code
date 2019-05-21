// Accepted
// 2019-01-07 14:14
// @lc id=389 lang=rust
// problem: find_the_difference
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn find_the_difference(s: String, t: String) -> char {
        let mut m1 = HashMap::new();
        s.chars().for_each(|c| {
            *(m1.entry(c).or_insert(0)) += 1;
        });
        let mut m2 = HashMap::new();
        t.chars().for_each(|c| {
            *(m2.entry(c).or_insert(0)) += 1;
        });
        for (k, v) in m2 {
            match m1.get(&k) {
                Some(n) if n == &v => (),
                _ => return k,
            }
        }
        return 'x';
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_the_difference("abcd".to_owned(), "abcde".to_owned())
    );
    println!(
        "{:?}",
        Solution::find_the_difference("a".to_owned(), "aa".to_owned())
    );
}
