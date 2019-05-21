// Accepted
// @lc id=447 lang=rust
// problem: number_of_boomerangs
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn number_of_boomerangs(points: Vec<Vec<i32>>) -> i32 {
        let n = points.len();
        let mut distance = vec![HashMap::new(); n];
        for i in 0..n {
            for j in i + 1..n {
                let (p1, p2) = (&points[i], &points[j]);
                let d = (p1[0] - p2[0]).pow(2) + (p1[1] - p2[1]).pow(2);
                *distance[i].entry(d).or_insert(0) += 1;
                *distance[j].entry(d).or_insert(0) += 1;
            }
        }
        distance
            .iter()
            .flat_map(|m| m.values().map(|k| k * (k - 1)))
            .sum()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::number_of_boomerangs(vec![
            vec![0, 0],
            vec![1, 0],
            vec![-1, 0],
            vec![0, 1],
            vec![0, -1],
        ])
    );
}
