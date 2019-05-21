// Accepted
// @lc id=454 lang=rust
// problem: 4sum_ii
use std::collections::HashMap;

struct Solution {}
impl Solution {
    pub fn four_sum_count(a: Vec<i32>, b: Vec<i32>, c: Vec<i32>, d: Vec<i32>) -> i32 {
        let mut counter: HashMap<i32, i32> = HashMap::new();
        for x in a.iter() {
            for y in b.iter() {
                *counter.entry(x + y).or_default() += 1;
            }
        }

        let mut ret = 0;
        for x in c.iter() {
            for y in d.iter() {
                ret += counter.get(&(-x - y)).unwrap_or(&0);
            }
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::four_sum_count(vec![1, 2], vec![-2, -1], vec![-1, 2], vec![0, 2],)
    );
}
