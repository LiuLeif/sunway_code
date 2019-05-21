// Accepted
// @lc id=814 lang=rust
// problem: binary_tree_pruning
extern crate util;
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;

impl Solution {
    pub fn prune_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        if root.is_none() {
            return None;
        }
        let root = root.unwrap();

        let left = Solution::prune_tree(root.borrow().left.clone());
        let right = Solution::prune_tree(root.borrow().right.clone());

        if left.is_none() && right.is_none() && root.borrow().val == 0 {
            return None;
        }
        root.borrow_mut().left = left;
        root.borrow_mut().right = right;
        return Some(root);
    }
}

fn main() {
    let root = tree!(1, 0, 0, 0);
    println!("{:?}", Solution::prune_tree(root));
}
