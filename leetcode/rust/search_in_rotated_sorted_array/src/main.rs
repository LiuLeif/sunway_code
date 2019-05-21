// Accepted
// @lc id=33 lang=rust
// problem: search_in_rotated_sorted_array
struct Solution {}
impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        if nums.is_empty() {
            return -1;
        }
        let (mut lo, mut hi) = (0, nums.len() - 1);
        while lo <= hi {
            let mid = (lo + hi) / 2;
            if nums[mid] == target {
                return mid as i32;
            }
            if nums[lo] <= nums[mid] {
                if target >= nums[lo] && target < nums[mid] {
                    hi = mid - 1;
                } else {
                    lo = mid + 1;
                }
            } else {
                if target > nums[mid] && target <= nums[hi] {
                    lo = mid + 1;
                } else {
                    hi = mid - 1;
                }
            }
        }
        return -1;
    }
}

fn main() {
    println!("{:?}", Solution::search(vec![4, 5, 6, 7, 0, 1, 2], 4));
    println!("{:?}", Solution::search(vec![5, 1, 3], 1));
}
