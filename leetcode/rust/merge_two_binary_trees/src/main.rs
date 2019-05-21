// Accepted
// @lc id=617 lang=rust
// problem: merge_two_binary_trees
extern crate util;
use util::binary_tree::*;
struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;

impl Solution {
    pub fn merge_trees(t1: Link, t2: Link) -> Link {
        if t1.is_none() && t2.is_none() {
            return None;
        }
        let mut t1_val = 0;
        let mut t2_val = 0;
        let mut t1_left = None;
        let mut t1_right = None;
        let mut t2_left = None;
        let mut t2_right = None;

        if let Some(node) = t1.clone() {
            t1_val = node.borrow().val;
            t1_left = node.borrow().left.clone();
            t1_right = node.borrow().right.clone();
        }
        if let Some(node) = t2.clone() {
            t2_val = node.borrow().val;
            t2_left = node.borrow().left.clone();
            t2_right = node.borrow().right.clone();
        }
        let mut root = TreeNode::new(t1_val + t2_val);
        root.left = Solution::merge_trees(t1_left, t2_left);
        root.right = Solution::merge_trees(t1_right, t2_right);
        return Some(Rc::new(RefCell::new(root)));
    }
}

fn main() {
    let root1 = BinaryTree::new(vec![Some(1), Some(2), Some(3)]);
    let root2 = BinaryTree::new(vec![Some(4), Some(5), Some(6)]);
    println!("{:?}", Solution::merge_trees(root1, root2));
}
