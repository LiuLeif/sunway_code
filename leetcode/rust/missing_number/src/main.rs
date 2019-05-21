// Accepted
// @lc id=268 lang=rust
// problem: missing_number
struct Solution {}
impl Solution {
    pub fn missing_number(nums: Vec<i32>) -> i32 {
        let n = nums.len() as i32;
        n * (n + 1) / 2 - nums.into_iter().sum::<i32>()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::missing_number(vec![9, 6, 4, 2, 3, 5, 7, 0, 1])
    );
}
