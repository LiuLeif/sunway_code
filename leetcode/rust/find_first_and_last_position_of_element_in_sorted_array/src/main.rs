// Accepted
// @lc id=34 lang=rust
// problem: find_first_and_last_position_of_element_in_sorted_array
struct Solution {}
impl Solution {
    fn first_smaller(nums: &[i32], target: i32) -> i32 {
        let (mut lo, mut hi) = (0, nums.len() as i32 - 1);
        while lo <= hi {
            let mid = (lo + hi) / 2;
            if nums[mid as usize] >= target {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        hi as i32
    }
    fn first_larger(nums: &[i32], target: i32) -> i32 {
        let (mut lo, mut hi) = (0, nums.len() as i32 - 1);
        while lo <= hi {
            let mid = (lo + hi) / 2;
            if nums[mid as usize] > target {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        lo as i32
    }
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        if nums.is_empty() {
            return vec![-1, -1];
        }
        let smaller = Solution::first_smaller(&nums, target);
        let larger = Solution::first_larger(&nums, target);
        if smaller != nums.len() as i32 - 1 && nums[(smaller + 1) as usize] == target {
            vec![smaller + 1, larger - 1]
        } else {
            vec![-1, -1]
        }
    }
}

fn main() {
    // println!("{:?}", Solution::search_range(vec![5, 7, 7, 8, 8, 10], 12));
    println!("{:?}", Solution::search_range(vec![1], 1));
}
