// Accepted
// @lc id=222 lang=rust
// problem: count_complete_tree_nodes
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;

impl Solution {
    fn left_height(root: Link) -> i32 {
        match root {
            Some(node) => Solution::left_height(node.borrow().left.clone()) + 1,
            None => 0,
        }
    }

    fn right_height(root: Link) -> i32 {
        match root {
            Some(node) => Solution::right_height(node.borrow().right.clone()) + 1,
            None => 0,
        }
    }

    pub fn count_nodes(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if root.is_none() {
            return 0;
        }
        let (left_height, right_height) = (
            Solution::left_height(root.clone()),
            Solution::right_height(root.clone()),
        );
        if left_height == right_height {
            2i32.pow(left_height as u32) - 1
        } else {
            let root = root.unwrap();
            let (left, right) = (root.borrow().left.clone(), root.borrow().right.clone());
            Solution::count_nodes(left) + Solution::count_nodes(right) + 1
        }
    }
}

fn main() {
    println!("{:?}", Solution::count_nodes(tree![1, 2, 3, 4, 5, 6]));
}
