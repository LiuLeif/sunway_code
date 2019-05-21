// Accepted
// @lc id=939 lang=rust
// problem: minimum_area_rectangle
use std::collections::HashSet;
struct Solution {}
impl Solution {
    pub fn min_area_rect(points: Vec<Vec<i32>>) -> i32 {
        let mut counter = HashSet::new();
        for v in points.iter() {
            counter.insert(v);
        }
        let mut ret = std::i32::MAX;
        for x in points.iter() {
            for y in points.iter() {
                if x[0] == y[0] || x[1] == y[1] {
                    continue;
                }
                if counter.contains(&vec![x[0], y[1]]) && counter.contains(&vec![y[0], x[1]]) {
                    ret = std::cmp::min(ret, (x[0] - y[0]).abs() * (x[1] - y[1]).abs())
                }
            }
        }
        if ret == std::i32::MAX {
            return 0;
        }
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::min_area_rect(vec![
            vec![1, 1],
            vec![1, 3],
            vec![3, 1],
            vec![3, 3],
            vec![2, 2],
        ])
    );
}
