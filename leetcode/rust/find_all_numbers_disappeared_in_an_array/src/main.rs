// Accepted
// @lc id=448 lang=rust
// problem: find_all_numbers_disappeared_in_an_array
struct Solution {}
impl Solution {
    pub fn find_disappeared_numbers(nums: Vec<i32>) -> Vec<i32> {
        let mut nums = nums;
        let n = nums.len();
        for i in 0..n {
            let x = nums[i].abs() as usize;
            if nums[x - 1] > 0 {
                nums[x - 1] = -nums[x - 1];
            }
        }
        nums.into_iter()
            .enumerate()
            .filter(|&(i, v)| v > 0)
            .map(|(i, _)| i as i32 + 1)
            .collect()
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_disappeared_numbers(vec![4, 3, 2, 7, 8, 2, 3, 1])
    );
}
