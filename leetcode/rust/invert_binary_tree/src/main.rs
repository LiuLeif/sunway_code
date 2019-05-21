// Accepted
// 2019-01-07 10:40
// @lc id=226 lang=rust
// problem: invert_binary_tree
extern crate util;
use util::binary_tree::*;

struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn invert_tree(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        if let Some(ref node) = root {
            let left = Solution::invert_tree(node.borrow().right.clone());
            let right = Solution::invert_tree(node.clone().borrow().left.clone());
            node.borrow_mut().left = left;
            node.borrow_mut().right = right;
        }
        root
    }
}

fn main() {
    let root = BinaryTree::new(vec![
        Some(4),
        Some(2),
        Some(7),
        Some(1),
        Some(3),
        Some(6),
        Some(9),
    ]);
    println!("{:?}", Solution::invert_tree(root));
}
