// Accepted
// @lc id=107 lang=rust
// problem: binary_tree_level_order_traversal_ii
use std::collections::VecDeque;
use util::binary_tree::*;
use util::*;
struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn level_order_bottom(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        if root.is_none() {
            return ret;
        }
        let root = root.unwrap();
        let mut queue = VecDeque::new();
        queue.push_back(root);
        while !queue.is_empty() {
            let mut curr = vec![];
            for _ in 0..queue.len() {
                let top = queue.pop_front().unwrap();
                let (left, right, val) = (
                    top.borrow().left.clone(),
                    top.borrow().right.clone(),
                    top.borrow().val,
                );
                curr.push(val);
                if let Some(node) = left {
                    queue.push_back(node);
                }
                if let Some(node) = right {
                    queue.push_back(node);
                }
            }
            ret.push(curr);
        }
        ret.reverse();
        return ret;
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::level_order_bottom(tree![3, 9, 20, None, None, 15, 7])
    );
}
