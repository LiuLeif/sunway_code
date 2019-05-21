// Accepted
// @lc id=28 lang=rust
// problem: implement_str_str
struct Solution {}
impl Solution {
    pub fn str_str(haystack: String, needle: String) -> i32 {
        if needle.len() > haystack.len() {
            return -1;
        }
        let haystack = haystack.as_bytes();
        let needle = needle.as_bytes();

        'again: for i in 0..haystack.len() - needle.len() + 1 {
            for j in 0..needle.len() {
                if haystack[i + j] != needle[j] {
                    continue 'again;
                }
            }
            return i as i32;
        }
        return -1;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::str_str("hello".to_owned(), "lll".to_owned())
    );
}
