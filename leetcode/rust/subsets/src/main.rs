// Accepted
// @lc id=78 lang=rust
// problem: subsets
struct Solution {}
impl Solution {
    fn backtrack(ret: &mut Vec<Vec<i32>>, curr: &mut Vec<i32>, nums: &[i32], start: usize) {
        ret.push(curr.clone());
        if start == nums.len() {
            return;
        }
        for i in start..nums.len() {
            curr.push(nums[i]);
            Solution::backtrack(ret, curr, nums, i + 1);
            curr.pop();
        }
    }

    pub fn subsets(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        Solution::backtrack(&mut ret, &mut vec![], &nums, 0);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::subsets(vec![1, 2, 3]));
}
