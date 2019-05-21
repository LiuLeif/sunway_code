// Accepted
// 2019-01-03 17:49
// @lc id=540 lang=rust
// problem: single_element_in_a_sorted_array
struct Solution {}
impl Solution {
    pub fn single_non_duplicate(nums: Vec<i32>) -> i32 {
        fn pair(x: usize) -> usize {
            if x % 2 == 0 {
                x + 1
            } else {
                x - 1
            }
        }

        let (mut lo, mut hi) = (0, nums.len() - 1);
        while lo < hi {
            let mid = (lo + hi) / 2;
            if nums[mid] == nums[pair(mid)] {
                lo = mid + 1;
            } else {
                hi = mid - 1;
            }
        }
        return nums[lo];
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::single_non_duplicate(vec![1, 1, 2, 3, 3, 4, 4, 8, 8])
    );
}
