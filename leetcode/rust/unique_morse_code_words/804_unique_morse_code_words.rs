// 2018-12-28 13:27
// @lc id=804 lang=rust
use std::collections::HashSet;
struct Solution {}
impl Solution {
    pub fn unique_morse_representations(words: Vec<String>) -> i32 {
        let morse = vec![
            ".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..",
            "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-",
            "-.--", "--..",
        ];

        let to_morse = |word: &str| -> String {
            word.chars()
                .fold("".to_owned(), |s, w| s + morse[(w as usize - 'a' as usize)])
        };

        words
            .iter()
            .map(|word| to_morse(word))
            .collect::<HashSet<_>>()
            .len() as i32
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::unique_morse_representations(vec![
            "gin".to_owned(),
            "zen".to_owned(),
            "gig".to_owned(),
            "msg".to_owned(),
        ])
    );
}
