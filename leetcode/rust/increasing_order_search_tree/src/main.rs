// Accepted
// @lc id=897 lang=rust
// problem: increasing_order_search_tree
extern crate util;
use util::binary_tree::*;
use util::*;

struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn increasing_subtree(root: Link) -> (Link, Link) {
        let root = root.unwrap();
        let left = root.borrow().left.clone();
        let right = root.borrow().right.clone();
        let (mut head, mut tail) = (Some(root.clone()), Some(root.clone()));
        if left.is_some() {
            root.borrow_mut().left = None;
            let (left_head, left_tail) = Solution::increasing_subtree(left);
            head = left_head;
            left_tail.unwrap().borrow_mut().right = Some(root.clone());
        }
        if right.is_some() {
            let (right_head, right_tail) = Solution::increasing_subtree(right);
            root.borrow_mut().right = right_head;
            tail = right_tail;
        }
        (head, tail)
    }

    pub fn increasing_bst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        let (head, _) = Solution::increasing_subtree(root);
        head
    }
}

fn main() {
    let root = tree!(5, 3, 6, 2, 4, None, 8);
    println!("{:?}", Solution::increasing_bst(root));
}
