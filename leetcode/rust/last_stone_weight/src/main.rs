// Accepted
// @lc id=1046 lang=rust
// problem: last_stone_weight
use std::collections::BinaryHeap;

struct Solution {}
impl Solution {
    pub fn last_stone_weight(stones: Vec<i32>) -> i32 {
        // 1 1 2 4 7 8
        let mut queue = BinaryHeap::new();
        for x in stones {
            queue.push(x);
        }
        while queue.len() > 1 {
            let (a, b) = (queue.pop().unwrap(), queue.pop().unwrap());
            queue.push(a - b);
        }
        return queue.pop().unwrap();
    }
}

fn main() {
    println!("{:?}", Solution::last_stone_weight(vec![2, 7, 4, 1, 8, 1]));
}
