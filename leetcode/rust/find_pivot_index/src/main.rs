// Accepted
// @lc id=724 lang=rust
// problem: find_pivot_index
struct Solution {}
impl Solution {
    pub fn pivot_index(nums: Vec<i32>) -> i32 {
        if nums.is_empty() {
            return -1;
        }
        let sum = nums.iter().sum::<i32>();
        let mut total = 0;
        for i in 0..nums.len() {
            if i != 0 {
                total += nums[i - 1];
            }

            if total * 2 + nums[i] == sum {
                return i as i32;
            }
        }
        return -1;
    }
}

fn main() {
    println!("{:?}", Solution::pivot_index(vec![1, 7, 3, 6, 5, 6]));
}
