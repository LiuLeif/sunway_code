// Accepted
// 2019-01-07 17:42
// @lc id=700 lang=rust
// problem: search_in_a_binary_search_tree
extern crate util;
use util::binary_tree::*;

struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn search_bst(
        root: Option<Rc<RefCell<TreeNode>>>,
        val: i32,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        let mut curr = root;
        while let Some(node) = curr.clone() {
            if node.borrow().val == val {
                return curr;
            }
            if node.borrow().val > val {
                curr = node.borrow().left.clone();
            } else {
                curr = node.borrow().right.clone();
            }
        }
        return curr;
    }
}

fn main() {
    let root = BinaryTree::new(vec![Some(4), Some(2), Some(7), Some(1), Some(3)]);
    println!("{:?}", Solution::search_bst(root, 4));
}
