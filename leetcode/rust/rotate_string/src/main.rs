// Accepted
// @lc id=796 lang=rust
// problem: rotate_string
struct Solution {}
impl Solution {
    pub fn rotate_string(a: String, b: String) -> bool {
        if a.len() != b.len() {
            return false;
        }
        (a.repeat(2)).contains(&b)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::rotate_string("abcde".to_owned(), "cdeab".to_owned())
    );
}
