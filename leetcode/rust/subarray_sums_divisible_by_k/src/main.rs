// Accepted
// @lc id=974 lang=rust
// problem: subarray_sums_divisible_by_k
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn subarrays_div_by_k(a: Vec<i32>, k: i32) -> i32 {
        let mut sum = 0;
        let mut counter = HashMap::new();
        counter.insert(0, 1);
        let mut ret = 0;
        for x in a {
            sum += x;
            sum = (sum % k + k) % k;
            let prev = counter.entry(sum).or_insert(0);
            ret += *prev;
            *prev += 1;
        }
        println!("{:?}", counter);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::subarrays_div_by_k(vec![-1, -9, -4, 0], 9));
}
