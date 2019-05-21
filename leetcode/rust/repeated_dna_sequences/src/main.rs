// Accepted
// @lc id=187 lang=rust
// problem: repeated_dna_sequences
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn find_repeated_dna_sequences(s: String) -> Vec<String> {
        if s.len() < 10 {
            return vec![];
        }
        let mut ret = vec![];
        let s = s.as_bytes();
        let mut counter = HashMap::new();
        let mut bits = HashMap::new();
        bits.insert(b'A', 0);
        bits.insert(b'G', 1);
        bits.insert(b'C', 2);
        bits.insert(b'T', 3);

        let mut hash = 0;
        for i in 0..10 {
            hash = hash << 2 | bits[&s[i]];
        }
        counter.insert(hash, 1);
        for i in 0..s.len() - 10 {
            hash = (hash << 2) & ((1 << 20) - 1) | bits[&s[i + 10]];
            let count = counter.entry(hash).or_insert(0);
            *count += 1;
            if *count == 2 {
                ret.push(String::from_utf8(s[i + 1..i + 11].to_vec()).unwrap());
            }
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_repeated_dna_sequences("AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT".to_owned())
    );
}
