// Accepted
// @lc id=213 lang=rust
// problem: house_robber_ii
struct Solution {}
impl Solution {
    fn do_rob(nums: &[i32]) -> i32 {
        let (mut rob, mut no_rob) = (0, 0);
        for x in nums {
            let (orig_rob, orig_no_rob) = (rob, no_rob);
            rob = x + orig_no_rob;
            no_rob = std::cmp::max(orig_rob, orig_no_rob)
        }
        std::cmp::max(rob, no_rob)
    }

    pub fn rob(nums: Vec<i32>) -> i32 {
        if nums.is_empty() {
            return 0;
        }
        if nums.len() == 1 {
            return nums[0];
        }
        std::cmp::max(
            Solution::do_rob(&nums[..nums.len() - 1]),
            Solution::do_rob(&nums[1..]),
        )
    }
}

fn main() {
    println!("{:?}", Solution::rob(vec![2, 3, 2]));
    println!("{:?}", Solution::rob(vec![1, 2]));
}
