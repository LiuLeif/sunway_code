// Accepted
// @lc id=112 lang=rust
// problem: path_sum
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn has_path_sum(root: Option<Rc<RefCell<TreeNode>>>, sum: i32) -> bool {
        if root.is_none() {
            return false;
        }
        let root = root.unwrap();
        let (left, right, val) = (
            root.borrow().left.clone(),
            root.borrow().right.clone(),
            root.borrow().val,
        );
        if left.is_none() && right.is_none() && val == sum {
            return true;
        }
        Solution::has_path_sum(left, sum - val) || Solution::has_path_sum(right, sum - val)
    }
}

fn main() {
    println!(
        "{:?}",
        Solution::has_path_sum(tree![5, 4, 8, 11, None, 13, 4, 7, 2], 22)
    );
    println!("{:?}", Solution::has_path_sum(tree![1, 2], 1));
}
