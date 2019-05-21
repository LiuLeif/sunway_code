// Accepted
// @lc id=198 lang=rust
// problem: house_robber
struct Solution {}
impl Solution {
    pub fn rob(nums: Vec<i32>) -> i32 {
        let (mut rob, mut no_rob) = (0, 0);
        for x in nums {
            let (orig_rob, orig_no_rob) = (rob, no_rob);
            rob = x + orig_no_rob;
            no_rob = std::cmp::max(orig_rob, orig_no_rob)
        }
        std::cmp::max(rob, no_rob)
    }
}

fn main() {
    println!("{:?}", Solution::rob(vec![2, 7, 9, 3, 1]));
}
