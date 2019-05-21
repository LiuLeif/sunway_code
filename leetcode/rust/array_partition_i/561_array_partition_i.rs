// 2018-12-26 15:21
// @lc id=561 lang=rust
struct Solution {}
impl Solution {
    pub fn array_pair_sum(nums: Vec<i32>) -> i32 {
        // 1 2 3 4
        let mut nums = nums;
        nums.sort();
        return nums.iter().step_by(2).sum();
    }
}

fn main() {
    println!("{:?}", Solution::array_pair_sum(vec![1, 4, 3, 2]));
}
