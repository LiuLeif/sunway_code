// Accepted
// @lc id=199 lang=rust
// problem: binary_tree_right_side_view
// tags: bfs
use std::collections::VecDeque;
use util::binary_tree::*;
use util::*;
struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn right_side_view(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut ret = vec![];
        if root.is_none() {
            return ret;
        }
        let mut queue = VecDeque::new();
        queue.push_back(root.unwrap());
        while !queue.is_empty() {
            for i in 0..queue.len() {
                let top = queue.pop_front().unwrap();
                let (left, right, val) = (
                    top.borrow().left.clone(),
                    top.borrow().right.clone(),
                    top.borrow().val,
                );
                if i == 0 {
                    ret.push(val);
                }
                if right.is_some() {
                    queue.push_back(right.unwrap());
                }
                if left.is_some() {
                    queue.push_back(left.unwrap());
                }
            }
        }
        ret
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::right_side_view(tree![1, 2, 3, None, 5, None, 4])
    );
}
