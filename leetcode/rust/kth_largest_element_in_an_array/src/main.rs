// Accepted
// @lc id=215 lang=rust
// problem: kth_largest_element_in_an_array
struct Solution {}
impl Solution {
    fn partition(nums: &mut [i32]) -> usize {
        let (mut p, mut lo, hi) = (0, 1, nums.len() - 1);
        while lo <= hi {
            if nums[p] >= nums[lo] {
                lo += 1;
            } else {
                nums.swap(lo, p + 1);
                nums.swap(p, p + 1);
                lo += 1;
                p += 1;
            }
        }
        p
    }

    fn quick_select(nums: &mut [i32], k: i32) -> i32 {
        let p = Solution::partition(nums);
        if p as i32 == k - 1 {
            return nums[p as usize];
        }
        if p as i32 > k - 1 {
            Solution::quick_select(&mut nums[..p], k)
        } else {
            Solution::quick_select(&mut nums[p + 1..], k - p as i32 - 1)
        }
    }

    pub fn find_kth_largest(nums: Vec<i32>, k: i32) -> i32 {
        let mut nums = nums;
        Solution::quick_select(&mut nums, k)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_kth_largest(vec![3, 2, 1, 5, 6, 4], 3)
    );
}
