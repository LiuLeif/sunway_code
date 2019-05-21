// Accepted
// @lc id=653 lang=rust
// problem: two_sum_iv_input_is_a_bst
extern crate util;
use util::binary_tree::*;
use util::*;

use std::collections::VecDeque;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;

type Link = Rc<RefCell<TreeNode>>;

enum Order {
    LEFT,
    RIGHT,
}

struct BSTIterator {
    stack: VecDeque<Link>,
    order: Order,
}

impl BSTIterator {
    fn new(root: Link, order: Order) -> Self {
        let mut stack = VecDeque::new();
        stack.push_back(root.clone());
        BSTIterator { stack, order }
    }
}
impl Iterator for BSTIterator {
    type Item = Link;
    fn next(&mut self) -> Option<Link> {
        while !self.stack.is_empty() {
            let top = self.stack.pop_back().unwrap();
            let top_val = top.borrow().val;
            let left = top.borrow().left.clone();
            let right = top.borrow().right.clone();

            if left.is_none() && right.is_none() {
                return Some(top);
            }
            match self.order {
                Order::LEFT => {
                    if let Some(node) = right {
                        self.stack.push_back(node);
                    }
                    self.stack
                        .push_back(Rc::new(RefCell::new(TreeNode::new(top_val))));
                    if let Some(node) = left {
                        self.stack.push_back(node);
                    }
                }
                Order::RIGHT => {
                    if let Some(node) = left {
                        self.stack.push_back(node);
                    }
                    self.stack
                        .push_back(Rc::new(RefCell::new(TreeNode::new(top_val))));
                    if let Some(node) = right {
                        self.stack.push_back(node);
                    }
                }
            }
        }
        None
    }
}
impl Solution {
    pub fn find_target(root: Option<Rc<RefCell<TreeNode>>>, k: i32) -> bool {
        if root.is_none() {
            return false;
        }

        let root = root.unwrap();
        let (mut left, mut right) = (
            BSTIterator::new(root.clone(), Order::LEFT),
            BSTIterator::new(root.clone(), Order::RIGHT),
        );
        let (mut a, mut b) = (left.next(), right.next());
        while a.is_some() && b.is_some() {
            let (x, y) = (a.clone().unwrap(), b.clone().unwrap());
            if x.as_ptr() == y.as_ptr() {
                return false;
            }
            let (x_val, y_val) = (x.borrow().val, y.borrow().val);

            if x_val + y_val == k {
                return true;
            }
            if x_val + y_val > k {
                b = right.next();
            } else {
                a = left.next();
            }
        }
        false
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::find_target(tree!(5, 3, 6, 2, 4, None, 7), 18)
    );
}
