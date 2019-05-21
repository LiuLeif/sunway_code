// Accepted
// @lc id=15 lang=rust
// problem: 3sum
use std::collections::HashMap;
struct Solution {}
impl Solution {
    pub fn three_sum(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        let mut nums = nums;
        nums.sort();
        let mut counter = HashMap::new();
        for (i, v) in nums.iter().enumerate().rev() {
            counter.entry(-v).or_insert(i);
        }
        for i in 0..nums.len() {
            if i != 0 && nums[i] == nums[i - 1] {
                continue;
            }
            for j in i + 1..nums.len() {
                if j != i + 1 && nums[j] == nums[j - 1] {
                    continue;
                }
                let index = *counter.entry(nums[j] + nums[i]).or_insert(0);
                if index > j {
                    ret.push(vec![nums[i], nums[j], -nums[i] - nums[j]]);
                }
            }
        }
        ret
    }
}

fn main() {
    println!("{:?}", Solution::three_sum(vec![-1, 0, 1, 2, -1, -4]));
    println!("{:?}", Solution::three_sum(vec![0, 0, 0, 1, -1]));
}
