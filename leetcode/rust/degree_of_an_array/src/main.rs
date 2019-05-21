// Accepted
// @lc id=697 lang=rust
// problem: degree_of_an_array
use std::collections::HashMap;
#[derive(Debug, Clone)]
struct Tuple {
    left: usize,
    right: usize,
    count: i32,
}
struct Solution {}
impl Solution {
    pub fn find_shortest_sub_array(nums: Vec<i32>) -> i32 {
        let mut counter = HashMap::new();
        for (i, n) in nums.iter().enumerate() {
            let entry = counter.entry(n).or_insert(Tuple {
                left: i,
                right: i,
                count: 0,
            });
            entry.right = std::cmp::max(entry.right, i);
            entry.count += 1;
        }
        let mut v = counter.values().collect::<Vec<&Tuple>>();
        v.sort_by(|t1, t2| {
            if t1.count == t2.count {
                return (t1.right - t1.left).cmp(&(t2.right - t2.left));
            }
            return t2.count.cmp(&t1.count);
        });
        v.first().map_or(0, |t| (t.right - t.left) as i32 + 1)
    }
}

fn main() {
    println!("{:?}", Solution::find_shortest_sub_array(vec![1, 2, 2, 1]));
}
