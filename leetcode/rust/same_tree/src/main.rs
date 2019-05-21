// Accepted
// @lc id=100 lang=rust
// problem: same_tree
extern crate util;
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn is_same_tree(
        p: Option<Rc<RefCell<TreeNode>>>,
        q: Option<Rc<RefCell<TreeNode>>>,
    ) -> bool {
        if p.is_none() && q.is_none() {
            return true;
        }
        if p.is_none() || q.is_none() {
            return false;
        }
        let p = p.unwrap();
        let q = q.unwrap();
        return p.borrow().val == q.borrow().val
            && Solution::is_same_tree(p.borrow().left.clone(), q.borrow().left.clone())
            && Solution::is_same_tree(p.borrow().right.clone(), q.borrow().right.clone());
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::is_same_tree(tree!(1, 2, 3), tree!(1, 2, 3))
    );
}
