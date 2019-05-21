// Accepted
// @lc id=14 lang=rust
// problem: longest_common_prefix
struct Solution {}
impl Solution {
    pub fn longest_common_prefix(strs: Vec<String>) -> String {
        if strs.is_empty() {
            return "".to_owned();
        }
        let len = strs.iter().map(|s| s.len()).min().unwrap();
        let mut ret = 0;
        'again: while ret < len {
            for j in 1..strs.len() {
                if strs[j].as_bytes()[ret] != strs[0].as_bytes()[ret] {
                    break 'again;
                }
            }
            ret += 1;
        }
        return strs[0][..ret].to_owned();
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::longest_common_prefix(vec!["abc".to_owned(), "acd".to_owned(), "x".to_owned(),])
    );
}
