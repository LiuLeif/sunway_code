// Accepted
// 2019-01-07 17:08
// @lc id=451 lang=rust
// problem: sort_characters_by_frequency
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn frequency_sort(s: String) -> String {
        let mut counter = HashMap::new();
        s.chars().for_each(|c| {
            *(counter.entry(c).or_insert(0)) += 1;
        });
        let mut v = counter.into_iter().collect::<Vec<_>>();
        v.sort_by_key(|v| 0 - v.1 as i32);
        v.iter().map(|v| v.0.to_string().repeat(v.1)).collect()
    }
}

fn main() {
    println!("{:?}", Solution::frequency_sort("Aabb".to_owned()));
}
