// Accepted
// @lc id=669 lang=rust
// problem: trim_a_binary_search_tree
extern crate util;
use util::binary_tree::*;
struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    pub fn trim_bst(
        root: Option<Rc<RefCell<TreeNode>>>,
        l: i32,
        r: i32,
    ) -> Option<Rc<RefCell<TreeNode>>> {
        if let Some(root) = root {
            let root_val = root.borrow().val;
            if l <= root_val {
                let left = Solution::trim_bst(root.borrow().left.clone(), l, r);
                root.borrow_mut().left = left;
            }
            if r >= root_val {
                let right = Solution::trim_bst(root.borrow().right.clone(), l, r);
                root.borrow_mut().right = right;
            }
            if l > root_val {
                return root.borrow().right.clone();
            } else if r < root_val {
                return root.borrow().left.clone();
            } else {
                return Some(root);
            }
        }
        return None;
    }
}

fn main() {
    let root = BinaryTree::new(vec![Some(1), Some(0), Some(2)]);
    println!("{:?}", Solution::trim_bst(root, 0, 2));
}
