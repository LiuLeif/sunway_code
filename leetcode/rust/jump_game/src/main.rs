// Accepted
// @lc id=55 lang=rust
// problem: jump_game
struct Solution {}
impl Solution {
    pub fn can_jump(nums: Vec<i32>) -> bool {
        let mut hi = 0;
        for lo in 0..nums.len() {
            if lo > hi {
                return false;
            }
            hi = std::cmp::max(hi, nums[lo] as usize + lo);
        }
        true
    }
}

fn main() {
    println!("{:?}", Solution::can_jump(vec![2, 3, 1, 1, 4]));
    println!("{:?}", Solution::can_jump(vec![0, 1]));
}
