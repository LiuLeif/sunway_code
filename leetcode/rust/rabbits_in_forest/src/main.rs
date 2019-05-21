// Accepted
// @lc id=781 lang=rust
// problem: rabbits_in_forest
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn num_rabbits(answers: Vec<i32>) -> i32 {
        // 1 1 2
        let mut bucket = HashMap::<i32, i32>::new();
        let mut ret = 0;
        for i in answers {
            let curr = bucket.entry(i).or_insert(0);
            if *curr == 0 {
                *curr = i;
                ret += i + 1;
            } else {
                *curr -= 1;
            }
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::num_rabbits(vec![10, 10, 10]));
}
