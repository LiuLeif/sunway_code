// Accepted
// @lc id=102 lang=rust
// problem: binary_tree_level_order_traversal
use std::collections::VecDeque;
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut ret = vec![];
        let mut queue = VecDeque::new();
        if root.is_none() {
            return ret;
        }
        queue.push_back(root.unwrap());
        while !queue.is_empty() {
            let mut row = vec![];
            for _ in 0..queue.len() {
                let top = queue.pop_front().unwrap();
                let (left, right, val) = (
                    top.borrow().left.clone(),
                    top.borrow().right.clone(),
                    top.borrow().val,
                );
                row.push(val);
                if left.is_some() {
                    queue.push_back(left.unwrap());
                }
                if right.is_some() {
                    queue.push_back(right.unwrap());
                }
            }
            ret.push(row);
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::level_order(tree![3, 9, 20, None, None, 15, 7])
    );
}
