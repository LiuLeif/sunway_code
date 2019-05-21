// Accepted
// 2019-01-04 09:50
// @lc id=104 lang=rust
// problem: maximum_depth_of_binary_tree
use util::binary_tree::*;

struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn max_depth(root: Option<Rc<RefCell<TreeNode>>>) -> i32 {
        match root {
            Some(ref node) => {
                std::cmp::max(
                    Solution::max_depth(node.borrow().left.clone()),
                    Solution::max_depth(node.borrow().right.clone()),
                ) + 1
            }
            None => 0,
        }
    }
}

fn main() {
    let root = BinaryTree::new(vec![
        Some(3),
        Some(9),
        Some(20),
        None,
        None,
        Some(15),
        Some(7),
    ]);
    println!("{:?}", Solution::max_depth(root));
}
