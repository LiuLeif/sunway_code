// Accepted
// @lc id=98 lang=rust
// problem: validate_binary_search_tree
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn dfs(root: Link) -> (bool, i32, i32) {
        let root = root.unwrap();
        let (left, right, val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            root.borrow().val,
        );
        let mut min_val = val;
        let mut max_val = val;
        if left.is_some() {
            let (is_valid, min, max) = Solution::dfs(left);
            if !is_valid || val <= max {
                return (false, 0, 0);
            }
            min_val = min;
        }
        if right.is_some() {
            let (is_valid, min, max) = Solution::dfs(right);
            if !is_valid || val >= min {
                return (false, 0, 0);
            }
            max_val = max;
        }
        (true, min_val, max_val)
    }

    pub fn is_valid_bst(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if root.is_none() {
            return true;
        }
        let (is_valid, _, _) = Solution::dfs(root);
        is_valid
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::is_valid_bst(tree![5, 1, 4, None, None, 3, 6])
    );
    println!("{:?}", Solution::is_valid_bst(tree![1, 1]));
}
