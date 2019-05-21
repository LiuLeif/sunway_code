// Accepted
// @lc id=217 lang=rust
// problem: contains_duplicate
struct Solution {}
impl Solution {
    pub fn contains_duplicate(nums: Vec<i32>) -> bool {
        let mut nums = nums;
        nums.sort();
        nums.windows(2).any(|w| w[0] == w[1])
    }
}

fn main() {
    println!("{:?}", Solution::contains_duplicate(vec![1, 2, 3, 1]));
}
