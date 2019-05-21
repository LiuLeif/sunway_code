// Accepted
// @lc id=209 lang=rust
// problem: minimum_size_subarray_sum
struct Solution {}
impl Solution {
    pub fn min_sub_array_len(s: i32, nums: Vec<i32>) -> i32 {
        // 2 3 1 2 4 3
        if nums.is_empty() {
            return 0;
        }
        let (mut lo, mut hi, mut sum, mut ret) = (0, 0, 0, std::usize::MAX);
        while hi < nums.len() {
            sum += nums[hi];
            hi += 1;
            while sum >= s {
                ret = std::cmp::min(ret, hi - lo);
                sum -= nums[lo];
                lo += 1;
            }
        }
        if ret == std::usize::MAX {
            return 0;
        }
        ret as i32
    }

    pub fn min_sub_array_len_2(s: i32, nums: Vec<i32>) -> i32 {
        // 2 3 1 2 4 3
        // 2 5 6 8 12 15
        let mut sum_array = vec![0; nums.len() + 1];
        for i in 1..=nums.len() {
            sum_array[i] = sum_array[i - 1] + nums[i - 1];
        }
        let (mut lo, mut hi) = (1, nums.len());
        while lo <= hi {
            let mid = (lo + hi) / 2;
            let max_sum = {
                let mut ret = 0;
                for i in 0..=nums.len() - mid {
                    ret = std::cmp::max(ret, sum_array[i + mid] - sum_array[i]);
                }
                ret
            };
            if max_sum >= s {
                hi = mid - 1;
            } else {
                lo = mid + 1;
            }
        }
        if lo == nums.len() + 1 {
            return 0;
        }

        lo as i32
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::min_sub_array_len(7, vec![2, 3, 1, 2, 4, 3])
    );
    println!("{:?}", Solution::min_sub_array_len(4, vec![]));
}
