// Accepted
// @lc id=162 lang=rust
// problem: find_peak_element
struct Solution {}
impl Solution {
    pub fn find_peak_element(nums: Vec<i32>) -> i32 {
        let (mut lo, mut hi) = (0, nums.len() - 1);
        while lo <= hi {
            let mid = (lo + hi) / 2;
            if (mid == 0 || (nums[mid] > nums[mid - 1]))
                && (mid == nums.len() - 1 || (nums[mid] > nums[mid + 1]))
            {
                return mid as i32;
            }
            if mid != 0 && nums[mid - 1] > nums[mid] {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        return lo as i32;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_peak_element(vec![1, 2, 1, 3, 5, 6, 4])
    );
    println!("{:?}", Solution::find_peak_element(vec![1, 0]));
}
