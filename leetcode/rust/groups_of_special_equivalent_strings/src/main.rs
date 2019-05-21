// Accepted
// @lc id=893 lang=rust
// problem: groups_of_special_equivalent_strings
struct Solution {}
use std::collections::HashSet;
impl Solution {
    pub fn num_special_equiv_groups(a: Vec<String>) -> i32 {
        let mut counter = HashSet::new();
        for s in a {
            let mut a = s.chars().step_by(2).collect::<Vec<char>>();
            a.sort();
            let mut b = s.chars().skip(1).step_by(2).collect::<Vec<char>>();
            b.sort();
            let x = a.into_iter().collect::<String>() + &b.into_iter().collect::<String>();
            counter.insert(x);
        }
        counter.len() as i32
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::num_special_equiv_groups(vec![
            "abcd".to_owned(),
            "cdab".to_owned(),
            "adcb".to_owned(),
            "cbad".to_owned(),
        ])
    );
}
