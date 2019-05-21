// Accepted
// @lc id=300 lang=rust
// problem: longest_increasing_subsequence
use std::collections::BTreeMap;
struct Solution {}
impl Solution {
    pub fn length_of_lis(nums: Vec<i32>) -> i32 {
        let n = nums.len();
        let mut dp = BTreeMap::new();
        let mut ret = 0;
        for i in 0..n {
            let max = dp.range(..nums[i]).map(|(_, &v)| v).max().unwrap_or(0) + 1;
            dp.insert(&nums[i], max);
            ret = std::cmp::max(ret, max);
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::length_of_lis(vec![1, 2, 3, 4, 5]));
}
