// Accepted
// @lc id=49 lang=rust
// problem: group_anagrams
use std::collections::HashMap;
struct Solution {}
impl Solution {
    fn count(str: &str) -> [i32; 26] {
        let mut ret = [0; 26];
        str.bytes().for_each(|b| {
            ret[(b - b'a') as usize] += 1;
        });
        ret
    }

    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut counter = HashMap::new();
        for s in strs {
            let array = Solution::count(&s);
            let v = counter.entry(array).or_insert(vec![]);
            v.push(s);
        }
        counter.values().cloned().collect()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::group_anagrams(vec![
            "eat".to_owned(),
            "tea".to_owned(),
            "tan".to_owned(),
            "ate".to_owned(),
            "nat".to_owned(),
            "bat".to_owned(),
        ])
    );
}
