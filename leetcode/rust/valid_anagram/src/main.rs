// Accepted
// @lc id=242 lang=rust
// problem: valid_anagram
struct Solution {}
impl Solution {
    pub fn is_anagram(s: String, t: String) -> bool {
        let mut s = s.chars().collect::<Vec<char>>();
        s.sort();
        let mut t = t.chars().collect::<Vec<char>>();
        t.sort();
        s == t
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::is_anagram("anagram".to_owned(), "nagaram".to_owned())
    );
}
