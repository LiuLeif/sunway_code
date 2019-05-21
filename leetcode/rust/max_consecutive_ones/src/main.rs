// Accepted
// 2019-01-07 12:11
// @lc id=485 lang=rust
// problem: max_consecutive_ones
struct Solution {}
impl Solution {
    pub fn find_max_consecutive_ones(nums: Vec<i32>) -> i32 {
        nums.split(|&n| n == 0).map(|g| g.len()).max().unwrap() as i32
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_max_consecutive_ones(vec![1, 1, 0, 1, 1, 1])
    );
}
