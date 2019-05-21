// Accepted
// @lc id=108 lang=rust
// problem: convert_sorted_array_to_binary_search_tree
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;

type Link = Option<Rc<RefCell<TreeNode>>>;

impl Solution {
    fn dfs(nums: &[i32]) -> Link {
        if nums.is_empty() {
            return None;
        }
        let mid = (nums.len() - 1) / 2;
        let mut ret = TreeNode::new(nums[mid]);
        ret.left = Solution::dfs(&nums[..mid]);
        ret.right = Solution::dfs(&nums[mid + 1..]);
        return Some(Rc::new(RefCell::new(ret)));
    }

    pub fn sorted_array_to_bst(nums: Vec<i32>) -> Option<Rc<RefCell<TreeNode>>> {
        Solution::dfs(&nums)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::sorted_array_to_bst(vec![-10, -3, 0, 5, 9])
    );
}
