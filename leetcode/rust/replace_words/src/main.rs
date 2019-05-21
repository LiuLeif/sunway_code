// Accepted
// @lc id=648 lang=rust
// problem: replace_words
use std::collections::HashMap;
struct Solution {}

#[derive(Default)]
struct TrieNode {
    value: String,
    children: HashMap<char, TrieNode>,
}

impl TrieNode {
    fn add(&mut self, s: String) {
        let mut root = self;
        for c in s.chars() {
            root = root.children.entry(c).or_default();
        }
        root.value = s;
    }

    fn query(&self, s: String) -> String {
        let mut root = self;
        for c in s.chars() {
            if !root.value.is_empty() {
                return root.value.clone();
            }
            match root.children.get(&c) {
                Some(node) => root = node,
                None => return s,
            }
        }
        s
    }
}

impl Solution {
    pub fn replace_words(dict: Vec<String>, sentence: String) -> String {
        let mut trie = TrieNode::default();
        for s in dict {
            trie.add(s);
        }
        sentence
            .split(' ')
            .map(|s| trie.query(s.to_owned()))
            .collect::<Vec<String>>()
            .join(" ")
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::replace_words(
            vec!["cat".to_owned(), "bat".to_owned(), "rat".to_owned(),],
            "the cattle was rattled by the battery".to_owned()
        )
    );
}
