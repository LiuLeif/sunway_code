// Accepted
// @lc id=404 lang=rust
// problem: sum_of_left_leaves
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    fn dfs(root: Option<Rc<RefCell<TreeNode>>>, check: bool) -> i32 {
        if root.is_none() {
            return 0;
        }
        let root = root.unwrap();
        let (left, right, val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            root.borrow().val,
        );
        let mut ret = 0;
        if check && left.is_none() && right.is_none() {
            ret += val;
        }
        if left.is_some() {
            ret += Solution::dfs(left, true);
        }
        if right.is_some() {
            ret += Solution::dfs(right, false);
        }
        ret
    }

    pub fn sum_of_left_leaves(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        Solution::dfs(root, false)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::sum_of_left_leaves(tree![3, 9, 20, None, None, 15, 7])
    );
}
