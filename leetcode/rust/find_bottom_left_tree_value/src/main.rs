// Accepted
// @lc id=513 lang=rust
// problem: find_bottom_left_tree_value
extern crate util;
use std::collections::VecDeque;
use util::binary_tree::*;

struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn find_bottom_left_value(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        let mut queue = VecDeque::new();
        queue.push_back(root);
        let mut ret = 0;
        while queue.len() != 0 {
            let top = queue.pop_front().unwrap().unwrap();
            ret = top.borrow().val;
            let left = top.borrow().left.clone();
            let right = top.borrow().right.clone();
            if right != None {
                queue.push_back(right)
            }
            if left != None {
                queue.push_back(left)
            }
        }
        ret
    }
}

fn main() {
    let root = BinaryTree::new(vec![
        Some(1),
        Some(2),
        Some(3),
        Some(4),
        None,
        Some(5),
        Some(6),
        None,
        None,
        Some(7),
    ]);
    println!("{:?}", Solution::find_bottom_left_value(root));
}
