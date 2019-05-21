// Accepted
// @lc id=521 lang=rust
// problem: longest_uncommon_subsequence_i
struct Solution {}
impl Solution {
    pub fn find_lu_slength(a: String, b: String) -> i32 {
        if a == b {
            return -1;
        }
        return std::cmp::max(a.len(), b.len()) as i32;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_lu_slength("aba".to_owned(), "cdc".to_owned());
    );
}
