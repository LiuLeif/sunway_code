// 2018-12-26 17:24
// @lc id=890 lang=rust
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn find_and_replace_pattern(words: Vec<String>, pattern: String) -> Vec<String> {
        fn convert(x: &str) -> String {
            let mut m = HashMap::<char, char>::new();
            let mut curr = 'a' as u8;
            x.chars().for_each(|c| {
                if !m.contains_key(&c) {
                    m.insert(c, curr as char);
                    curr += 1;
                }
            });
            return x.chars().map(|c| m.get(&c).unwrap()).collect();
        }

        let pattern = convert(&pattern);
        return words
            .iter()
            .filter(|&w| convert(w) == pattern)
            .map(|x| x.clone())
            .collect();
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_and_replace_pattern(
            vec![
                "abc".to_owned(),
                "deq".to_owned(),
                "mee".to_owned(),
                "aqq".to_owned(),
                "dkd".to_owned(),
                "ccc".to_owned(),
            ],
            "abb".to_owned()
        )
    );
}
