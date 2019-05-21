// Accepted
// @lc id=783 lang=rust
// problem: minimum_distance_between_bst_nodes
extern crate util;
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Rc<RefCell<TreeNode>>;

impl Solution {
    fn inorder(root: Option<Link>, prev: &mut i32, min: &mut i32) {
        if let Some(node) = root {
            let node_value = node.borrow().val;
            let left = node.borrow().left.clone();
            let right = node.borrow().right.clone();
            Solution::inorder(left, prev, min);
            if *prev != std::i32::MIN {
                *min = std::cmp::min(*min, node_value - *prev);
            }
            *prev = node_value;
            Solution::inorder(right, prev, min);
        }
    }

    pub fn min_diff_in_bst(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut prev = std::i32::MIN;
        let mut ret = std::i32::MAX;
        Solution::inorder(root, &mut prev, &mut ret);
        ret
    }
}

fn main() {
    println!("{:?}", Solution::min_diff_in_bst(tree!(4, 2, 6)));
}
