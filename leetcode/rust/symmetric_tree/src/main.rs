// Accepted
// @lc id=101 lang=rust
// problem: symmetric_tree
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn is_same(a: Link, b: Link) -> bool {
        if a.is_none() && b.is_none() {
            return true;
        }
        if a.is_none() || b.is_none() {
            return false;
        }
        let (a, b) = (a.unwrap(), b.unwrap());
        if a.borrow().val != b.borrow().val {
            return false;
        }
        let (a_left, a_right, b_left, b_right) = (
            a.borrow().left.clone(),
            a.borrow().right.clone(),
            b.borrow().left.clone(),
            b.borrow().right.clone(),
        );
        Solution::is_same(a_left, b_right) && Solution::is_same(a_right, b_left)
    }

    pub fn is_symmetric(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        if root.is_none() {
            return true;
        }
        let root = root.unwrap();
        let (left, right) = (root.borrow().left.clone(), root.borrow().right.clone());
        Solution::is_same(left, right)
    }
}

fn main() {
    println!("{:?}", Solution::is_symmetric(tree![1, 2, 2, 3, 4, 4, 3]));
}
