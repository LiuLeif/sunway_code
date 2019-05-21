// @lc id=830 lang=rust
// problem: positions_of_large_groups
struct Solution {}
impl Solution {
    pub fn large_group_positions(s: String) -> Vec<Vec<i32>> {
        let mut len = 0;
        let mut ret = vec![];
        let s = s.as_bytes();
        for i in 1..s.len() {
            if s[i] == s[i - 1] {
                len += 1;
            } else {
                if len >= 3 {
                    ret.push(vec![i - len, i - 1])
                }
                len = 0;
            }
        }
        
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::large_group_positions("abcdddeeeeaabbbcd".to_owned())
    );
}
