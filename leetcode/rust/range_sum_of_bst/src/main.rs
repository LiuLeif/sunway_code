// 2018-12-24 17:59
// @lc id=938 lang=rust
extern crate util;
use util::binary_tree::{BinaryTree, TreeNode};

use std::cell::RefCell;
use std::rc::Rc;

struct Solution {}

impl Solution {
    pub fn range_sum_bst(root: Option<Rc<RefCell<TreeNode>>>, l: i32, r: i32) -> i32 {
        if root == None {
            return 0;
        }
        let mut ret = 0;

        let root = root.unwrap();
        let root_val = root.borrow().val;
        if root_val >= l && root_val <= r {
            ret += root_val;
        }
        if root_val <= r {
            ret += Solution::range_sum_bst(root.borrow().right.clone(), l, r);
        }
        if root_val >= l {
            ret += Solution::range_sum_bst(root.borrow().left.clone(), l, r);
        }
        return ret;
    }
}

fn main() {
    let root = BinaryTree::new(vec![
        Some(10),
        Some(5),
        Some(15),
        Some(3),
        Some(7),
        None,
        Some(18),
    ]);
    println!("{:?}", Solution::range_sum_bst(root, 7, 15));
}
