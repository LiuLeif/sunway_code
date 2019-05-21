// Accepted
// @lc id=127 lang=rust
// problem: word_ladder
use std::collections::HashSet;
use std::collections::VecDeque;
struct Solution {}
impl Solution {
    pub fn ladder_length(begin_word: String, end_word: String, word_list: Vec<String>) -> i32 {
        let mut word_set = word_list
            .into_iter()
            .map(|s| s.into_bytes())
            .collect::<HashSet<_>>();

        let mut queue = VecDeque::new();
        let (begin_word, end_word) = (begin_word.into_bytes(), end_word.into_bytes());
        queue.push_back(begin_word);
        let mut count = 1;
        while !queue.is_empty() {
            count += 1;
            for _ in 0..queue.len() {
                let mut top = queue.pop_front().unwrap();
                for i in 0..top.len() {
                    let orig_c = top[i];
                    for c in b'a'..b'z' {
                        if top[i] == c {
                            continue;
                        }
                        top[i] = c;
                        if !word_set.contains(&top) {
                            continue;
                        }
                        word_set.remove(&top);
                        queue.push_back(top.clone());
                        if top == end_word {
                            return count;
                        }
                    }
                    top[i] = orig_c;
                }
            }
        }
        return 0;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::ladder_length(
            "hit".to_owned(),
            "cog".to_owned(),
            vec![
                "hot".to_owned(),
                "dot".to_owned(),
                "dog".to_owned(),
                "lot".to_owned(),
                "log".to_owned(),
                "cog".to_owned(),
            ]
        )
    );
}
