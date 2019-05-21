// 2018-12-28 17:52
// @lc id=344 lang=rust
struct Solution {}
impl Solution {
    pub fn reverse_string(s: String) -> String {
        s.chars().rev().collect()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::reverse_string("A man, a plan, a canal: Panama".to_owned())
    );
}
