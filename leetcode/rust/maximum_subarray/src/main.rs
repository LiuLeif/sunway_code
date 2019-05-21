// Accepted
// @lc id=53 lang=rust
// problem: maximum_subarray
struct Solution {}
impl Solution {
    pub fn max_sub_array(nums: Vec<i32>) -> i32 {
        if nums.is_empty() {
            return 0;
        }
        let mut dp = vec![0; nums.len()];
        dp[0] = nums[0];
        for i in 1..nums.len() {
            dp[i] = std::cmp::max(dp[i - 1] + nums[i], nums[i]);
        }
        *dp.iter().max().unwrap()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::max_sub_array(vec![-2, 1, -3, 4, -1, 2, 1, -5, 4])
    );
}
