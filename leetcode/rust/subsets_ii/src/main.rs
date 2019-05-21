// Accepted
// @lc id=90 lang=rust
// problem: subsets_ii
struct Solution {}
impl Solution {
    fn backtrack(ret: &mut Vec<Vec<i32>>, curr: &mut Vec<i32>, index: usize, nums: &[i32]) -> () {
        ret.push(curr.clone());

        if index == nums.len() {
            return;
        }
        for i in index..nums.len() {
            if i != index && nums[i] == nums[i - 1] {
                continue;
            }
            curr.push(nums[i]);
            Solution::backtrack(ret, curr, i + 1, nums);
            curr.pop();
        }
    }

    pub fn subsets_with_dup(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        let mut nums = nums;
        nums.sort();
        Solution::backtrack(&mut ret, &mut vec![], 0, &nums);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::subsets_with_dup(vec![1, 2, 2]));
}
