// Accepted
// @lc id=554 lang=rust
// problem: brick_wall
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn least_bricks(wall: Vec<Vec<i32>>) -> i32 {
        let n = wall.len();
        let mut dp = HashMap::<i32, i32>::new();
        for row in wall {
            let mut offset = 0;
            for j in &row[..row.len() - 1] {
                offset += j;
                *dp.entry(offset).or_insert(0) -= 1;
            }
        }
        n as i32 + dp.values().min().unwrap_or(&0)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::least_bricks(vec![
            vec![1, 2, 2, 1],
            vec![3, 1, 2],
            vec![1, 3, 2],
            vec![2, 4],
            vec![3, 1, 2],
            vec![1, 3, 1, 1],
        ])
    );
}
