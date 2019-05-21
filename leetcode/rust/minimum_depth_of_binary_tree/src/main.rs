// Accepted
// @lc id=111 lang=rust
// problem: minimum_depth_of_binary_tree
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn min_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        if root.is_none() {
            return 0;
        }
        let root = root.unwrap();
        let (left, right) = (root.borrow().left.clone(), root.borrow().right.clone());
        if left.is_none() && right.is_none() {
            return 1;
        }
        let mut ret = std::i32::MAX;
        if left.is_some() {
            ret = std::cmp::min(Solution::min_depth(left), ret);
        }
        if right.is_some() {
            ret = std::cmp::min(Solution::min_depth(right), ret);
        }
        1 + ret
    }
}

fn main() {
    println!("{:?}", Solution::min_depth(tree![1, 2]));
}
