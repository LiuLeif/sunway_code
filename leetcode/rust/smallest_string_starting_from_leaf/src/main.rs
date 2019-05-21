// Accepted
// @lc id=988 lang=rust
// problem: smallest_string_starting_from_leaf
use std::collections::HashMap;
use std::collections::VecDeque;
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn smallest_from_leaf(root: Option<Rc<RefCell<TreeNode>>>) -> String {
        if root.is_none() {
            return "~".to_owned();
        }

        let root = root.unwrap();
        let (left, right, val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            (root.borrow().val as u8 + b'a') as char,
        );

        if left == right {
            return val.to_string();
        }
        return std::cmp::min(
            format!("{}{}", Solution::smallest_from_leaf(left), val),
            format!("{}{}", Solution::smallest_from_leaf(right), val),
        );
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::smallest_from_leaf(tree![2, 2, 1, None, 1, 0, None, 0])
    );
    println!(
        "{:?}",
        Solution::smallest_from_leaf(tree![3, 9, 20, None, None, 15, 7])
    );
}
