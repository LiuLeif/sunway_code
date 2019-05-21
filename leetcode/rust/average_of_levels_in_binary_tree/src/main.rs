// Accepted
// 2019-01-04 10:22
// @lc id=637 lang=rust
// problem: average_of_levels_in_binary_tree
extern crate util;
use util::binary_tree::*;

struct Solution {}
use std::cell::RefCell;
use std::collections::VecDeque;
use std::rc::Rc;
impl Solution {
    pub fn average_of_levels(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<f64> {
        let mut ret = vec![];
        let mut queue = VecDeque::new();
        queue.push_back(root);
        queue.push_back(None);
        let mut total = 0 as i64;
        let mut count = 0 as i64;
        while queue.len() != 1 {
            let top = queue.pop_front().unwrap();
            if top == None {
                queue.push_back(None);
                ret.push(total as f64 / count as f64);
                total = 0;
                count = 0;
                continue;
            }
            let top = top.unwrap();
            total += top.borrow().val as i64;
            count += 1;
            let left = top.borrow().left.clone();
            if left != None {
                queue.push_back(left);
            }
            let right = top.borrow().right.clone();
            if right != None {
                queue.push_back(right);
            }
        }
        ret.push(total as f64 / count as f64);
        ret
    }
}

fn main() {
    let root = BinaryTree::new(vec![Some(2147483647), Some(2147483647), Some(2147483647)]);
    println!("{:?}", Solution::average_of_levels(root));
}
