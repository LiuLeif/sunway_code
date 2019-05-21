// Accepted
// @lc id=47 lang=rust
// problem: permutations_ii
use std::collections::HashSet;
struct Solution {}
impl Solution {
    fn backtrack(
        ret: &mut Vec<Vec<i32>>,
        curr: &mut Vec<i32>,
        visited_index: &mut HashSet<usize>,
        nums: &Vec<i32>,
    ) {
        if visited_index.len() == nums.len() {
            ret.push(curr.clone());
            return;
        }
        let mut visited_val = HashSet::new();
        for i in 0..nums.len() {
            if visited_val.contains(&nums[i]) {
                continue;
            }
            if visited_index.contains(&i) {
                continue;
            }
            visited_val.insert(nums[i]);
            visited_index.insert(i);
            curr.push(nums[i]);
            Solution::backtrack(ret, curr, visited_index, nums);
            curr.pop();
            visited_index.remove(&i);
        }
    }

    pub fn permute_unique(nums: Vec<i32>) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        Solution::backtrack(&mut ret, &mut vec![], &mut HashSet::new(), &nums);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::permute_unique(vec![1, 1, 2]));
}
