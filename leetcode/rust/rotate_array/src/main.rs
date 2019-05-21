// Accepted
// @lc id=189 lang=rust
// problem: rotate_array
struct Solution {}
impl Solution {
    pub fn rotate(nums: &mut Vec<i32>, k: i32) {
        let k = k % nums.len() as i32;
        nums.reverse();
        &nums[..k as usize].reverse();
        &nums[k as usize..].reverse();
    }
}

fn main() {
    let mut v = vec![1, 2, 3, 4, 5, 6, 7];
    Solution::rotate(&mut v, 3);
    println!("{:?}", v);
}
