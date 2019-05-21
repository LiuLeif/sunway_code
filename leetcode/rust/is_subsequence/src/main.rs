// Accepted
// @lc id=392 lang=rust
// problem: is_subsequence
struct Solution {}
impl Solution {
    pub fn is_subsequence(s: String, t: String) -> bool {
        let s = s.as_bytes();
        let t = t.as_bytes();
        let (mut a, mut b) = (0, 0);
        while a < s.len() && b < t.len() {
            if s[a] == t[b] {
                a += 1;
            }
            b += 1;
        }
        a == s.len()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::is_subsequence("abc".to_owned(), "abcc".to_owned())
    );
}
