// Accepted
// @lc id=283 lang=rust
// problem: move_zeroes
struct Solution {}
impl Solution {
    pub fn move_zeroes(nums: &mut Vec<i32>) {
        let mut lo = 0;
        for i in 0..nums.len() {
            if nums[i] != 0 {
                nums.swap(lo, i);
                lo += 1
            }
        }
    }
}

fn main() {
    println!("{:?}", Solution::move_zeroes(&mut vec![0, 1, 0, 3, 12]));
}
