// 2018-12-28 14:31
// @lc id=763 lang=rust
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn partition_labels(s: String) -> Vec<i32> {
        let mut ret = vec![];
        let mut mark = HashMap::new();
        s.chars().rev().enumerate().for_each(|(i, c)| {
            mark.entry(c).or_insert(s.len() - i - 1);
        });
        let mut max_offset = 0;
        let mut begin_offset = 0 as usize;
        s.chars().enumerate().for_each(|(i, c)| {
            if i > max_offset {
                ret.push((i - begin_offset) as i32);
                begin_offset = i;
            }
            max_offset = std::cmp::max(max_offset, *mark.get(&c).unwrap());
        });
        if s.len() != begin_offset {
            ret.push((s.len() - begin_offset) as i32);
        }
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::partition_labels("ababcbacadefegdehijhklij".to_owned())
    );
}
