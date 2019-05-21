// Accepted
// 2019-01-05 21:26
// @lc id=888 lang=rust
// problem: fair_candy_swap
use std::collections::HashSet;
struct Solution {}
impl Solution {
    pub fn fair_candy_swap(a: Vec<i32>, b: Vec<i32>) -> Vec<i32> {
        let set_b = b.iter().collect::<HashSet<_>>();
        let (sum_a, sum_b): (i32, i32) = (a.iter().sum(), b.iter().sum());
        if sum_a == sum_b {
            return vec![0, 0];
        }
        let delta = (sum_b - sum_a) / 2;
        for i in a {
            if set_b.contains(&(i + delta)) {
                return vec![i, i + delta];
            }
        }
        return vec![];
    }
}

fn main() {
    println!("{:?}", Solution::fair_candy_swap(vec![1, 2], vec![2, 3]));
}
