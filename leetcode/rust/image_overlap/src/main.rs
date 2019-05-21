// Accepted
// @lc id=835 lang=rust
// problem: image_overlap
struct Solution {}
use std::collections::HashMap;
impl Solution {
    pub fn largest_overlap(a: Vec<Vec<i32>>, b: Vec<Vec<i32>>) -> i32 {
        let a = {
            let mut ret = vec![];
            for i in 0..a.len() {
                for j in 0..a[0].len() {
                    if a[i][j] == 1 {
                        ret.push((i as i32, j as i32))
                    }
                }
            }
            ret
        };
        let b = {
            let mut ret = vec![];
            for i in 0..b.len() {
                for j in 0..b[0].len() {
                    if b[i][j] == 1 {
                        ret.push((i as i32, j as i32))
                    }
                }
            }
            ret
        };
        let mut counter = HashMap::new();
        for (x0, y0) in a {
            for (x1, y1) in &b {
                *counter.entry((x1 - x0, y1 - y0)).or_insert(0) += 1;
            }
        }
        match counter.iter().map(|t| *t.1).max() {
            Some(x) => x,
            None => 0,
        }
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::largest_overlap(
            vec![vec![1, 1, 0], vec![0, 1, 0], vec![0, 1, 0],],
            vec![vec![0, 0, 0], vec![0, 1, 1], vec![0, 0, 1],]
        )
    );
    println!(
        "{:?}",
        Solution::largest_overlap(vec![vec![0]], vec![vec![0]],)
    );
}
