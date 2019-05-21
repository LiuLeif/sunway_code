// Accepted
// @lc id=965 lang=rust
// problem: univalued_binary_tree
extern crate util;
use util::binary_tree::*;
struct Solution {}
use std::cell::RefCell;
use std::rc::Rc;
impl Solution {
    fn traverse(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        if root == None {
            return vec![];
        }
        let root = root.unwrap();
        let mut ret = vec![root.borrow().val];

        let left = root.borrow().left.clone();
        let right = root.borrow().right.clone();
        
        ret.extend(Solution::traverse(left));
        ret.extend(Solution::traverse(right));
        ret
    }

    pub fn is_unival_tree(root: Option<Rc<RefCell<TreeNode>>>) -> bool {
        let v = Solution::traverse(root);
        v.windows(2).all(|w| w[0] == w[1])
    }
}

fn main() {
    let root = BinaryTree::new(vec![Some(1), Some(1), Some(1), Some(1)]);
    println!("{:?}", Solution::is_unival_tree(root));
}
