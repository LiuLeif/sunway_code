// 2019-01-03 09:46
// @lc id=500 lang=rust
// problem: keyboard_row
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn find_words(words: Vec<String>) -> Vec<String> {
        let mut mapping: HashMap<char, i32> = HashMap::new();
        "qwertyuiop".chars().for_each(|c| {
            mapping.insert(c, 1);
        });
        "asdfghjkl".chars().for_each(|c| {
            mapping.insert(c, 2);
        });
        "zxcvbnm".chars().for_each(|c| {
            mapping.insert(c, 3);
        });
        words
            .iter()
            .filter(|w| {
                let w = w.to_lowercase();
                let id = *mapping.get(&w.chars().next().unwrap()).unwrap();
                w.chars().all(|c| mapping[&c] == id)
            })
            .map(|w| w.clone())
            .collect()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_words(vec![
            "Hello".to_owned(),
            "Alaska".to_owned(),
            "Dad".to_owned(),
            "Peace".to_owned(),
        ])
    );
}
