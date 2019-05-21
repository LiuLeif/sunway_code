// 2018-12-24 17:18
// @lc id=950 lang=rust
use std::collections::VecDeque;
struct Solution {}
impl Solution {
    pub fn deck_revealed_increasing(deck: Vec<i32>) -> Vec<i32> {
        let mut ret = vec![0; deck.len()];
        let mut deck = deck;
        deck.sort();
        // 2,3,5,7,11,13,17
        let mut q = VecDeque::new();
        for i in 0..deck.len() {
            q.push_back(i)
        }
        let mut index = 0;
        loop {
            if q.is_empty() {
                break;
            }
            let top = q.pop_front().unwrap();
            ret[top] = deck[index];
            index += 1;

            if q.is_empty() {
                break;
            }
            let top = q.pop_front().unwrap();
            q.push_back(top);
        }
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::deck_revealed_increasing(vec![17, 13, 11, 2, 3, 5, 7])
    );
}
