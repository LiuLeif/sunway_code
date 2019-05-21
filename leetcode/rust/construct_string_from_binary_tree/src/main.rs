// Accepted
// @lc id=606 lang=rust
// problem: construct_string_from_binary_tree
extern crate util;
use util::binary_tree::*;
use util::*;

struct Solution {}

use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn tree2str(t: Option<Rc<RefCell<TreeNode>>>) -> String {
        // "1(2()(4))(3)"
        if let Some(node) = t {
            let left = Solution::tree2str(node.borrow().left.clone());
            let right = Solution::tree2str(node.borrow().right.clone());
            let val = node.borrow().val;
            if left.is_empty() && right.is_empty() {
                return format!("{}", val);
            }
            if right.is_empty() {
                return format!("{}({})", val, left);
            }
            return format!("{}({})({})", val, left, right);
        } else {
            return "".to_owned();
        }
    }
}

fn main() {
    println!("{:?}", Solution::tree2str(tree!(1, 2, 3, None, 4)));
}
