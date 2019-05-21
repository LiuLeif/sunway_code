// Accepted
// @lc id=701 lang=rust
// problem: insert_into_a_binary_search_tree

extern crate util;
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn insert_into_bst(
        root: Option<Rc<RefCell<TreeNode>>>,
        val: i32,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if root.is_none() {
            return Some(Rc::new(RefCell::new(TreeNode::new(val))));
        }
        let root = root.unwrap();
        let left = root.borrow().left.clone();
        let right = root.borrow().right.clone();
        if val > root.borrow().val {
            root.borrow_mut().right = Solution::insert_into_bst(right, val);
        } else {
            root.borrow_mut().left = Solution::insert_into_bst(left, val);
        }
        return Some(root);
    }
}

fn main() {
    let root = tree!(4, 2, 7, 1, 3);
    println!("{:?}", Solution::insert_into_bst(root, 5));
}
