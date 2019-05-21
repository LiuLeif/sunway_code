// Accepted
// @lc id=1002 lang=rust
// problem: find_common_characters
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn common_chars(a: Vec<String>) -> Vec<String> {
        let mut counter = [std::i32::MAX; 26];
        let mut ret = vec![];
        for s in a {
            let mut tmp = [0; 26];
            for b in s.bytes() {
                let index = (b - b'a') as usize;
                tmp[index] += 1;
            }
            for i in 0..26 {
                counter[i] = std::cmp::min(counter[i], tmp[i])
            }
        }
        for i in 0..26 {
            if counter[i] != std::i32::MAX {
                for _ in 0..counter[i] {
                    ret.push(((i as u8 + b'a') as char).to_string());
                }
            }
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::common_chars(vec![
            "bella".to_owned(),
            "label".to_owned(),
            "roller".to_owned(),
        ])
    );
}
