// Accepted
// @lc id=46 lang=rust
// problem: permutations
struct Solution {}
impl Solution {
    fn backtrack(ret: &mut Vec<Vec<i32>>, curr: &mut Vec<i32>, nums: &[i32], mask: &mut Vec<bool>) {
        if curr.len() == nums.len() {
            ret.push(curr.clone());
            return;
        }
        for i in 0..nums.len() {
            if mask[i] {
                continue;
            }
            mask[i] = true;
            curr.push(nums[i]);
            Solution::backtrack(ret, curr, nums, mask);
            curr.pop();
            mask[i] = false;
        }
    }

    pub fn permute(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        Solution::backtrack(&mut ret, &mut vec![], &nums, &mut vec![false; nums.len()]);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::permute(vec![1, 2, 3]));
}
