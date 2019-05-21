// Accepted
// @lc id=334 lang=rust
// problem: increasing_triplet_subsequence
struct Solution {}
impl Solution {
    pub fn increasing_triplet(nums: Vec<i32>) -> bool {
        let n = nums.len();
        if n < 3 {
            return false;
        }
        let (mut smaller, mut larger) = (vec![0; n], vec![0; n]);
        let mut min = std::i32::MAX;
        for i in 0..n {
            smaller[i] = min;
            min = std::cmp::min(min, nums[i]);
        }
        let mut max = std::i32::MIN;
        for i in (0..n).rev() {
            larger[i] = max;
            max = std::cmp::max(max, nums[i]);
        }
        for i in 1..n - 1 {
            if nums[i] > smaller[i] && nums[i] < larger[i] {
                return true;
            }
        }
        return false;
    }
}

fn main() {
    println!("{:?}", Solution::increasing_triplet(vec![1, 2, 3, 4, 5]));
    println!("{:?}", Solution::increasing_triplet(vec![3, 2, 1, 4, 9]));
}
