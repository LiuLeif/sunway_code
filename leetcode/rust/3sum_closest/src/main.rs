// Accepted
// @lc id=16 lang=rust
// problem: 3sum_closest
struct Solution {}
impl Solution {
    pub fn three_sum_closest(nums: Vec<i32>, target: i32) -> i32 {
        // -4 -1 1 2  -> 1
        let mut nums = nums;
        nums.sort();
        let mut min_delta = std::i32::MAX;
        let mut ret = 0;
        for i in 0..nums.len() {
            let (mut left, mut right) = (0, nums.len() - 1);
            while left < right {
                if left == i {
                    left += 1;
                    continue;
                }
                if right == i {
                    right -= 1;
                    continue;
                }
                let sum = nums[left] + nums[right] + nums[i];
                if sum == target {
                    return target;
                }
                let mut delta = (sum - target).abs();
                if delta < min_delta {
                    min_delta = delta;
                    ret = sum;
                }
                if sum > target {
                    right -= 1;
                } else {
                    left += 1;
                }
            }
        }
        return ret;
    }
}

fn main() {
    println!("{:?}", Solution::three_sum_closest(vec![-1, 2, 1, -4], 1));
}
