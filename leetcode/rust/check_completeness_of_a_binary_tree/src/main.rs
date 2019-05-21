// Accepted
// @lc id=958 lang=rust
// problem: check_completeness_of_a_binary_tree
use util::binary_tree::*;
use util::*;

use std::collections::VecDeque;
struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn is_complete_tree(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        let mut q = VecDeque::new();
        q.push_back(root.clone());
        let mut last = false;
        while !q.is_empty() {
            for _ in 0..q.len() {
                match q.pop_front().unwrap() {
                    Some(top) => {
                        if last {
                            return false;
                        }
                        let (left, right) = (top.borrow().left.clone(), top.borrow().right.clone());
                        q.push_back(left);
                        q.push_back(right);
                    }
                    None => last = true,
                }
            }
        }
        true
    }
}

fn main() {
    println!("{:?}", Solution::is_complete_tree(tree![1, 2, 3, 4, 5, 6]));
    println!("{:?}", Solution::is_complete_tree(tree![1, None, 2]));
}
