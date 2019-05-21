// Accepted
// @lc id=538 lang=rust
// problem: convert_bst_to_greater_tree
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
type Link = Option<Rc<RefCell<TreeNode>>>;
impl Solution {
    fn traversal(root: Link, mut total: i32) -> i32 {
        if root.is_none() {
            return total;
        }
        let root = root.unwrap();
        let (left, right, val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            root.borrow().val,
        );
        total = Solution::traversal(right, total);
        root.borrow_mut().val += total;
        Solution::traversal(left, total + val)
    }

    pub fn convert_bst(root: Option<Rc<RefCell<TreeNode>>>) -> Option<Rc<RefCell<TreeNode>>> {
        Solution::traversal(root.clone(), 0);
        root
    }
}

fn main() {
    println!("{:?}", Solution::convert_bst(tree![5, 2, 13]));
}
